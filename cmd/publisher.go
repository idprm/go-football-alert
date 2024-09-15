package cmd

import (
	"encoding/json"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/gorm"
	loggerDb "gorm.io/gorm/logger"
)

var publisherRenewalCmd = &cobra.Command{
	Use:   "pub_renewal",
	Short: "Renewal CLI",
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
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_RENEWAL_EXCHANGE,
			true,
			RMQ_RENEWAL_QUEUE,
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_RENEWAL, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_RENEWAL,
						IsUnlocked: false,
					},
				)

				go func() {
					populateRenewal(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

var publisherRetryCmd = &cobra.Command{
	Use:   "pub_retry",
	Short: "Retry CLI",
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
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_RETRY_EXCHANGE,
			true,
			RMQ_RETRY_QUEUE,
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_RETRY, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_RETRY,
						IsUnlocked: false,
					},
				)

				go func() {
					populateRetry(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

var publisherPredictionCmd = &cobra.Command{
	Use:   "pub_prediction",
	Short: "Publisher Prediction CLI",
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_PREDICTION_EXCHANGE, true, RMQ_PREDICTION_QUEUE)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_PREDICTION, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_PREDICTION,
						IsUnlocked: false,
					},
				)

				go func() {
					populatePrediction(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

var publisherCreditCmd = &cobra.Command{
	Use:   "pub_credit_goal",
	Short: "Publisher Credit CLI",
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_CREDIT_EXCHANGE, true, RMQ_CREDIT_QUEUE)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_CREDIT_GOAL, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_CREDIT_GOAL,
						IsUnlocked: false,
					},
				)

				go func() {
					populateGoalCredit(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

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

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_SCRAPING, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_SCRAPING,
						IsUnlocked: false,
					},
				)

				go func() {
					scrapingFixturesAndNews(db)
				}()
			}

			if scheduleService.IsUnlocked(ACT_SCRAPING, timeNow) {
				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_SCRAPING,
						IsUnlocked: true,
					},
				)

			}

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherNewsCmd = &cobra.Command{
	Use:   "pub_news",
	Short: "Publisher News CLI",
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

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_NEWS, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_NEWS,
						IsUnlocked: false,
					},
				)

				go func() {
					populateNews(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

func populateRenewal(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Renewal()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_RENEWAL_EXCHANGE, RMQ_RENEWAL_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func populateRetry(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Retry()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_RETRY_EXCHANGE, RMQ_RETRY_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func populatePrediction(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Prediction()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_PREDICTION_EXCHANGE, RMQ_PREDICTION_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func populateGoalCredit(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Renewal()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_CREDIT_EXCHANGE, RMQ_CREDIT_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func populateNews(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.News()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_NEWS_EXCHANGE, RMQ_NEWS_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func scrapingFixturesAndNews(db *gorm.DB) {

	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	seasonRepo := repository.NewSeasonRepository(db)
	seasonService := services.NewSeasonService(seasonRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	homeRepo := repository.NewHomeRepository(db)
	homeService := services.NewHomeService(homeRepo)
	awayRepo := repository.NewAwayRepository(db)
	awayService := services.NewAwayService(awayRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	liveScoreRepo := repository.NewLiveScoreRepository(db)
	liveScoreService := services.NewLiveScoreService(liveScoreRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		leagueService,
		seasonService,
		fixtureService,
		homeService,
		awayService,
		teamService,
		liveScoreService,
		predictionService,
		newsService,
	)

	h.Fixtures()
	h.News()
}
