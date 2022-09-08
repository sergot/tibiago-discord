package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/sergot/tibiago/src/models"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:     "seed",
	Aliases: []string{},
	Short:   "Run db migrations",
	Long:    `Synchronize Ent with your db.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add colors
		fmt.Println("Seeding the database...")

		client, err := models.ConnectDatabase()
		if err != nil {
			log.Fatalf("failed connecting to database: %v", err)
		}
		defer client.Close()

		// TODO: provide dev db dump
		_, err = client.Boss.Create().
			SetName("feru").
			Save(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
