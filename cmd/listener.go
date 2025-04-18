package cmd

import (
	"database/sql"
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
			&entity.LiveMatch{},
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
			&entity.MT{},
			&entity.Pronostic{},
			&entity.SMSProno{},
			&entity.SMSActu{},
			&entity.MO{},
			&entity.SummaryRevenue{},
			&entity.Country{},
			&entity.User{},
			&entity.SummaryDashboard{},
			&entity.SummaryRevenue{},
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_SMS_PRONO_EXCHANGE, true, RMQ_SMS_PRONO_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MO_EXCHANGE, true, RMQ_MO_QUEUE)

		r := routeUrlListener(db, &sql.DB{}, rds, rmq, logger)
		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeUrlListener(db *gorm.DB, sqlDb *sql.DB, rds *redis.Client, rmq rmqp.AMQP, logger *logger.Logger) *fiber.App {

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

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	menuRepo := repository.NewMenuRepository(db, rds)
	menuService := services.NewMenuService(menuRepo)

	leagueRepo := repository.NewLeagueRepository(db)
	leagueService := services.NewLeagueService(leagueRepo)

	teamRepo := repository.NewTeamRepository(db)
	teamService := services.NewTeamService(teamRepo)

	fixtureRepo := repository.NewFixtureRepository(db)
	fixtureService := services.NewFixtureService(fixtureRepo)

	livematchRepo := repository.NewLiveMatchRepository(db)
	livematchService := services.NewLiveMatchService(livematchRepo)

	livescoreRepo := repository.NewLiveScoreRepository(db)
	livescoreService := services.NewLiveScoreService(livescoreRepo)

	standingRepo := repository.NewStandingRepository(db)
	standingService := services.NewStandingService(standingRepo)

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

	subscriptionCreditGoalRepo := repository.NewSubscriptionCreditGoalRepository(db)
	subscriptionCreditGoalService := services.NewSubscriptionCreditGoalService(subscriptionCreditGoalRepo)

	subscriptionPredictWinRepo := repository.NewSubscriptionPredictWinRepository(db)
	subscriptionPredictWinService := services.NewSubscriptionPredictWinService(subscriptionPredictWinRepo)

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

	moRepo := repository.NewMORepository(db)
	moService := services.NewMOService(moRepo)

	mtRepo := repository.NewMTRepository(db)
	mtService := services.NewMTService(mtRepo)

	smsAlerteRepo := repository.NewSMSAlerteRespository(db)
	smsAlerteService := services.NewSMSAlerteService(smsAlerteRepo)

	pronosticRepo := repository.NewPronosticRepository(db)
	pronosticService := services.NewPronosticService(pronosticRepo)

	summaryDashboardRepo := repository.NewSummaryDashboardRepository(db, sqlDb)
	summaryDashboardService := services.NewSummaryDashboardService(summaryDashboardRepo)

	summaryRevenueRepo := repository.NewSummaryRevenueRepository(db, sqlDb)
	summaryRevenueService := services.NewSummaryRevenueService(summaryRevenueRepo)

	h := handler.NewIncomingHandler(
		rds,
		rmq,
		logger,
		menuService,
		ussdService,
		leagueService,
		teamService,
		fixtureService,
		livematchService,
		livescoreService,
		standingService,
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
		subscriptionService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		&entity.News{},
	)

	dcbHandler := handler.NewDCBHandler(
		rmq,
		userService,
		leagueService,
		teamService,
		menuService,
		ussdService,
		scheduleService,
		serviceService,
		contentService,
		subscriptionService,
		subscriptionCreditGoalService,
		subscriptionPredictWinService,
		subscriptionFollowLeagueService,
		subscriptionFollowTeamService,
		transactionService,
		historyService,
		moService,
		mtService,
		smsAlerteService,
		pronosticService,
		summaryDashboardService,
		summaryRevenueService,
	)

	r.Get("/mo", h.MessageOriginated)
	r.Post("/sub", h.CreateSub)
	r.Post("/otp", h.VerifySub)

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
	leagues.Put("/", leagueHandler.Update)
	leagues.Delete("/", leagueHandler.Delete)

	teams := v1.Group("teams")
	teams.Get("/", teamHandler.GetAllPaginate)
	teams.Get("/:slug", teamHandler.GetBySlug)
	teams.Put("/", teamHandler.Update)
	teams.Put("/toggle", teamHandler.Update)
	teams.Delete("/", teamHandler.Delete)

	fixtures := v1.Group("fixtures")
	fixtures.Get("/", fixtureHandler.GetAllPaginate)
	fixtures.Get("/current", fixtureHandler.GetAllCurrent)
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

	// pronostics
	pronostics := v1.Group("pronostics")
	pronostics.Get("/", dcbHandler.GetAllPronosticPaginate)
	pronostics.Post("/", dcbHandler.SavePronostic)

	// callback
	ussd := v1.Group("ussd")
	// main menu in ussd
	ussd.Get("/", h.Main)
	ussd.Get("/q", h.Menu)
	ussd.Get("/q/:name", h.Detail)
	ussd.Get("/confirm", h.Confirm)
	ussd.Get("/confirm-stop", h.ConfirmStop)
	ussd.Get("/buy", h.Buy)
	ussd.Get("/stop", h.Stop)

	// landing page
	p := v1.Group("p")
	p.Get("/alertesms", h.LPAlerteSMS)
	p.Get("/alertesms/:code/sub", h.Sub)
	p.Get("/alertesms/:code/unsub", h.UnSub)

	// auth
	auth := v1.Group("auth")
	auth.Post("/login", dcbHandler.Login)

	// summaries
	summaries := dcb.Group("summaries")
	summaries.Get("/dashboard", dcbHandler.GetAllSummaryDashboardPaginate)
	summaries.Get("/revenue", dcbHandler.GetAllSummaryRevenuePaginate)

	// charts
	charts := dcb.Group("charts")
	charts.Get("/revenue", dcbHandler.GetAllChartRevenue)

	// badges
	badges := dcb.Group("badges")
	badges.Get("/dashboard", dcbHandler.GetBadgeDashboard)

	// menus
	menus := dcb.Group("menus")
	menus.Get("/", dcbHandler.GetAllMenuPaginate)
	menus.Post("/", dcbHandler.SaveMenu)

	// schedules
	schedules := dcb.Group("schedules")
	schedules.Get("/", dcbHandler.GetAllSchedulePaginate)

	// services
	services := dcb.Group("services")
	services.Get("/", dcbHandler.GetAllServicePaginate)
	services.Post("/", dcbHandler.SaveService)

	// contents
	contents := dcb.Group("contents")
	contents.Get("/", dcbHandler.GetAllContentPaginate)
	contents.Post("/", dcbHandler.SaveContent)

	// subscriptions
	subscriptions := dcb.Group("subscriptions")
	subscriptions.Get("/", dcbHandler.GetAllSubscriptionPaginate)
	subscriptions.Put("/unsub/", dcbHandler.Unsubscription)

	// transactions
	transactions := dcb.Group("transactions")
	transactions.Get("/", dcbHandler.GetAllTransactionPaginate)

	// histories
	histories := dcb.Group("histories")
	histories.Get("/", dcbHandler.GetAllHistoryPaginate)

	// message originated
	mo := dcb.Group("mo")
	mo.Get("/", dcbHandler.GetAllMOPaginate)

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
