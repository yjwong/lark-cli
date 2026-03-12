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

		CREATE TABLE IF NOT EXISTS message_bodies (
			mailbox TEXT NOT NULL,
			uid INTEGER NOT NULL,
			body BLOB NOT NULL,
			fetched_at INTEGER NOT NULL,
			PRIMARY KEY (mailbox, uid)
		);
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

// CountEnvelopes returns the number of cached envelopes for a mailbox
func (c *Cache) CountEnvelopes(mailbox string) (int, error) {
	var count int
	err := c.db.QueryRow(`SELECT COUNT(*) FROM envelopes WHERE mailbox = ?`, mailbox).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("counting envelopes: %w", err)
	}
	return count, nil
}

// GetCachedUIDs returns all cached UIDs for a mailbox as a set
func (c *Cache) GetCachedUIDs(mailbox string) (map[uint32]bool, error) {
	rows, err := c.db.Query(`SELECT uid FROM envelopes WHERE mailbox = ?`, mailbox)
	if err != nil {
		return nil, fmt.Errorf("querying cached UIDs: %w", err)
	}
	defer rows.Close()

	uids := make(map[uint32]bool)
	for rows.Next() {
		var uid uint32
		if err := rows.Scan(&uid); err != nil {
			return nil, fmt.Errorf("scanning UID: %w", err)
		}
		uids[uid] = true
	}
	return uids, nil
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

	if _, err := tx.Exec(`DELETE FROM message_bodies WHERE mailbox = ?`, mailbox); err != nil {
		return fmt.Errorf("clearing message bodies: %w", err)
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

// CachedMessageBody represents a stored RFC822 message body.
type CachedMessageBody struct {
	UID       uint32
	Body      []byte
	FetchedAt time.Time
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

// UpsertMessageBody stores a single RFC822 message body in the cache.
func (c *Cache) UpsertMessageBody(mailbox string, uid uint32, body []byte) error {
	_, err := c.db.Exec(
		`INSERT INTO message_bodies (mailbox, uid, body, fetched_at)
		 VALUES (?, ?, ?, ?)
		 ON CONFLICT(mailbox, uid) DO UPDATE SET
			body = excluded.body,
			fetched_at = excluded.fetched_at`,
		mailbox, uid, body, time.Now().Unix(),
	)
	if err != nil {
		return fmt.Errorf("storing message body: %w", err)
	}
	return nil
}

// GetCachedBodyUIDs returns the UIDs that already have cached bodies.
func (c *Cache) GetCachedBodyUIDs(mailbox string) (map[uint32]bool, error) {
	rows, err := c.db.Query(`SELECT uid FROM message_bodies WHERE mailbox = ?`, mailbox)
	if err != nil {
		return nil, fmt.Errorf("querying cached body UIDs: %w", err)
	}
	defer rows.Close()

	uids := make(map[uint32]bool)
	for rows.Next() {
		var uid uint32
		if err := rows.Scan(&uid); err != nil {
			return nil, fmt.Errorf("scanning cached body UID: %w", err)
		}
		uids[uid] = true
	}
	return uids, nil
}

// CountMessageBodies returns the number of cached bodies for a mailbox.
func (c *Cache) CountMessageBodies(mailbox string) (int, error) {
	var count int
	err := c.db.QueryRow(`SELECT COUNT(*) FROM message_bodies WHERE mailbox = ?`, mailbox).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("counting message bodies: %w", err)
	}
	return count, nil
}

// GetMessageBody retrieves a cached RFC822 message by UID.
func (c *Cache) GetMessageBody(mailbox string, uid uint32) (*CachedMessageBody, error) {
	row := c.db.QueryRow(
		`SELECT uid, body, fetched_at FROM message_bodies WHERE mailbox = ? AND uid = ?`,
		mailbox, uid,
	)

	var msg CachedMessageBody
	var fetchedAt int64
	err := row.Scan(&msg.UID, &msg.Body, &fetchedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("querying cached message body: %w", err)
	}

	msg.FetchedAt = time.Unix(fetchedAt, 0)
	return &msg, nil
}

// SearchOptions specifies search filters
type SearchOptions struct {
	From    string
	Subject string
	Body    string
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
	query := `SELECT e.uid, e.message_id, e.date, e.from_addr, e.from_name, e.subject
			  FROM envelopes e`
	if opts != nil && opts.Body != "" {
		query += ` JOIN message_bodies mb ON mb.mailbox = e.mailbox AND mb.uid = e.uid`
	}
	query += ` WHERE e.mailbox = ?`
	args := []any{mailbox}

	if opts != nil {
		if opts.From != "" {
			query += ` AND e.from_addr LIKE ?`
			args = append(args, "%"+opts.From+"%")
		}
		if opts.Subject != "" {
			query += ` AND e.subject LIKE ?`
			args = append(args, "%"+opts.Subject+"%")
		}
		if opts.Body != "" {
			query += ` AND CAST(mb.body AS TEXT) LIKE ?`
			args = append(args, "%"+opts.Body+"%")
		}
		if opts.Since != nil {
			query += ` AND e.date >= ?`
			args = append(args, opts.Since.Unix())
		}
		if opts.Before != nil {
			query += ` AND e.date < ?`
			args = append(args, opts.Before.Unix())
		}
	}

	query += ` ORDER BY e.date DESC`

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
