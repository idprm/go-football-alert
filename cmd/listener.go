package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

		// DEBUG ON CONSOLE
		db.Logger = logger.Default.LogMode(logger.Info)

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
			&entity.Reward{},
		)

		r := routeUrlListener(db)
		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeUrlListener(db *gorm.DB) *fiber.App {
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

	h := handler.NewListenerHandler(
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

	// landing page
	p := v1.Group("p")
	p.Get("sub", h.Sub)
	p.Get("unsub", h.UnSub)

	// callback
	ussd := v1.Group("ussd")
	ussd.Post("callback", ussdHandler.Callback)
	ussd.Post("event", ussdHandler.Event)

	return app
}
