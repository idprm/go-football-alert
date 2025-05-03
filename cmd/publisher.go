package cmd

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/providers/rabbit"
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RENEWAL_EXCHANGE, true, RMQ_RENEWAL_QUEUE)

		/**
		 * Looping schedule per 5 minute
		 */
		timeDuration := time.Duration(5)

		for {
			/**
			** Populate retry if queue message is zero or 0
			**/
			p := rabbit.NewRabbitMQ()

			q, err := p.Queue(RMQ_RENEWAL_QUEUE)
			if err != nil {
				log.Println(err)
			}

			var res *model.RabbitMQResponse
			json.Unmarshal(q, &res)

			// if queue is empty
			if !res.IsRunning() {
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
		 * Looping schedule per 2 hours
		 */
		timeDuration := time.Duration(2)

		for {

			go func() {
				populateRetry(db, rmq)
			}()

			time.Sleep(timeDuration * time.Hour)
		}
	},
}

var publisherRetryUnderpaymentCmd = &cobra.Command{
	Use:   "pub_retry_underpayment",
	Short: "Pub Retry Underpayment CLI",
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
			RMQ_RETRY_UP_EXCHANGE,
			true,
			RMQ_RETRY_UP_QUEUE,
		)

		/**
		 * Looping schedule per 3 hours
		 */
		timeDuration := time.Duration(3)

		for {

			go func() {
				populateRetryUnderpayment(db, rmq)
			}()

			time.Sleep(timeDuration * time.Hour)
		}
	},
}

var publisherReminder48HBeforeChargingCmd = &cobra.Command{
	Use:   "pub_reminder_48h_before_charging",
	Short: "Publisher Reminder 48H Before Charging CLI",
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
			RMQ_REMINDER_48H_BEFORE_CHARGING_EXCHANGE,
			true,
			RMQ_REMINDER_48H_BEFORE_CHARGING_QUEUE,
		)

		/**
		 * Looping schedule per 1 hour
		 */
		timeDuration := time.Duration(1)

		for {

			go func() {
				populateReminder48HBeforeCharging(db, rmq)
			}()

			time.Sleep(timeDuration * time.Hour)
		}
	},
}

var publisherReminderAfterTrialEndsCmd = &cobra.Command{
	Use:   "pub_reminder_after_trial_ends",
	Short: "Pub Reminder After Trial Ends CLI",
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
			RMQ_REMINDER_AFTER_TRIAL_ENDS_EXCHANGE,
			true,
			RMQ_REMINDER_AFTER_TRIAL_ENDS_QUEUE,
		)

		/**
		 * Looping schedule per 50 minute
		 */
		timeDuration := time.Duration(50)

		for {

			go func() {
				populateReminderAfterTrialEnds(db, rmq)
			}()

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherCreditGoalCmd = &cobra.Command{
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
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_CREDIT_GOAL_EXCHANGE,
			true,
			RMQ_CREDIT_GOAL_QUEUE,
		)

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
					populateCreditGoal(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

var publisherPredictWinCmd = &cobra.Command{
	Use:   "pub_predict_win",
	Short: "Publisher Predict Win CLI",
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
			RMQ_PREDICT_WIN_EXCHANGE,
			true,
			RMQ_PREDICT_WIN_QUEUE,
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_PREDICT_WIN, timeNow) {

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_PREDICT_WIN,
					},
				)

				go func() {
					populatePredictWin(db, rmq)
				}()
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
		timeDuration := time.Duration(25)

		for {

			go func() {
				scrapingNews(db, rmq)
			}()

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherScrapingMasterCmd = &cobra.Command{
	Use:   "pub_scraping_master",
	Short: "Publisher Scraping Master Service CLI",
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
		timeDuration := time.Duration(6)

		for {

			go func() {
				scrapingLeagues(db)
				scrapingTeams(db)
				scrapingStandings(db)
			}()

			time.Sleep(timeDuration * time.Hour)
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
		timeDuration := time.Duration(30)

		for {

			go func() {
				scrapingFixtures(db)
			}()

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherScrapingLiveMatchesCmd = &cobra.Command{
	Use:   "pub_scraping_livematches",
	Short: "Publisher Scraping LiveMatches Service CLI",
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

			go func() {
				scrapingLiveMatch(db)
			}()

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

			if scheduleService.IsUnlocked(ACT_PREDICTION, timeNow) {

				scheduleService.UpdateLocked(
					&entity.Schedule{
						Name: ACT_PREDICTION,
					},
				)

				go func() {
					scrapingPredictions(db)
				}()
			}

			if scheduleService.IsUnlocked(ACT_PREDICTION, timeNow) {
				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_PREDICTION,
						IsUnlocked: true,
					},
				)

			}

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

var publisherReportCmd = &cobra.Command{
	Use:   "pub_report",
	Short: "Publisher Report Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		// connect sqldb
		sqlDb, err := connectSqlDb()
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
			if timeNow == "23:58" || timeNow == "16:00" || timeNow == "08:00" || timeNow == "03:00" {
				go func() {
					populateReport(db, sqlDb)
				}()
			}
			time.Sleep(timeDuration * time.Minute)
		}

	},
}

func populateRenewal(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	if subscriptionService.IsPopulateRenewal() {

		subs := subscriptionService.Renewal()

		if len(*subs) > 0 {
			for _, s := range *subs {
				var sub entity.Subscription

				sub.ID = s.ID
				sub.ServiceID = s.ServiceID
				sub.Msisdn = s.Msisdn
				sub.Code = s.Code
				sub.LatestKeyword = s.LatestKeyword
				sub.LatestSubject = s.LatestSubject
				sub.CreatedAt = s.CreatedAt

				json, _ := json.Marshal(sub)

				rmq.IntegratePublish(RMQ_RENEWAL_EXCHANGE, RMQ_RENEWAL_QUEUE, RMQ_DATA_TYPE, "", string(json))

				time.Sleep(100 * time.Microsecond)
			}
		}
	}

}

func populateRetry(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	if subscriptionService.IsPopulateRetry() {

		subs := subscriptionService.Retry()

		if len(*subs) > 0 {
			for _, s := range *subs {
				var sub entity.Subscription

				sub.ID = s.ID
				sub.ServiceID = s.ServiceID
				sub.Msisdn = s.Msisdn
				sub.Code = s.Code
				sub.LatestKeyword = s.LatestKeyword
				sub.LatestSubject = s.LatestSubject
				sub.CreatedAt = s.CreatedAt

				json, _ := json.Marshal(sub)

				rmq.IntegratePublish(RMQ_RETRY_EXCHANGE, RMQ_RETRY_QUEUE, RMQ_DATA_TYPE, "", string(json))

				time.Sleep(100 * time.Microsecond)
			}
		}
	}

}

func populateRetryUnderpayment(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	if subscriptionService.IsPopulateRetryUnderpayment() {

		subs := subscriptionService.RetryUnderpayment()

		if len(*subs) > 0 {
			for _, s := range *subs {
				var sub entity.Subscription

				sub.ID = s.ID
				sub.ServiceID = s.ServiceID
				sub.Msisdn = s.Msisdn
				sub.Code = s.Code
				sub.LatestKeyword = s.LatestKeyword
				sub.LatestSubject = s.LatestSubject
				sub.CreatedAt = s.CreatedAt
				sub.TotalUnderpayment = s.TotalUnderpayment

				json, _ := json.Marshal(sub)

				rmq.IntegratePublish(RMQ_RETRY_UP_EXCHANGE, RMQ_RETRY_UP_QUEUE, RMQ_DATA_TYPE, "", string(json))

				time.Sleep(100 * time.Microsecond)
			}
		}
	}
}

func populatePredictWin(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.PredictWin()

	if len(*subs) > 0 {
		for _, s := range *subs {
			var sub entity.Subscription

			sub.ID = s.ID
			sub.ServiceID = s.ServiceID
			sub.Msisdn = s.Msisdn
			sub.Code = s.Code
			sub.LatestKeyword = s.LatestKeyword
			sub.LatestSubject = s.LatestSubject
			sub.CreatedAt = s.CreatedAt

			json, _ := json.Marshal(sub)

			rmq.IntegratePublish(RMQ_PREDICT_WIN_EXCHANGE, RMQ_PREDICT_WIN_QUEUE, RMQ_DATA_TYPE, "", string(json))

			time.Sleep(100 * time.Microsecond)
		}
	}
}

func populateCreditGoal(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Renewal()

	if len(*subs) > 0 {
		for _, s := range *subs {
			var sub entity.Subscription

			sub.ID = s.ID
			sub.ServiceID = s.ServiceID
			sub.Msisdn = s.Msisdn
			sub.Code = s.Code
			sub.LatestKeyword = s.LatestKeyword
			sub.LatestSubject = s.LatestSubject
			sub.CreatedAt = s.CreatedAt

			json, _ := json.Marshal(sub)

			rmq.IntegratePublish(RMQ_CREDIT_GOAL_EXCHANGE, RMQ_CREDIT_GOAL_QUEUE, RMQ_DATA_TYPE, "", string(json))

			time.Sleep(100 * time.Microsecond)
		}
	}
}

func populateReminder48HBeforeCharging(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	if subscriptionService.IsPopulateReminder48HBeforeCharging() {
		subs := subscriptionService.Reminder48HBeforeCharging()
		if len(*subs) > 0 {
			for _, s := range *subs {
				var sub entity.Subscription

				sub.ID = s.ID
				sub.ServiceID = s.ServiceID
				sub.Msisdn = s.Msisdn
				sub.Code = s.Code
				sub.LatestKeyword = s.LatestKeyword
				sub.LatestSubject = s.LatestSubject
				sub.FreeAt = s.FreeAt
				sub.RenewalAt = s.RenewalAt
				sub.CreatedAt = s.CreatedAt

				json, _ := json.Marshal(sub)

				rmq.IntegratePublish(RMQ_REMINDER_48H_BEFORE_CHARGING_EXCHANGE, RMQ_REMINDER_48H_BEFORE_CHARGING_QUEUE, RMQ_DATA_TYPE, "", string(json))

				time.Sleep(100 * time.Microsecond)
			}
		}
	}

}

func populateReminderAfterTrialEnds(db *gorm.DB, rmq rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	if subscriptionService.IsPopulateReminderAfterTrialEnds() {
		subs := subscriptionService.ReminderAfterTrialEnds()

		if len(*subs) > 0 {
			for _, s := range *subs {
				var sub entity.Subscription

				sub.ID = s.ID
				sub.ServiceID = s.ServiceID
				sub.Msisdn = s.Msisdn
				sub.Code = s.Code
				sub.LatestKeyword = s.LatestKeyword
				sub.LatestSubject = s.LatestSubject
				sub.FreeAt = s.FreeAt
				sub.RenewalAt = s.RenewalAt
				sub.CreatedAt = s.CreatedAt

				json, _ := json.Marshal(sub)

				rmq.IntegratePublish(RMQ_REMINDER_AFTER_TRIAL_ENDS_EXCHANGE, RMQ_REMINDER_AFTER_TRIAL_ENDS_QUEUE, RMQ_DATA_TYPE, "", string(json))

				time.Sleep(100 * time.Microsecond)
			}
		}
	}

}

func populateReport(db *gorm.DB, sqlDb *sql.DB) {

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)
	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	summaryDashboardRepo := repository.NewSummaryDashboardRepository(db, sqlDb)
	summaryDashboardService := services.NewSummaryDashboardService(summaryDashboardRepo)
	summaryRevenueRepo := repository.NewSummaryRevenueRepository(db, sqlDb)
	summaryRevenueService := services.NewSummaryRevenueService(summaryRevenueRepo)
	summaryTotalDailyRepo := repository.NewSummaryTotalDailyRepository(db, sqlDb)
	summaryTotalDailyService := services.NewSummaryTotalDailyService(summaryTotalDailyRepo)

	h := handler.NewReportHandler(
		serviceService,
		subscriptionService,
		transactionService,
		summaryDashboardService,
		summaryRevenueService,
		summaryTotalDailyService,
	)

	// based trans
	h.PopulateRevenue()

	// based sub
	h.GetTotalActiveSub()
	h.GetTotalRevenue()

	h.PopulateTotalDaily()
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
		newsService,
	)

	h.Fixtures()
}

func scrapingLiveMatch(db *gorm.DB) {
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
		newsService,
	)

	h.LiveMatches()
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
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
	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)
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
		livematchService,
		newsService,
	)

	// maxifoot
	// h.NewsMaxiFoot()
	// madeinfoot
	// h.NewsMadeInFoot()
	// africatopsports
	// h.NewsAfricaTopSports()
	// footmercato
	// h.NewsFootMercato()
	// rmcsport
	// h.NewsRmcSport()

	// mobimiumnews
	h.MobimiumNews()
}
