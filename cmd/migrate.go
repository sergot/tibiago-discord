package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/sergot/tibiago/src/models"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Aliases: []string{"db", "migration"},
	Short:   "Run db migrations",
	Long:    `Synchronize Ent with your db.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add colors
		fmt.Println("Applying database migrations...")

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Fatalf("failed connecting to database: %v", err)
		}
		defer client.Close()

		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		fmt.Println("DONE")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
