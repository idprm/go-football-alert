package cmd

import "github.com/spf13/cobra"

var scraperCmd = &cobra.Command{
	Use:   "scraper",
	Short: "Scraper Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		p := NewProcessor(db)
		p.Scraping()

	},
}
