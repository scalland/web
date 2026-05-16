package cmd

import (
	"github.com/spf13/cobra"

	"scalland/pkg/utils"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := utils.GetUtils()
		u.ConnectWriter()
		defer u.CloseDB()

		u.Logger.Infof("Seeding database...")
		// Add your seed logic here
		u.Logger.Infof("Database seeded successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
