package mail

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"github.com/emersion/go-imap/v2"
)

// SyncOptions configures the sync operation
type SyncOptions struct {
	Workers       int
	IncludeBodies bool
	Progress      io.Writer // If set, progress is written here
}

// SyncResult contains the result of a sync operation
type SyncResult struct {
	Mailbox      string `json:"mailbox"`
	NewMessages  int    `json:"new_messages"`
	TotalCached  int    `json:"total_cached"`
	BodiesCached int    `json:"bodies_cached,omitempty"`
	Message      string `json:"message"`
}

// Sync fetches new messages from the server and updates the cache
func Sync(mailbox string, opts *SyncOptions) (*SyncResult, error) {
	if opts == nil {
		opts = &SyncOptions{}
	}
	if opts.Workers <= 0 {
		opts.Workers = 1
	}

	// Open cache
	cache, err := OpenCache()
	if err != nil {
		return nil, err
	}
	defer cache.Close()

	// Connect to IMAP to get mailbox info
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Select mailbox
	mbox, err := client.SelectMailbox(mailbox)
	if err != nil {
		return nil, err
	}

	// Check cache state
	state, err := cache.GetMailboxState(mailbox)
	if err != nil {
		return nil, err
	}

	// Check UIDVALIDITY - if it changed, we need to clear and resync
	if state != nil && state.UIDValidity != mbox.UIDValidity {
		if err := cache.ClearMailbox(mailbox); err != nil {
			return nil, fmt.Errorf("clearing stale cache: %w", err)
		}
		state = nil
	}

	result := &SyncResult{
		Mailbox: mailbox,
	}

	if mbox.NumMessages == 0 {
		if err := cache.UpdateMailboxState(mailbox, mbox.UIDValidity, 0); err != nil {
			return nil, err
		}
		result.Message = "mailbox is empty"
		return result, nil
	}

	// Get cached UIDs
	cachedUIDs, err := cache.GetCachedUIDs(mailbox)
	if err != nil {
		return nil, err
	}

	cachedBodyUIDs := map[uint32]bool{}
	if opts.IncludeBodies {
		cachedBodyUIDs, err = cache.GetCachedBodyUIDs(mailbox)
		if err != nil {
			return nil, err
		}
	}

	// Get all server UIDs
	if opts.Progress != nil {
		fmt.Fprintf(opts.Progress, "Checking for new messages...\n")
	}
	serverUIDs, err := client.GetAllUIDs()
	if err != nil {
		return nil, fmt.Errorf("getting server UIDs: %w", err)
	}

	// Find missing UIDs
	var missingUIDs []imap.UID
	for _, uid := range serverUIDs {
		if !cachedUIDs[uint32(uid)] {
			missingUIDs = append(missingUIDs, uid)
		}
	}

	if len(missingUIDs) == 0 {
		// Update state to track current max UID
		if len(serverUIDs) > 0 {
			maxUID := serverUIDs[len(serverUIDs)-1]
			if err := cache.UpdateMailboxState(mailbox, mbox.UIDValidity, uint32(maxUID)); err != nil {
				return nil, err
			}
		}

		result.TotalCached = len(cachedUIDs)
		if opts.IncludeBodies {
			bodyCount, err := cache.CountMessageBodies(mailbox)
			if err != nil {
				return nil, err
			}
			result.BodiesCached = bodyCount
		}
		result.Message = "already up to date"
	} else {
		if opts.Progress != nil {
			fmt.Fprintf(opts.Progress, "Found %d messages to sync\n", len(missingUIDs))
		}

		// Fetch missing UIDs in parallel
		var newCount int
		if opts.Workers > 1 && len(missingUIDs) > 100 {
			newCount, err = fetchMissingUIDsParallel(cache, mailbox, mbox.UIDValidity, missingUIDs, opts)
		} else {
			newCount, err = fetchMissingUIDs(cache, mailbox, mbox.UIDValidity, missingUIDs, opts)
		}
		if err != nil {
			return nil, err
		}

		result.NewMessages = newCount
		result.TotalCached = len(cachedUIDs) + newCount

		if result.NewMessages == 0 {
			result.Message = "already up to date"
		} else {
			result.Message = fmt.Sprintf("synced %d new messages", result.NewMessages)
		}
	}

	if opts.IncludeBodies {
		var missingBodyUIDs []imap.UID
		for _, uid := range serverUIDs {
			if !cachedBodyUIDs[uint32(uid)] {
				missingBodyUIDs = append(missingBodyUIDs, uid)
			}
		}

		bodyCount, err := cache.CountMessageBodies(mailbox)
		if err != nil {
			return nil, err
		}
		result.BodiesCached = bodyCount

		if len(missingBodyUIDs) > 0 {
			if opts.Progress != nil {
				fmt.Fprintf(opts.Progress, "Found %d message bodies to cache\n", len(missingBodyUIDs))
			}

			if opts.Workers > 1 && len(missingBodyUIDs) > 25 {
				if err := fetchMissingBodiesParallel(cache, mailbox, missingBodyUIDs, opts); err != nil {
					return nil, err
				}
			} else {
				if err := fetchMissingBodies(cache, mailbox, missingBodyUIDs, opts); err != nil {
					return nil, err
				}
			}

			bodyCount, err = cache.CountMessageBodies(mailbox)
			if err != nil {
				return nil, err
			}
			result.BodiesCached = bodyCount
		}
	}

	return result, nil
}

// fetchMissingUIDs fetches specific UIDs sequentially
func fetchMissingUIDs(cache *Cache, mailbox string, uidValidity uint32, uids []imap.UID, opts *SyncOptions) (int, error) {
	const batchSize = 500

	client, err := Connect()
	if err != nil {
		return 0, err
	}
	defer client.Close()

	if _, err := client.SelectMailbox(mailbox); err != nil {
		return 0, err
	}

	var totalFetched int
	var maxUID imap.UID

	for i := 0; i < len(uids); i += batchSize {
		end := i + batchSize
		if end > len(uids) {
			end = len(uids)
		}

		batch := uids[i:end]
		envelopes, err := client.FetchEnvelopesByUID(batch)
		if err != nil {
			return totalFetched, err
		}

		if len(envelopes) > 0 {
			if err := cache.InsertEnvelopes(mailbox, envelopes); err != nil {
				return totalFetched, err
			}

			for _, env := range envelopes {
				if env.UID > maxUID {
					maxUID = env.UID
				}
			}

			if err := cache.UpdateMailboxState(mailbox, uidValidity, uint32(maxUID)); err != nil {
				return totalFetched, err
			}
		}

		totalFetched += len(envelopes)

		if opts.Progress != nil {
			fmt.Fprintf(opts.Progress, "\rSyncing: %d / %d messages (%.1f%%)", totalFetched, len(uids), float64(totalFetched)/float64(len(uids))*100)
		}
	}

	if opts.Progress != nil {
		fmt.Fprintln(opts.Progress)
	}

	return totalFetched, nil
}

// fetchMissingUIDsParallel fetches specific UIDs using multiple parallel connections
func fetchMissingUIDsParallel(cache *Cache, mailbox string, uidValidity uint32, uids []imap.UID, opts *SyncOptions) (int, error) {
	const batchSize = 500

	numWorkers := opts.Workers
	if numWorkers > len(uids)/batchSize+1 {
		numWorkers = len(uids)/batchSize + 1
	}
	if numWorkers < 1 {
		numWorkers = 1
	}

	type batch struct {
		envelopes []Envelope
		err       error
	}
	batchChan := make(chan batch, numWorkers*2)

	var wg sync.WaitGroup
	var fetched atomic.Uint64

	// Progress reporter
	var progressDone chan struct{}
	if opts.Progress != nil {
		progressDone = make(chan struct{})
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-progressDone:
					return
				case <-ticker.C:
					f := fetched.Load()
					fmt.Fprintf(opts.Progress, "\rSyncing: %d / %d messages (%.1f%%)", f, len(uids), float64(f)/float64(len(uids))*100)
				}
			}
		}()
	}

	// Divide UIDs among workers
	uidsPerWorker := len(uids) / numWorkers
	remainder := len(uids) % numWorkers

	start := 0
	for i := 0; i < numWorkers; i++ {
		count := uidsPerWorker
		if i < remainder {
			count++
		}
		if count == 0 {
			continue
		}

		workerUIDs := uids[start : start+count]
		start += count

		wg.Add(1)
		go func(myUIDs []imap.UID) {
			defer wg.Done()

			client, err := Connect()
			if err != nil {
				batchChan <- batch{nil, fmt.Errorf("connect: %w", err)}
				return
			}
			defer client.Close()

			if _, err := client.SelectMailbox(mailbox); err != nil {
				batchChan <- batch{nil, fmt.Errorf("select: %w", err)}
				return
			}

			// Fetch in batches
			for i := 0; i < len(myUIDs); i += batchSize {
				end := i + batchSize
				if end > len(myUIDs) {
					end = len(myUIDs)
				}

				envs, err := client.FetchEnvelopesByUID(myUIDs[i:end])
				if err != nil {
					batchChan <- batch{nil, fmt.Errorf("fetch: %w", err)}
					return
				}

				batchChan <- batch{envs, nil}
				fetched.Add(uint64(len(envs)))
			}
		}(workerUIDs)
	}

	// Close batch channel when all workers are done
	go func() {
		wg.Wait()
		close(batchChan)
	}()

	// Write batches to cache as they arrive
	var totalFetched int
	var maxUID imap.UID
	var writeErr error
	var batchesSinceStateUpdate int

	for b := range batchChan {
		if b.err != nil {
			writeErr = b.err
			continue
		}
		if writeErr != nil {
			continue
		}

		if len(b.envelopes) > 0 {
			if err := cache.InsertEnvelopes(mailbox, b.envelopes); err != nil {
				writeErr = err
				continue
			}

			for _, env := range b.envelopes {
				if env.UID > maxUID {
					maxUID = env.UID
				}
			}

			totalFetched += len(b.envelopes)
			batchesSinceStateUpdate++

			if batchesSinceStateUpdate >= 10 && maxUID > 0 {
				cache.UpdateMailboxState(mailbox, uidValidity, uint32(maxUID))
				batchesSinceStateUpdate = 0
			}
		}
	}

	if opts.Progress != nil {
		close(progressDone)
		fmt.Fprintf(opts.Progress, "\rSyncing: %d / %d messages (100.0%%)\n", len(uids), len(uids))
	}

	if writeErr != nil {
		return totalFetched, writeErr
	}

	// Update final state
	if maxUID > 0 {
		if err := cache.UpdateMailboxState(mailbox, uidValidity, uint32(maxUID)); err != nil {
			return totalFetched, err
		}
	}

	return totalFetched, nil
}

func fetchMissingBodies(cache *Cache, mailbox string, uids []imap.UID, opts *SyncOptions) error {
	client, err := Connect()
	if err != nil {
		return err
	}
	defer client.Close()

	if _, err := client.SelectMailbox(mailbox); err != nil {
		return err
	}

	for i, uid := range uids {
		body, _, err := client.FetchMessage(uid)
		if err != nil {
			return err
		}
		if err := cache.UpsertMessageBody(mailbox, uint32(uid), body); err != nil {
			return err
		}

		if opts.Progress != nil {
			fmt.Fprintf(opts.Progress, "\rCaching bodies: %d / %d messages (%.1f%%)", i+1, len(uids), float64(i+1)/float64(len(uids))*100)
		}
	}

	if opts.Progress != nil {
		fmt.Fprintln(opts.Progress)
	}

	return nil
}

func fetchMissingBodiesParallel(cache *Cache, mailbox string, uids []imap.UID, opts *SyncOptions) error {
	numWorkers := opts.Workers
	if numWorkers > len(uids) {
		numWorkers = len(uids)
	}
	if numWorkers < 1 {
		numWorkers = 1
	}

	type bodyBatch struct {
		uid  imap.UID
		body []byte
		err  error
	}

	batchChan := make(chan bodyBatch, numWorkers*2)
	var fetched atomic.Uint64
	var wg sync.WaitGroup

	var progressDone chan struct{}
	if opts.Progress != nil {
		progressDone = make(chan struct{})
		go func() {
			ticker := time.NewTicker(5 * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-progressDone:
					return
				case <-ticker.C:
					f := fetched.Load()
					fmt.Fprintf(opts.Progress, "\rCaching bodies: %d / %d messages (%.1f%%)", f, len(uids), float64(f)/float64(len(uids))*100)
				}
			}
		}()
	}

	uidsPerWorker := len(uids) / numWorkers
	remainder := len(uids) % numWorkers
	start := 0

	for i := 0; i < numWorkers; i++ {
		count := uidsPerWorker
		if i < remainder {
			count++
		}
		if count == 0 {
			continue
		}

		workerUIDs := uids[start : start+count]
		start += count

		wg.Add(1)
		go func(myUIDs []imap.UID) {
			defer wg.Done()

			client, err := Connect()
			if err != nil {
				batchChan <- bodyBatch{err: fmt.Errorf("connect: %w", err)}
				return
			}
			defer client.Close()

			if _, err := client.SelectMailbox(mailbox); err != nil {
				batchChan <- bodyBatch{err: fmt.Errorf("select: %w", err)}
				return
			}

			for _, uid := range myUIDs {
				body, _, err := client.FetchMessage(uid)
				if err != nil {
					batchChan <- bodyBatch{err: fmt.Errorf("fetch body %d: %w", uid, err)}
					return
				}
				batchChan <- bodyBatch{uid: uid, body: body}
				fetched.Add(1)
			}
		}(workerUIDs)
	}

	go func() {
		wg.Wait()
		close(batchChan)
	}()

	var writeErr error
	for b := range batchChan {
		if b.err != nil {
			writeErr = b.err
			continue
		}
		if writeErr != nil {
			continue
		}
		if err := cache.UpsertMessageBody(mailbox, uint32(b.uid), b.body); err != nil {
			writeErr = err
		}
	}

	if opts.Progress != nil {
		close(progressDone)
		fmt.Fprintf(opts.Progress, "\rCaching bodies: %d / %d messages (100.0%%)\n", len(uids), len(uids))
	}

	return writeErr
}
