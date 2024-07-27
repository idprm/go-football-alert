package cmd

import "github.com/spf13/cobra"

var scraperCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scraper Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db

	},
}
