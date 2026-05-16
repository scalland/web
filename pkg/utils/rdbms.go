package utils

import (
	"database/sql"
	"fmt"
)

// QueryRow executes a query and scans a single row.
func QueryRow(db *sql.DB, query string, args []interface{}, dest ...interface{}) error {
	return db.QueryRow(query, args...).Scan(dest...)
}

// ExecInsert executes an INSERT and returns the last insert ID.
func ExecInsert(db *sql.DB, query string, args ...interface{}) (int64, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// ExecUpdate executes an UPDATE/DELETE and returns rows affected.
func ExecUpdate(db *sql.DB, query string, args ...interface{}) (int64, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// InPlaceholders generates "?,?,?" for IN clauses.
func InPlaceholders(n int) string {
	if n <= 0 {
		return ""
	}
	s := "?"
	for i := 1; i < n; i++ {
		s += ",?"
	}
	return fmt.Sprintf("(%s)", s)
}
