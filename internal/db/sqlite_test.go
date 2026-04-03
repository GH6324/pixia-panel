package db

import (
	"context"
	"database/sql"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

func TestBuildSQLiteDSN(t *testing.T) {
	dsn := buildSQLiteDSN("pixia.db")
	wantParts := []string{
		"file:pixia.db",
		"_pragma=journal_mode%28WAL%29",
		"_pragma=busy_timeout%285000%29",
		"_pragma=foreign_keys%28ON%29",
	}
	for _, want := range wantParts {
		if !strings.Contains(dsn, want) {
			t.Fatalf("dsn %q missing %q", dsn, want)
		}
	}
}

func TestOpenAppliesPragmasToEveryConnection(t *testing.T) {
	db, err := Open(t.TempDir() + "/pixia.db")
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(2)

	ctx := context.Background()
	conn1, err := db.Conn(ctx)
	if err != nil {
		t.Fatalf("conn1: %v", err)
	}
	defer conn1.Close()

	conn2, err := db.Conn(ctx)
	if err != nil {
		t.Fatalf("conn2: %v", err)
	}
	defer conn2.Close()

	for idx, conn := range []*sql.Conn{conn1, conn2} {
		if got := pragmaInt(t, conn, "busy_timeout"); got != 5000 {
			t.Fatalf("conn%d busy_timeout=%d want 5000", idx+1, got)
		}
		if got := pragmaInt(t, conn, "foreign_keys"); got != 1 {
			t.Fatalf("conn%d foreign_keys=%d want 1", idx+1, got)
		}
		if got := pragmaText(t, conn, "journal_mode"); got != "wal" {
			t.Fatalf("conn%d journal_mode=%q want wal", idx+1, got)
		}
	}
}

func pragmaInt(t *testing.T, conn *sql.Conn, name string) int {
	t.Helper()
	var got int
	if err := conn.QueryRowContext(context.Background(), "PRAGMA "+name).Scan(&got); err != nil {
		t.Fatalf("pragma %s: %v", name, err)
	}
	return got
}

func pragmaText(t *testing.T, conn *sql.Conn, name string) string {
	t.Helper()
	var got string
	if err := conn.QueryRowContext(context.Background(), "PRAGMA "+name).Scan(&got); err != nil {
		t.Fatalf("pragma %s: %v", name, err)
	}
	return got
}
