package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version    = "dev"
	GitCommit  = "none"
	GitBranch  = "unknown"
	GitState   = "unknown"
	GitSummary = "none"
	BuildDate  = "unknown"
)

// SetVersionInfo sets version info from main.go ldflags.
func SetVersionInfo(version, commit, branch, state, summary, date string) {
	Version = version
	GitCommit = commit
	GitBranch = branch
	GitState = state
	GitSummary = summary
	BuildDate = date
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(" %s\n", Version)
		fmt.Printf("  Commit:  %s\n", GitCommit)
		fmt.Printf("  Branch:  %s\n", GitBranch)
		fmt.Printf("  State:   %s\n", GitState)
		fmt.Printf("  Summary: %s\n", GitSummary)
		fmt.Printf("  Built:   %s\n", BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
