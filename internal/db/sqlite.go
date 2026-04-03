package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"
)

// Open opens a SQLite database and applies required PRAGMAs.
// The caller must register a SQLite driver (e.g. modernc.org/sqlite or mattn/go-sqlite3).
func Open(path string) (*sql.DB, error) {
	if strings.TrimSpace(path) == "" {
		return nil, fmt.Errorf("db path is empty")
	}

	db, err := sql.Open("sqlite", buildSQLiteDSN(path))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func buildSQLiteDSN(path string) string {
	base := strings.TrimSpace(path)
	if !strings.HasPrefix(base, "file:") {
		base = "file:" + base
	}

	u, err := url.Parse(base)
	if err != nil {
		return base
	}

	q := u.Query()
	q.Add("_pragma", "journal_mode(WAL)")
	q.Add("_pragma", "busy_timeout(5000)")
	q.Add("_pragma", "foreign_keys(ON)")
	u.RawQuery = q.Encode()
	return u.String()
}
