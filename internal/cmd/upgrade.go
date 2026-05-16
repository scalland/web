package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"scalland/internal/updater"
	"scalland/pkg/utils"
)

var upgradeChannel string

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Self-upgrade to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := utils.GetUtils()
		up := updater.NewUpdater(u.Config.Upgrade, Version)

		if upgradeChannel != "" {
			return up.UpgradeViaChannel(upgradeChannel)
		}
		return up.AutoUpgrade()
	},
}

var upgradeListCmd = &cobra.Command{
	Use:   "channels",
	Short: "List available upgrade channels",
	Run: func(cmd *cobra.Command, args []string) {
		u := utils.GetUtils()
		up := updater.NewUpdater(u.Config.Upgrade, Version)
		channels := up.ListAvailableChannels()
		for _, ch := range channels {
			fmt.Printf("  %s\n", ch)
		}
	},
}

func init() {
	upgradeCmd.Flags().StringVar(&upgradeChannel, "channel", "", "Upgrade channel (github, s3, http, local)")
	upgradeCmd.AddCommand(upgradeListCmd)
	rootCmd.AddCommand(upgradeCmd)
}
