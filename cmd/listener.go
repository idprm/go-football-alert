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
		)

		/**
		 * Seeder
		 **/
		seederDB(db)

		/**
		 * Setup channel
		 */
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

	h := handler.NewIncomingHandler(
		rmq,
		logger,
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

	ussdRepo := repository.NewUssdRepository(db)
	ussdService := services.NewUssdService(ussdRepo)
	ussdHandler := handler.NewUssdHandler(ussdService)

	app.Post("/mo", h.MessageOriginated)

	v1 := app.Group(API_VERSION)
	leagues := v1.Group("leagues")
	leagues.Get("/", h.USSD)

	fixtures := v1.Group("fixtures")
	fixtures.Get("/", h.USSD)
	fixtures.Get("/:id", h.USSD)
	fixtures.Post("/", h.USSD)
	fixtures.Put("/:id", h.USSD)

	teams := v1.Group("teams")
	teams.Get("/", h.USSD)
	teams.Get("/:id", h.USSD)
	teams.Post("/", h.USSD)
	teams.Put("/:id", h.USSD)

	home := v1.Group("home")
	home.Get("/", h.USSD)
	home.Get("/:id", h.USSD)
	home.Post("/", h.USSD)
	home.Put("/:id", h.USSD)

	away := v1.Group("home")
	away.Get("/", h.USSD)
	away.Get("/:id", h.USSD)
	away.Post("/", h.USSD)
	away.Put("/:id", h.USSD)

	predictions := v1.Group("predictions")
	predictions.Get("/", h.USSD)
	predictions.Get("/:id", h.USSD)
	predictions.Post("/", h.USSD)
	predictions.Put("/:id", h.USSD)

	news := v1.Group("news")
	news.Get("/", h.USSD)
	news.Get("/:id", h.USSD)
	news.Post("/", h.USSD)
	news.Put("/:id", h.USSD)

	// callback
	ussd := v1.Group("ussd")
	ussd.Post("callback", ussdHandler.Callback)
	ussd.Post("event", ussdHandler.Event)

	// landing page
	p := v1.Group("p")
	p.Get("sub", h.Sub)
	p.Get("unsub", h.UnSub)

	return app
}

func seederDB(db *gorm.DB) {
	var country []entity.Country
	var service []entity.Service
	var content []entity.Content

	var countries = []entity.Country{
		{
			ID:       1,
			Name:     "MALI",
			Code:     "223",
			TimeZone: "GMT",
		},
		{
			ID:       2,
			Name:     "GUINEE",
			Code:     "224",
			TimeZone: "GMT",
		},
	}

	var services = []entity.Service{
		{
			ID:         1,
			CountryID:  1,
			Category:   "FB-ALERT",
			Name:       "FB 100",
			Code:       "FB100",
			Package:    "1",
			Price:      100,
			RenewalDay: 1,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS0001",
			UrlMT:      "http://10.106.0.3/",
			UserMT:     "admin",
			PassMT:     "admin",
		},
	}

	var contents = []entity.Content{
		{
			ServiceID: 1,
			Name:      ACT_FIRSTPUSH,
			Value:     "Test",
		},
		{
			ServiceID: 1,
			Name:      ACT_RENEWAL,
			Value:     "Test",
		},
	}

	if db.Find(&country).RowsAffected == 0 {
		for i, _ := range countries {
			db.Model(&entity.Country{}).Create(&countries[i])
		}
		log.Println("countries migrated")
	}

	if db.Find(&service).RowsAffected == 0 {
		for i, _ := range services {
			db.Model(&entity.Service{}).Create(&services[i])
		}
		log.Println("services migrated")
	}

	if db.Find(&content).RowsAffected == 0 {
		for i, _ := range contents {
			db.Model(&entity.Content{}).Create(&contents[i])
		}
		log.Println("contents migrated")
	}

}
