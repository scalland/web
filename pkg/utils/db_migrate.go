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

	// Walk migrations FS and apply pending ones
	entries, err := u.MigrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	// Build map of version -> filename for .sql files
	pending := make(map[int]string)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		var version int
		var name string
		_, scanErr := fmt.Sscanf(entry.Name(), "migration_%d", &version)
		if scanErr != nil {
			u.Logger.Warnf("Skipping non-standard migration file: %s", entry.Name())
			continue
		}
		name = strings.TrimSuffix(entry.Name(), ".sql")
		if !applied[version] {
			pending[version] = name
		}
	}

	if len(pending) == 0 {
		u.Logger.Infof("No pending migrations")
		return nil
	}

	for _, version := range sortedKeys(pending) {
		name := pending[version]
		filename := fmt.Sprintf("migrations/migration_%03d.sql", version)
		sqlBytes, readErr := u.MigrationsFS.ReadFile(filename)
		if readErr != nil {
			return fmt.Errorf("read migration %s: %w", filename, readErr)
		}

		// Split on semicolons and execute each statement
		stmts := strings.Split(string(sqlBytes), ";")
		for _, stmt := range stmts {
			// Strip comment lines and whitespace
			lines := strings.Split(stmt, "\n")
			var codeLines []string
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed != "" && !strings.HasPrefix(trimmed, "--") {
					codeLines = append(codeLines, trimmed)
				}
			}
			stmt = strings.TrimSpace(strings.Join(codeLines, "\n"))
			if stmt == "" {
				continue
			}
			if _, execErr := db.Exec(stmt); execErr != nil {
				return fmt.Errorf("execute migration %s: %w", filename, execErr)
			}
		}

		// Record as applied
		_, recErr := db.Exec("INSERT INTO schema_migrations (version, name) VALUES (?, ?)", version, name)
		if recErr != nil {
			return fmt.Errorf("record migration %s: %w", name, recErr)
		}
		u.Logger.Infof("Applied migration: %s", name)
	}

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
