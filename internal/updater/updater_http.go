package updater

// HTTPUpdater upgrades via HTTP download.
type HTTPUpdater struct {
	BaseURL     string
	ManifestURL string
}
