package utils

import (
	"database/sql"
	"embed"
	"html/template"
	"sync"
	"time"

	"scalland/pkg/log"
	"scalland/pkg/schemas"
)

// Utils is the singleton utility object holding shared resources.
type Utils struct {
	Logger       *log.Logger
	Config       *Config
	DBWriter     *sql.DB
	DBReader     *sql.DB
	DBMigrator   *sql.DB
	TimeZone     *time.Location
	Debug        bool
	TemplatesMap map[string]*template.Template
	Semver       *schemas.SemanticVersion
	MigrationsFS embed.FS
}

var (
	instance *Utils
	once     sync.Once
)

// GetUtils returns the singleton Utils instance.
func GetUtils() *Utils {
	once.Do(func() {
		instance = &Utils{
			Config: &Config{},
		}
	})
	return instance
}

// SetTimeZone sets the application timezone.
func (u *Utils) SetTimeZone(tz string) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		loc = time.UTC
	}
	u.TimeZone = loc
}
