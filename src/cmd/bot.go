package cmd

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sergot/tibiago/src/bot"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runBot()
	},
}

func init() {
	rootCmd.AddCommand(botCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runBot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	bot.Connect("db")
}
