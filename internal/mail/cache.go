package mail

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

// Cache provides SQLite-based email metadata caching
type Cache struct {
	db *sql.DB
}

// OpenCache opens or creates the cache database
func OpenCache() (*Cache, error) {
	path := CacheFilePath()

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening cache database: %w", err)
	}

	cache := &Cache{db: db}
	if err := cache.init(); err != nil {
		db.Close()
		return nil, err
	}

	return cache, nil
}

// Close closes the cache database
func (c *Cache) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

func (c *Cache) init() error {
	schema := `
		CREATE TABLE IF NOT EXISTS mailboxes (
			name TEXT PRIMARY KEY,
			uidvalidity INTEGER NOT NULL,
			last_uid INTEGER NOT NULL DEFAULT 0,
			last_sync INTEGER NOT NULL DEFAULT 0
		);

		CREATE TABLE IF NOT EXISTS envelopes (
			mailbox TEXT NOT NULL,
			uid INTEGER NOT NULL,
			message_id TEXT,
			date INTEGER,
			from_addr TEXT,
			from_name TEXT,
			subject TEXT,
			PRIMARY KEY (mailbox, uid)
		);

		CREATE INDEX IF NOT EXISTS idx_envelopes_date ON envelopes(mailbox, date DESC);
		CREATE INDEX IF NOT EXISTS idx_envelopes_from ON envelopes(mailbox, from_addr);
		CREATE INDEX IF NOT EXISTS idx_envelopes_subject ON envelopes(mailbox, subject);
	`

	_, err := c.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("initializing cache schema: %w", err)
	}

	return nil
}

// MailboxState holds sync state for a mailbox
type MailboxState struct {
	Name        string
	UIDValidity uint32
	LastUID     uint32
	LastSync    time.Time
}

// GetMailboxState returns the cached state for a mailbox
func (c *Cache) GetMailboxState(mailbox string) (*MailboxState, error) {
	row := c.db.QueryRow(
		`SELECT name, uidvalidity, last_uid, last_sync FROM mailboxes WHERE name = ?`,
		mailbox,
	)

	var state MailboxState
	var lastSyncUnix int64
	err := row.Scan(&state.Name, &state.UIDValidity, &state.LastUID, &lastSyncUnix)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying mailbox state: %w", err)
	}

	state.LastSync = time.Unix(lastSyncUnix, 0)
	return &state, nil
}

// UpdateMailboxState updates the sync state for a mailbox
func (c *Cache) UpdateMailboxState(mailbox string, uidValidity, lastUID uint32) error {
	_, err := c.db.Exec(
		`INSERT INTO mailboxes (name, uidvalidity, last_uid, last_sync)
		 VALUES (?, ?, ?, ?)
		 ON CONFLICT(name) DO UPDATE SET
			uidvalidity = excluded.uidvalidity,
			last_uid = excluded.last_uid,
			last_sync = excluded.last_sync`,
		mailbox, uidValidity, lastUID, time.Now().Unix(),
	)
	if err != nil {
		return fmt.Errorf("updating mailbox state: %w", err)
	}
	return nil
}

// ClearMailbox removes all cached data for a mailbox (used when UIDVALIDITY changes)
func (c *Cache) ClearMailbox(mailbox string) error {
	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM envelopes WHERE mailbox = ?`, mailbox); err != nil {
		return fmt.Errorf("clearing envelopes: %w", err)
	}

	if _, err := tx.Exec(`DELETE FROM mailboxes WHERE name = ?`, mailbox); err != nil {
		return fmt.Errorf("clearing mailbox state: %w", err)
	}

	return tx.Commit()
}

// CachedEnvelope represents a cached email envelope
type CachedEnvelope struct {
	UID       uint32
	MessageID string
	Date      time.Time
	FromAddr  string
	FromName  string
	Subject   string
}

// InsertEnvelopes adds envelopes to the cache
func (c *Cache) InsertEnvelopes(mailbox string, envelopes []Envelope) error {
	if len(envelopes) == 0 {
		return nil
	}

	tx, err := c.db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(
		`INSERT OR REPLACE INTO envelopes (mailbox, uid, message_id, date, from_addr, from_name, subject)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
	)
	if err != nil {
		return fmt.Errorf("preparing insert: %w", err)
	}
	defer stmt.Close()

	for _, env := range envelopes {
		_, err := stmt.Exec(mailbox, uint32(env.UID), env.MessageID, env.Date, env.FromAddr, env.FromName, env.Subject)
		if err != nil {
			return fmt.Errorf("inserting envelope: %w", err)
		}
	}

	return tx.Commit()
}

// SearchOptions specifies search filters
type SearchOptions struct {
	From    string
	Subject string
	Since   *time.Time
	Before  *time.Time
	Limit   int
}

// SearchResult contains search results with cache metadata
type SearchResult struct {
	Mailbox     string           `json:"mailbox"`
	LastSync    time.Time        `json:"last_sync"`
	Freshness   string           `json:"freshness"`
	TotalCached int              `json:"total_cached"`
	Results     []CachedEnvelope `json:"results"`
	Count       int              `json:"count"`
}

// Search queries the cache for matching envelopes
func (c *Cache) Search(mailbox string, opts *SearchOptions) (*SearchResult, error) {
	// Get mailbox state for freshness info
	state, err := c.GetMailboxState(mailbox)
	if err != nil {
		return nil, err
	}

	result := &SearchResult{
		Mailbox: mailbox,
		Results: []CachedEnvelope{},
	}

	if state != nil {
		result.LastSync = state.LastSync
		result.Freshness = formatFreshness(state.LastSync)
	} else {
		result.Freshness = "never synced"
	}

	// Count total cached
	row := c.db.QueryRow(`SELECT COUNT(*) FROM envelopes WHERE mailbox = ?`, mailbox)
	row.Scan(&result.TotalCached)

	// Build query
	query := `SELECT uid, message_id, date, from_addr, from_name, subject
			  FROM envelopes WHERE mailbox = ?`
	args := []any{mailbox}

	if opts != nil {
		if opts.From != "" {
			query += ` AND from_addr LIKE ?`
			args = append(args, "%"+opts.From+"%")
		}
		if opts.Subject != "" {
			query += ` AND subject LIKE ?`
			args = append(args, "%"+opts.Subject+"%")
		}
		if opts.Since != nil {
			query += ` AND date >= ?`
			args = append(args, opts.Since.Unix())
		}
		if opts.Before != nil {
			query += ` AND date < ?`
			args = append(args, opts.Before.Unix())
		}
	}

	query += ` ORDER BY date DESC`

	limit := 50
	if opts != nil && opts.Limit > 0 {
		limit = opts.Limit
	}
	query += fmt.Sprintf(` LIMIT %d`, limit)

	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("searching cache: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var env CachedEnvelope
		var dateUnix int64
		var messageID, fromAddr, fromName, subject sql.NullString

		err := rows.Scan(&env.UID, &messageID, &dateUnix, &fromAddr, &fromName, &subject)
		if err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}

		env.MessageID = messageID.String
		env.Date = time.Unix(dateUnix, 0)
		env.FromAddr = fromAddr.String
		env.FromName = fromName.String
		env.Subject = subject.String

		result.Results = append(result.Results, env)
	}

	result.Count = len(result.Results)
	return result, nil
}

// GetEnvelope retrieves a single envelope by UID
func (c *Cache) GetEnvelope(mailbox string, uid uint32) (*CachedEnvelope, error) {
	row := c.db.QueryRow(
		`SELECT uid, message_id, date, from_addr, from_name, subject
		 FROM envelopes WHERE mailbox = ? AND uid = ?`,
		mailbox, uid,
	)

	var env CachedEnvelope
	var dateUnix int64
	var messageID, fromAddr, fromName, subject sql.NullString

	err := row.Scan(&env.UID, &messageID, &dateUnix, &fromAddr, &fromName, &subject)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying envelope: %w", err)
	}

	env.MessageID = messageID.String
	env.Date = time.Unix(dateUnix, 0)
	env.FromAddr = fromAddr.String
	env.FromName = fromName.String
	env.Subject = subject.String

	return &env, nil
}

func formatFreshness(t time.Time) string {
	if t.IsZero() {
		return "never synced"
	}

	d := time.Since(t)

	switch {
	case d < time.Minute:
		return "just now"
	case d < time.Hour:
		mins := int(d.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", mins)
	case d < 24*time.Hour:
		hours := int(d.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	default:
		days := int(d.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}
}
