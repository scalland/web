package utils

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

// ConnectWriter opens the writer DB connection.
func (u *Utils) ConnectWriter() {
	cfg := u.Config.Databases["writer"]
	db, err := connectDB(cfg)
	if err != nil {
		u.Logger.Fatalf("Writer DB: %s", err.Error())
	}
	u.DBWriter = db
	u.Logger.Infof("Writer DB connected")
}

// ConnectReader opens the reader DB connection.
func (u *Utils) ConnectReader() {
	cfg, ok := u.Config.Databases["reader"]
	if !ok {
		cfg = u.Config.Databases["writer"]
	}
	db, err := connectDB(cfg)
	if err != nil {
		u.Logger.Fatalf("Reader DB: %s", err.Error())
	}
	u.DBReader = db
	u.Logger.Infof("Reader DB connected")
}

// ConnectMigrator opens the migrator DB connection.
func (u *Utils) ConnectMigrator() {
	cfg, ok := u.Config.Databases["migrator"]
	if !ok {
		cfg = u.Config.Databases["writer"]
	}
	db, err := connectDB(cfg)
	if err != nil {
		u.Logger.Fatalf("Migrator DB: %s", err.Error())
	}
	u.DBMigrator = db
	u.Logger.Infof("Migrator DB connected")
}

// CloseDB closes all open DB connections.
func (u *Utils) CloseDB() {
	if u.DBWriter != nil {
		u.DBWriter.Close()
	}
	if u.DBReader != nil {
		u.DBReader.Close()
	}
	if u.DBMigrator != nil {
		u.DBMigrator.Close()
	}
}

func connectDB(cfg DatabaseConfig) (*sql.DB, error) {
	dsn := buildDSN(cfg)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}
	return db, nil
}

func buildDSN(cfg DatabaseConfig) string {
	switch cfg.Driver {
	case "mysql":
		tz := cfg.Timezone
		if tz == "" {
			tz = "UTC"
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, tz)
	case "sqlite3":
		return cfg.Name
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	default:
		return cfg.Name
	}
}
