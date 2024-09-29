package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
)

var consumerTestLeagueCmd = &cobra.Command{
	Use:   "test_league",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingLeagues(db)
	},
}

var consumerTestTeamCmd = &cobra.Command{
	Use:   "test_team",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingTeams(db)
	},
}

var consumerTestFixtureCmd = &cobra.Command{
	Use:   "test_fixture",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingFixtures(db)
	},
}

var consumerTestPredictionCmd = &cobra.Command{
	Use:   "test_prediction",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingPredictions(db)
	},
}

var consumerTestStandingCmd = &cobra.Command{
	Use:   "test_standing",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingStandings(db)
	},
}

var consumerTestLineupCmd = &cobra.Command{
	Use:   "test_lineup",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingLineups(db)
	},
}

var consumerTestNewsCmd = &cobra.Command{
	Use:   "test_news",
	Short: "Consumer Test Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		scrapingNews(db, rmqp.AMQP{})
	},
}
