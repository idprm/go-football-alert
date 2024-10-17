package cmd

import (
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
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

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_NEWS_EXCHANGE, true, RMQ_NEWS_QUEUE)

		scrapingNews(db, rmq)
	},
}

var consumerTestBalanceCmd = &cobra.Command{
	Use:   "test_balance",
	Short: "Consumer Test Balance Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		subscriptionRepo := repository.NewSubscriptionRepository(db)
		subscriptionService := services.NewSubscriptionService(subscriptionRepo)

		h := handler.NewTestHandler(&logger.Logger{}, subscriptionService)

		h.TestBalance()

	},
}

var consumerTestChargeCmd = &cobra.Command{
	Use:   "test_charge",
	Short: "Consumer Test Charge Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		subscriptionRepo := repository.NewSubscriptionRepository(db)
		subscriptionService := services.NewSubscriptionService(subscriptionRepo)

		h := handler.NewTestHandler(&logger.Logger{}, subscriptionService)

		h.TestCharge()

	},
}

var consumerTestChargeFailedCmd = &cobra.Command{
	Use:   "test_charge_failed",
	Short: "Consumer Test Charge Failed Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		subscriptionRepo := repository.NewSubscriptionRepository(db)
		subscriptionService := services.NewSubscriptionService(subscriptionRepo)

		h := handler.NewTestHandler(&logger.Logger{}, subscriptionService)

		h.TestChargeFailed()

	},
}

var consumerTestUpdateFalseCmd = &cobra.Command{
	Use:   "test_update_false",
	Short: "Consumer Test Update False Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		subscriptionRepo := repository.NewSubscriptionRepository(db)
		subscriptionService := services.NewSubscriptionService(subscriptionRepo)

		h := handler.NewTestHandler(&logger.Logger{}, subscriptionService)

		h.TestUpdateToFalse()

	},
}
