package cmd

import (
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/spf13/cobra"
)

var publisherScrapingCmd = &cobra.Command{
	Use:   "pub_scraping",
	Short: "Publisher Scraping Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		p := NewProcessor(db, rds, rmq, logger)
		p.Scraping()

	},
}
