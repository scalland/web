package main

import (
	"embed"

	"scalland/internal/cmd"
	"scalland/pkg/log"

	_ "github.com/go-sql-driver/mysql"
)



//go:embed all:migrations
var migrationsFS embed.FS

//go:embed all:web
var webFS embed.FS

// Build info — injected via ldflags
var (
	Version    = "dev"
	GitCommit  = "none"
	GitBranch  = "unknown"
	GitState   = "unknown"
	GitSummary = "none"
	BuildDate  = "unknown"
)

func main() {
	l := log.NewLogger()
	cmd.SetEmbedFS(migrationsFS, webFS)
	cmd.SetVersionInfo(Version, GitCommit, GitBranch, GitState, GitSummary, BuildDate)
	l.Debugf("Starting Scalland...")
	l.Debugf("Version: %s\nGitCommit: %s\nGitBranch: %s\nGitState: %s\nGitSummary: %s\nBuildDate: %s\n", Version, GitCommit, GitBranch, GitState, GitSummary, BuildDate)
	cmd.Execute()
}
