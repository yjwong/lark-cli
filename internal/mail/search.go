package mail

import (
	"time"
)

// Search performs a local cache search with the given options
func Search(mailbox string, opts *SearchOptions) (*SearchResult, error) {
	cache, err := OpenCache()
	if err != nil {
		return nil, err
	}
	defer cache.Close()

	return cache.Search(mailbox, opts)
}

// ParseSearchOptions parses command-line style options into SearchOptions
func ParseSearchOptions(from, subject, since, before string, limit int) (*SearchOptions, error) {
	opts := &SearchOptions{
		From:    from,
		Subject: subject,
		Limit:   limit,
	}

	if since != "" {
		t, err := time.Parse("2006-01-02", since)
		if err != nil {
			return nil, err
		}
		opts.Since = &t
	}

	if before != "" {
		t, err := time.Parse("2006-01-02", before)
		if err != nil {
			return nil, err
		}
		opts.Before = &t
	}

	return opts, nil
}
