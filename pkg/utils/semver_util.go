package utils

import "scalland/pkg/schemas"

// NewSemver creates a SemanticVersion from build info.
func NewSemver(version, commit, branch, state, summary, date string) *schemas.SemanticVersion {
	return &schemas.SemanticVersion{
		Version:    version,
		GitCommit:  commit,
		GitBranch:  branch,
		GitState:   state,
		GitSummary: summary,
		BuildDate:  date,
	}
}
