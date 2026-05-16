package schemas

// SemanticVersion holds build info injected via ldflags.
type SemanticVersion struct {
	GitBranch  string
	GitState   string
	GitSummary string
	BuildDate  string
	Version    string
	GitCommit  string
}
