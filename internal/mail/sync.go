package mail

import (
	"fmt"

	"github.com/emersion/go-imap/v2"
)

// SyncResult contains the result of a sync operation
type SyncResult struct {
	Mailbox     string `json:"mailbox"`
	NewMessages int    `json:"new_messages"`
	TotalCached int    `json:"total_cached"`
	Message     string `json:"message"`
}

// Sync fetches new messages from the server and updates the cache
func Sync(mailbox string) (*SyncResult, error) {
	// Open cache
	cache, err := OpenCache()
	if err != nil {
		return nil, err
	}
	defer cache.Close()

	// Connect to IMAP
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
		state = nil // Force full sync
	}

	var lastUID imap.UID
	if state != nil {
		lastUID = imap.UID(state.LastUID)
	}

	result := &SyncResult{
		Mailbox: mailbox,
	}

	if mbox.NumMessages == 0 {
		// Empty mailbox
		if err := cache.UpdateMailboxState(mailbox, mbox.UIDValidity, 0); err != nil {
			return nil, err
		}
		result.Message = "mailbox is empty"
		return result, nil
	}

	var envelopes []Envelope

	if lastUID == 0 {
		// Full sync - fetch all envelopes in batches
		envelopes, err = fetchAllEnvelopes(client, mbox.NumMessages)
		if err != nil {
			return nil, err
		}
	} else {
		// Incremental sync - only fetch new messages
		envelopes, err = client.FetchNewEnvelopes(lastUID)
		if err != nil {
			return nil, err
		}
	}

	// Insert into cache
	if len(envelopes) > 0 {
		if err := cache.InsertEnvelopes(mailbox, envelopes); err != nil {
			return nil, err
		}

		// Find max UID
		var maxUID imap.UID
		for _, env := range envelopes {
			if env.UID > maxUID {
				maxUID = env.UID
			}
		}
		lastUID = maxUID
	}

	// Update mailbox state
	if err := cache.UpdateMailboxState(mailbox, mbox.UIDValidity, uint32(lastUID)); err != nil {
		return nil, err
	}

	result.NewMessages = len(envelopes)

	// Get total cached count
	searchResult, err := cache.Search(mailbox, &SearchOptions{Limit: 1})
	if err == nil {
		result.TotalCached = searchResult.TotalCached
	}

	if result.NewMessages == 0 {
		result.Message = "already up to date"
	} else {
		result.Message = fmt.Sprintf("synced %d new messages", result.NewMessages)
	}

	return result, nil
}

// fetchAllEnvelopes fetches all envelopes in batches
func fetchAllEnvelopes(client *Client, numMessages uint32) ([]Envelope, error) {
	const batchSize = 100

	var allEnvelopes []Envelope

	for end := numMessages; end >= 1; {
		start := end - batchSize + 1
		if start < 1 {
			start = 1
		}

		envelopes, err := client.FetchEnvelopes(start, end)
		if err != nil {
			return nil, err
		}

		allEnvelopes = append(allEnvelopes, envelopes...)
		end = start - 1
	}

	return allEnvelopes, nil
}
