package main

import (
	"embed"

	"/internal/cmd"
)

//go:embed all:configs
var configsFS embed.FS

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
	cmd.SetEmbedFS(configsFS, migrationsFS, webFS)
	cmd.SetVersionInfo(Version, GitCommit, GitBranch, GitState, GitSummary, BuildDate)
	cmd.Execute()
}
