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

/**
	RENEWAL at
	RETRY at
	SCRAPING_FIXTURES at
	SCRAPING_CREDITGOAL at
	SCRAPING_PREDICTION at
	CREDIT_GOAL at
	PREDICTION at
	FOLLOW_COMPETITION at
	FOLLOW_COMPETITION at
	FOLLOW_COMPETITION at
	FOLLOW_TEAM at
	FOLLOW_TEAM at
	FOLLOW_TEAM at
**/

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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_RENEWAL,
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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_RETRY,
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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_PREDICTION,
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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_CREDIT_GOAL,
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

var publisherScrapingFixturesCmd = &cobra.Command{
	Use:   "pub_scraping_fixtures",
	Short: "Publisher Scraping Fixture Service CLI",
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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_SCRAPING,
					},
				)

				go func() {
					scrapingFixtures(db)
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

var publisherScrapingPredictionCmd = &cobra.Command{
	Use:   "pub_scraping_prediction",
	Short: "Publisher Scraping Prediction Service CLI",
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

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_SCRAPING,
					},
				)

				go func() {
					scrapingPredictions(db)
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

var publisherScrapingNewsCmd = &cobra.Command{
	Use:   "pub_scraping_news",
	Short: "Publisher Scraping News Service CLI",
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
		timeDuration := time.Duration(90)

		for {

			go func() {
				scrapingNews(db, rmq)
			}()

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherSMSAlerteCmd = &cobra.Command{
	Use:   "pub_sms_alerte",
	Short: "Publisher SMS Alerte CLI",
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_SMS_ALERTE_EXCHANGE, true, RMQ_SMS_ALERTE_QUEUE)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_SMS_ALERTE, timeNow) {

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_SMS_ALERTE,
					},
				)

				go func() {
					populateSMSAlerte(db, rmq)
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

func populateSMSAlerte(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.FollowCompetition()

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

		rmq.IntegratePublish(RMQ_SMS_ALERTE_EXCHANGE, RMQ_SMS_ALERTE_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func scrapingLeagues(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	h.Leagues()
}

func scrapingTeams(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)
	h.Teams()
}

func scrapingFixtures(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	h.Fixtures()
}

func scrapingPredictions(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	h.Predictions()
}

func scrapingStandings(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	h.Standings()
}

func scrapingLineups(db *gorm.DB) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmqp.AMQP{},
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	h.Lineups()
}

func scrapingNews(db *gorm.DB, rmq rmqp.AMQP) {
	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)
	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)
	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)
	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)
	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)
	lineupRepo := repository.NewLineupRepository(db)
	lineupService := services.NewLineupService(lineupRepo)
	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	h := handler.NewScraperHandler(
		rmq,
		leagueService,
		teamService,
		fixtureService,
		predictionService,
		standingService,
		lineupService,
		newsService,
	)

	// maxifoot
	h.NewsMaxiFoot()
	// madeinfoot
	h.NewsMadeInFoot()
	// africatopsports
	h.NewsAfricaTopSports()
	// footmercato
	h.NewsFootMercato()
}
