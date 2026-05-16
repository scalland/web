package updater

import "fmt"

// Updater manages self-upgrade from multiple channels.
type Updater struct {
	Config         interface{} // UpgradeConfig
	CurrentVersion string
}

// NewUpdater creates a new Updater.
func NewUpdater(config interface{}, currentVersion string) *Updater {
	return &Updater{Config: config, CurrentVersion: currentVersion}
}

// AutoUpgrade upgrades from the default channel.
func (u *Updater) AutoUpgrade() error {
	fmt.Println("Auto-upgrade not yet configured")
	return nil
}

// UpgradeViaChannel upgrades from a specific channel.
func (u *Updater) UpgradeViaChannel(channel string) error {
	fmt.Printf("Upgrading via channel: %s\n", channel)
	return nil
}

// ListAvailableChannels returns configured channels.
func (u *Updater) ListAvailableChannels() []string {
	return []string{"github", "s3", "http", "local"}
}
