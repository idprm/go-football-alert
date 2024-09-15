package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/gorm"
	loggerDb "gorm.io/gorm/logger"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Listener Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
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
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * setup log
		 */
		logger := logger.NewLogger()

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		// TODO: Add migrations
		db.AutoMigrate(
			&entity.Menu{},
			&entity.Ussd{},
			&entity.League{},
			&entity.Season{},
			&entity.Team{},
			&entity.Fixture{},
			&entity.Home{},
			&entity.Away{},
			&entity.Prediction{},
			&entity.News{},
			&entity.Country{},
			&entity.Schedule{},
			&entity.Service{},
			&entity.Content{},
			&entity.Subscription{},
			&entity.Transaction{},
			&entity.History{},
			&entity.Reward{},
			&entity.Summary{},
		)

		/**
		 * Seeder
		 **/
		seederDB(db)

		/**
		 * Setup channel
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_USSD_EXCHANGE, true, RMQ_USSD_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MO_EXCHANGE, true, RMQ_MO_QUEUE)

		r := routeUrlListener(db, rds, rmq, logger)
		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeUrlListener(db *gorm.DB, rds *redis.Client, rmq rmqp.AMQP, logger *logger.Logger) *fiber.App {
	app := fiber.New()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	app.Static(PATH_STATIC, path+"/public")

	app.Use(cors.New())

	menuRepo := repository.NewMenuRepository(db)
	menuService := services.NewMenuService(menuRepo)

	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)

	seasonRepo := repository.NewSeasonRepository(db)
	seasonService := services.NewSeasonService(seasonRepo)

	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)

	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)

	homeRepo := repository.NewHomeRepository(db)
	homeService := services.NewHomeService(homeRepo)

	awayRepo := repository.NewAwayRepository(db)
	awayService := services.NewAwayService(awayRepo)

	livescoreRepo := repository.NewLiveScoreRepository(db)
	livescoreService := services.NewLiveScoreService(livescoreRepo)

	predictionRepo := repository.NewPredictionRepository(db)
	predictionService := services.NewPredictionService(predictionRepo)

	newsRepo := repository.NewNewsRepository(db)
	newsService := services.NewNewsService(newsRepo)

	scheduleRepo := repository.NewScheduleRepository(db)
	scheduleService := services.NewScheduleService(scheduleRepo)

	serviceRepo := repository.NewServiceRepository(db)
	serviceService := services.NewServiceService(serviceRepo)

	contentRepo := repository.NewContentRepository(db)
	contentService := services.NewContentService(contentRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)

	rewardRepo := repository.NewRewardRepository(db)
	rewardService := services.NewRewardService(rewardRepo)

	ussdRepo := repository.NewUssdRepository(db)
	ussdService := services.NewUssdService(ussdRepo)

	h := handler.NewIncomingHandler(
		rmq,
		logger,
		menuService,
		ussdService,
		leagueService,
		seasonService,
		teamService,
		fixtureService,
		homeService,
		awayService,
		livescoreService,
		predictionService,
		newsService,
		scheduleService,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		rewardService,
	)

	leagueHandler := handler.NewLeagueHandler(leagueService)
	teamHandler := handler.NewTeamHandler(teamService)
	fixtureHandler := handler.NewFixtureHandler(fixtureService)
	predictionHandler := handler.NewPredictionHandler(predictionService)
	newsHandler := handler.NewNewsHandler(newsService)

	app.Post("/mo", h.MessageOriginated)

	v1 := app.Group(API_VERSION)
	leagues := v1.Group("leagues")
	leagues.Get("/", leagueHandler.GetAllPaginate)
	leagues.Get("/:slug", leagueHandler.GetBySlug)
	leagues.Post("/", leagueHandler.Save)
	leagues.Put("/:id", leagueHandler.Update)
	leagues.Delete("/:id", leagueHandler.Delete)

	teams := v1.Group("teams")
	teams.Get("/", teamHandler.GetAllPaginate)
	teams.Get("/:slug", teamHandler.GetBySlug)
	teams.Post("/", teamHandler.Save)
	teams.Put("/:id", teamHandler.Update)
	teams.Delete("/:id", teamHandler.Delete)

	fixtures := v1.Group("fixtures")
	fixtures.Get("/", fixtureHandler.GetAllPaginate)
	fixtures.Get("/:slug", fixtureHandler.GetBySlug)
	fixtures.Post("/", fixtureHandler.Save)
	fixtures.Put("/:id", fixtureHandler.Update)
	fixtures.Delete("/:id", fixtureHandler.Delete)

	predictions := v1.Group("predictions")
	predictions.Get("/", predictionHandler.GetAllPaginate)
	predictions.Get("/:slug", predictionHandler.GetBySlug)
	predictions.Post("/", predictionHandler.Save)
	predictions.Put("/:id", predictionHandler.Update)
	predictions.Delete("/:id", predictionHandler.Delete)

	news := v1.Group("news")
	news.Get("/", newsHandler.GetAllPaginate)
	news.Get("/:slug", newsHandler.GetBySlug)
	news.Post("/", newsHandler.Save)
	news.Put("/:id", newsHandler.Delete)

	// callback
	ussd := v1.Group("ussd")
	ussd.Post("callback", h.Callback)
	// ussd.Post("event", ussdHandler.Event)

	// landing page
	p := v1.Group("p")
	p.Get("sub", h.Sub)
	p.Get("unsub", h.UnSub)

	return app
}
