package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	loggerFiber "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
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
			&entity.Team{},
			&entity.LeagueTeam{},
			&entity.Fixture{},
			&entity.Prediction{},
			&entity.Lineup{},
			&entity.Standing{},
			&entity.News{},
			&entity.NewsLeagues{},
			&entity.NewsTeams{},
			&entity.Schedule{},
			&entity.Service{},
			&entity.Content{},
			&entity.Subscription{},
			&entity.SubscriptionCreditGoal{},
			&entity.SubscriptionPredict{},
			&entity.SubscriptionFollowLeague{},
			&entity.SubscriptionFollowTeam{},
			&entity.SMSAlerte{},
			&entity.Transaction{},
			&entity.History{},
			&entity.Betting{},
			&entity.Summary{},
			&entity.MT{},
		)

		/**
		 * Seeder
		 **/
		seederDB(db)

		/**
		 * Setup channel
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_USSD_EXCHANGE, true, RMQ_USSD_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_SMS_EXCHANGE, true, RMQ_SMS_QUEUE)

		r := routeUrlListener(db, rds, rmq, logger)
		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeUrlListener(db *gorm.DB, rds *redis.Client, rmq rmqp.AMQP, logger *logger.Logger) *fiber.App {

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	engine := html.New(path+"/views", ".html")

	r := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	/**
	 * Initialize default config
	 */
	r.Use(cors.New())

	/**
	 * Access log on browser
	 */
	r.Use(PATH_LOG, filesystem.New(
		filesystem.Config{
			Root:         http.Dir(PATH_LOG),
			Browse:       true,
			Index:        "index.html",
			NotFoundFile: "404.html",
		},
	))

	f, err := logAccess()
	if err != nil {
		log.Println(err)
	}

	r.Use(loggerFiber.New(
		loggerFiber.Config{
			Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${url} | ${referer} | ${error}\n",
			TimeFormat: "02-Jan-2006 15:04:05",
			TimeZone:   APP_TZ,
			Output:     f,
		},
	))

	r.Static(PATH_STATIC, path+"/public")

	r.Use(cors.New())

	summaryRepo := repository.NewSummaryRepository(db)
	summaryService := services.NewSummaryService(summaryRepo)

	menuRepo := repository.NewMenuRepository(db, rds)
	menuService := services.NewMenuService(menuRepo)

	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)

	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)

	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)

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

	subscriptionFollowLeagueRepo := repository.NewSubscriptionFollowLeagueRepository(db)
	subscriptionFollowLeagueService := services.NewSubscriptionFollowLeagueService(subscriptionFollowLeagueRepo)

	subscriptionFollowTeamRepo := repository.NewSubscriptionFollowTeamRepository(db)
	subscriptionFollowTeamService := services.NewSubscriptionFollowTeamService(subscriptionFollowTeamRepo)

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)

	historyRepo := repository.NewHistoryRepository(db)
	historyService := services.NewHistoryService(historyRepo)

	bettingRepo := repository.NewBettingRepository(db)
	bettingService := services.NewBettingService(bettingRepo)

	ussdRepo := repository.NewUssdRepository(db, rds)
	ussdService := services.NewUssdService(ussdRepo)

	mtRepo := repository.NewMTRepository(db)
	mtService := services.NewMTService(mtRepo)

	smsAlerteRepo := repository.NewSMSAlerteRespository(db)
	smsAlerteService := services.NewSMSAlerteService(smsAlerteRepo)

	h := handler.NewIncomingHandler(
		rds,
		rmq,
		logger,
		menuService,
		ussdService,
		leagueService,
		teamService,
		fixtureService,
		livescoreService,
		predictionService,
		newsService,
		scheduleService,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		bettingService,
	)

	leagueHandler := handler.NewLeagueHandler(leagueService)
	teamHandler := handler.NewTeamHandler(teamService)
	fixtureHandler := handler.NewFixtureHandler(fixtureService)

	predictionHandler := handler.NewPredictionHandler(
		rmq,
		logger,
		&entity.Subscription{},
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		predictionService,
	)

	newsHandler := handler.NewNewsHandler(
		rmq,
		leagueService,
		teamService,
		newsService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		&entity.News{},
	)

	dcbHandler := handler.NewDCBHandler(
		summaryService,
		menuService,
		ussdService,
		scheduleService,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		historyService,
		mtService,
		smsAlerteService,
	)

	r.Get("/mo", h.MessageOriginated)
	r.Post("/sub", h.CreateSub)
	r.Post("/otp", h.VerifySub)

	lp := r.Group("p")
	lp.Get("/:service", h.LandingPage)

	// v1
	v1 := r.Group(API_VERSION)
	// fiture
	fiture := r.Group("fiture")
	// direct carrier billing
	dcb := r.Group("dcb")
	// test
	test := r.Group("test")

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
	fixtures.Get("/:id", fixtureHandler.Get)
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
	news.Get("/:id", newsHandler.GetById)
	news.Post("/", newsHandler.Save)
	news.Put("/:id", newsHandler.Delete)

	// callback
	ussd := v1.Group("ussd")
	// main menu in ussd
	ussd.Get("/", h.Main)
	ussd.Get("/q", h.Menu)
	ussd.Get("/q/:name", h.Detail)
	ussd.Get("/confirm", h.Confirm)
	ussd.Get("/buy", h.Buy)

	// landing page
	p := v1.Group("p")
	p.Get("sub", h.Sub)
	p.Get("unsub", h.UnSub)

	// summaries
	summaries := dcb.Group("summaries")
	summaries.Get("/", dcbHandler.GetAllSummaryPaginate)

	// menus
	menus := dcb.Group("menus")
	menus.Get("/", dcbHandler.GetAllMenuPaginate)

	// schedules
	schedules := dcb.Group("schedules")
	schedules.Get("/", dcbHandler.GetAllSchedulePaginate)

	// services
	services := dcb.Group("services")
	services.Get("/", dcbHandler.GetAllServicePaginate)

	// contents
	contents := dcb.Group("contents")
	contents.Get("/", dcbHandler.GetAllContentPaginate)

	// subscriptions
	subscriptions := dcb.Group("subscriptions")
	subscriptions.Get("/", dcbHandler.GetAllSubscriptionPaginate)

	// transactions
	transactions := dcb.Group("transactions")
	transactions.Get("/", dcbHandler.GetAllSubscriptionPaginate)

	// histories
	histories := dcb.Group("histories")
	histories.Get("/", dcbHandler.GetAllHistoryPaginate)

	// message terminated
	mt := dcb.Group("mt")
	mt.Get("/", dcbHandler.GetAllMTPaginate)

	// sms alertes
	smsalerte := fiture.Group("smsalertes")
	smsalerte.Get("/", dcbHandler.GetAllSMSAlertePaginate)

	test.Post("/balance", h.TestBalance)
	test.Post("/charge", h.TestCharge)
	test.Post("/charge-failed", h.TestChargeFailed)

	return r
}

func logAccess() (*os.File, error) {
	file, err := os.OpenFile(PATH_LOG+"/access/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return nil, err
	}
	return file, nil
}
