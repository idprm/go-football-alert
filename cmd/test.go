package cmd

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

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

var consumerTestMigrateSubCmd = &cobra.Command{
	Use:   "test_migrate_sub",
	Short: "Consumer Test Migrate Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "./logs/migrate_xxx.txt"

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 {
				msisdn := strings.TrimSpace(line)
				m, err := migrateHttp(msisdn)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println(string(m))
			}
		}

	},
}

func migrateHttp(msisdn string) ([]byte, error) {
	// http://165.22.122.64:9100/v1/migrate/sub?category=SMSALERTE_COMPETITION&code=SAC30&unique_code=PL&msisdn=
	q := url.Values{}
	q.Add("category", "SMSALERTE_COMPETITION")
	q.Add("code", "SAC30")
	q.Add("unique_code", "PL")
	q.Add("msisdn", msisdn)

	req, err := http.NewRequest("GET", "http://165.22.122.64:9100/v1/migrate/sub?"+q.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	tr := &http.Transport{
		MaxIdleConns:       30,
		IdleConnTimeout:    60 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   60 * time.Second,
		Transport: tr,
	}

	log.Println(req)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	log.Println(body)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return body, nil
}
