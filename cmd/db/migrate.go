package db

import (
	"github.com/d3ta-go/project-template/interface/cmd-apps/database"
	"github.com/spf13/cobra"
)

// migrateCmd represents the db migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Shows the db migrate command.",
	Long:  `Shows the db migrate command.`,
	Run: func(cmd *cobra.Command, args []string) {

		database.RunDBMigrate()
	},
}

func init() {
	DBCmd.AddCommand(migrateCmd)
}