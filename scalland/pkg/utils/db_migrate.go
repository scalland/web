package utils

import (
	"fmt"
	"sort"
	"strings"
)

// MigrateDB runs all pending SQL migrations.
func (u *Utils) MigrateDB(createDB bool) error {
	db := u.DBMigrator
	if db == nil {
		return fmt.Errorf("migrator DB not connected")
	}

	// Create schema_migrations table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version INT NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	// Get applied versions
	applied := make(map[int]bool)
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("query applied: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var v int
		rows.Scan(&v)
		applied[v] = true
	}

	u.Logger.Infof("Found %d applied migrations", len(applied))
	// Migration files will be loaded from embedded FS at runtime
	// This is a placeholder for the migration runner logic
	return nil
}

// DebugSQL logs a parameterized query with its values substituted.
func DebugSQL(query string, args ...interface{}) string {
	q := query
	for _, arg := range args {
		q = strings.Replace(q, "?", fmt.Sprintf("'%v'", arg), 1)
	}
	return q
}

func sortedKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
