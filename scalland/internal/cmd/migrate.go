package cmd

import (
	"github.com/spf13/cobra"

	"/pkg/utils"
)

var createDB bool

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		u := utils.GetUtils()
		u.ConnectMigrator()
		defer u.CloseDB()
		return u.MigrateDB(createDB)
	},
}

func init() {
	migrateCmd.Flags().BoolVar(&createDB, "create-db", false, "Create the database if it doesn't exist")
	rootCmd.AddCommand(migrateCmd)
}
