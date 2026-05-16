package updater

// GitHubUpdater upgrades via GitHub releases.
type GitHubUpdater struct {
	Owner        string
	Repo         string
	Token        string
	AssetPattern string
}
