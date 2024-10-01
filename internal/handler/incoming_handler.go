package handler

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/idprm/go-football-alert/internal/utils"
	"github.com/sirupsen/logrus"
	"github.com/wiliehidayat87/rmqp"
)

var (
	APP_URL           string = utils.GetEnv("APP_URL")
	APP_PORT          string = utils.GetEnv("APP_PORT")
	APP_TZ            string = utils.GetEnv("APP_TZ")
	API_VERSION       string = utils.GetEnv("API_VERSION")
	URI_MYSQL         string = utils.GetEnv("URI_MYSQL")
	URI_REDIS         string = utils.GetEnv("URI_REDIS")
	RMQ_HOST          string = utils.GetEnv("RMQ_HOST")
	RMQ_USER          string = utils.GetEnv("RMQ_USER")
	RMQ_PASS          string = utils.GetEnv("RMQ_PASS")
	RMQ_PORT          string = utils.GetEnv("RMQ_PORT")
	RMQ_VHOST         string = utils.GetEnv("RMQ_VHOST")
	RMQ_URL           string = utils.GetEnv("RMQ_URL")
	AUTH_SECRET       string = utils.GetEnv("AUTH_SECRET")
	PATH_STATIC       string = utils.GetEnv("PATH_STATIC")
	PATH_IMAGE        string = utils.GetEnv("PATH_IMAGE")
	PATH_VIEWS        string = utils.GetEnv("PATH_VIEWS")
	API_FOOTBALL_URL  string = utils.GetEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = utils.GetEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = utils.GetEnv("API_FOOTBALL_HOST")
)

var (
	RMQ_DATA_TYPE       string = "application/json"
	RMQ_USSD_EXCHANGE   string = "E_USSD"
	RMQ_USSD_QUEUE      string = "Q_USSD"
	RMQ_MO_EXCHANGE     string = "E_MO"
	RMQ_MO_QUEUE        string = "Q_MO"
	RMQ_MT_EXCHANGE     string = "E_MT"
	RMQ_MT_QUEUE        string = "Q_MT"
	RMQ_NEWS_EXCHANGE   string = "E_NEWS"
	RMQ_NEWS_QUEUE      string = "Q_NEWS"
	MT_FIRSTPUSH        string = "FIRSTPUSH"
	MT_RENEWAL          string = "RENEWAL"
	MT_PREDICTION       string = "PREDICTION"
	MT_CREDIT_GOAL      string = "CREDIT_GOAL"
	MT_NEWS             string = "NEWS"
	MT_UNSUB            string = "UNSUB"
	STATUS_SUCCESS      string = "SUCCESS"
	STATUS_FAILED       string = "FAILED"
	SUBJECT_FIRSTPUSH   string = "FIRSTPUSH"
	SUBJECT_DAILYPUSH   string = "DAILYPUSH"
	SUBJECT_FP_SMS      string = "FP_SMS"
	SUBJECT_DP_SMS      string = "DP_SMS"
	SUBJECT_RENEWAL     string = "RENEWAL"
	SUBJECT_PREDICTION  string = "PREDICTION"
	SUBJECT_CREDIT_GOAL string = "CREDIT_GOAL"
	SUBJECT_NEWS        string = "NEWS"
	SUBJECT_UNSUB       string = "UNSUB"
	ACT_PREDICTION      string = "PREDICTION"
	ACT_NEWS            string = "NEWS"
	CHANNEL_USSD        string = "USSD"
	CHANNEL_SMS         string = "SMS"
)

type IncomingHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	menuService         services.IMenuService
	ussdService         services.IUssdService
	leagueService       services.ILeagueService
	teamService         services.ITeamService
	fixtureService      services.IFixtureService
	livescoreService    services.ILiveScoreService
	predictionService   services.IPredictionService
	newsService         services.INewsService
	scheduleService     services.IScheduleService
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	rewardService       services.IRewardService
}

func NewIncomingHandler(
	rmq rmqp.AMQP,
	logger *logger.Logger,
	menuService services.IMenuService,
	ussdService services.IUssdService,
	leagueService services.ILeagueService,
	teamService services.ITeamService,
	fixtureService services.IFixtureService,
	livescoreService services.ILiveScoreService,
	predictionService services.IPredictionService,
	newsService services.INewsService,
	scheduleService services.IScheduleService,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	rewardService services.IRewardService,
) *IncomingHandler {
	return &IncomingHandler{
		rmq:                 rmq,
		logger:              logger,
		menuService:         menuService,
		ussdService:         ussdService,
		leagueService:       leagueService,
		teamService:         teamService,
		fixtureService:      fixtureService,
		livescoreService:    livescoreService,
		predictionService:   predictionService,
		newsService:         newsService,
		scheduleService:     scheduleService,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		rewardService:       rewardService,
	}

}

var validate = validator.New()

func ValidateStruct(data interface{}) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (h *IncomingHandler) Sub(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *IncomingHandler) UnSub(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *IncomingHandler) LandingPage(c *fiber.Ctx) error {
	if h.serviceService.IsService(c.Params("service")) {
		return c.Render("fb-alert/sub", fiber.Map{})
	}
	return c.Redirect("https://www.google.com/")
}

func (h *IncomingHandler) MessageOriginated(c *fiber.Ctx) error {
	l := h.logger.Init("mo", true)

	req := new(model.MORequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadGateway,
				Message:    err.Error(),
			},
		)
	}

	l.WithFields(logrus.Fields{"request": req}).Info("MO")

	h.rmq.IntegratePublish(
		RMQ_MO_EXCHANGE,
		RMQ_MO_QUEUE,
		RMQ_DATA_TYPE, "", string(jsonData),
	)

	return c.Status(fiber.StatusOK).JSON(
		&model.WebResponse{
			Error:      false,
			StatusCode: fiber.StatusOK,
			Message:    "Successful",
		},
	)
}

func (h *IncomingHandler) Main(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	menu, err := h.menuService.GetBySlug("home")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"&", "&amp;",
	)
	replace := replacer.Replace(menu.GetTemplateXML())
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Menu(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	req.SetMsisdn(c.Get("User-MSISDN"))

	// if menu unavailable
	if !h.menuService.IsSlug(req.GetSlug()) {
		menu, err := h.menuService.GetBySlug("404")
		if err != nil {
			return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(menu.GetTemplateXML())
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	menu, err := h.menuService.GetBySlug(req.GetSlug())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	// if is_confirm = true
	if menu.IsConfirm {
		// // if sub not active
		// if !h.subscriptionService.IsActiveSubscriptionByCategory(menu.GetCategory(), req.GetMsisdn()) {
		// 	services, err := h.serviceService.GetAllByCategory(menu.GetCategory())
		// 	if err != nil {
		// 		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		// 	}

		// 	// package
		// 	var servicesData []string
		// 	for _, s := range services {
		// 		row := `<a href="` + APP_URL + `/` + API_VERSION + `/ussd/select?slug=` + req.GetSlug() + `&category=` + menu.GetCategory() + `&package=` + s.GetPackage() + `">` + s.GetName() + " (" + s.GetPriceToString() + ")" + "</a>"
		// 		servicesData = append(servicesData, row)
		// 	}
		// 	servicesString := strings.Join(servicesData, "\n")

		// 	replacer := strings.NewReplacer(
		// 		"{{.url}}", APP_URL,
		// 		"{{.version}}", API_VERSION,
		// 		"{{.title}}", req.GetTitle(),
		// 		"{{.data}}", servicesString,
		// 		"&", "&amp;",
		// 	)

		// 	data, err := os.ReadFile(PATH_VIEWS + "/xml/package.xml")
		// 	if err != nil {
		// 		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		// 	}
		// 	replace := replacer.Replace(string(data))
		// 	return c.Status(fiber.StatusOK).SendString(replace)
		// }

		var data string
		page := strconv.Itoa(req.GetPage() + 1)
		paginate := `<a href="` + APP_URL + `/` + API_VERSION + `/ussd/q?slug=` + req.GetSlug() + `&amp;page=` + page + `">Suiv</a>`

		if req.GetSlug() == "lm-live-match" {
			data = h.LiveMatchs()
		}

		if req.GetSlug() == "lm-schedule" {
			data = h.Schedules()
		}

		if req.GetSlug() == "lm-schedule" {
			data = h.Schedules()
		}

		if req.GetSlug() == "flash-news" {
			data = h.FlashNews(req.GetPage() + 1)
		}

		if req.GetSlug() == "credit-goal" {
			data = ""
		}

		if req.GetSlug() == "champ" {
			data = ""
		}

		if req.GetSlug() == "champ-results" {
			data = ""
		}

		if req.GetSlug() == "champ-standings" {
			data = ""
		}

		if req.GetSlug() == "champ-schedule" {
			data = ""
		}

		if req.GetSlug() == "champ-team" {
			data = ""
		}

		if req.GetSlug() == "champ-credit-score" {
			data = ""
		}

		if req.GetSlug() == "champ-creditgoal" {
			data = ""
		}

		if req.GetSlug() == "champ-sms-alerte" {
			data = ""
		}

		if req.GetSlug() == "champ-sms-alerte-equipe" {
			data = ""
		}

		if req.GetSlug() == "prediction" {
			data = ""
		}

		if req.GetSlug() == "sms-alerte" {
			data = ""
		}

		if req.GetSlug() == "sms-kit-foot" {
			data = ""
		}

		if req.GetSlug() == "sms-foot-europe" {
			data = ""
		}

		if req.GetSlug() == "sms-foot-afrique" {
			data = ""
		}

		if req.GetSlug() == "sms-alerte-equipe" {
			data = ""
		}

		if req.GetSlug() == "sms-foot-international" {
			data = ""
		}

		if req.GetSlug() == "kit-foot" {
			data = ""
		}

		if req.GetSlug() == "kit-foot-champ" {
			data = ""
		}

		if req.GetSlug() == "kit-foot-premier-league" {
			data = ""
		}

		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"{{.data}}", data,
			"{{.paginate}}", paginate,
			"&", "&amp;",
		)
		replace := replacer.Replace(string(menu.GetTemplateXML()))
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	var data string
	if req.GetSlug() == "foot-europe-premier-league" {
		data = h.PremierLeagues()
	}

	if req.GetSlug() == "foot-europe-la-liga" {
		data = h.LaLigas()
	}

	if req.GetSlug() == "foot-europe-ligue-1" {
		data = h.Ligue1s()
	}

	if req.GetSlug() == "foot-europe-serie-a" {
		data = h.SerieA()
	}

	if req.GetSlug() == "foot-europe-bundesligua" {
		data = h.Bundesliguas()
	}

	if req.GetSlug() == "foot-europe-champ-portugal" {
		data = h.PrimeiraLigas()
	}

	page := strconv.Itoa(req.GetPage() + 1)
	paginate := `<a href="` + APP_URL + `/` + API_VERSION + `/ussd/q?slug=` + req.GetSlug() + `&amp;page=` + page + `">Suiv</a>`
	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"{{.data}}", data,
		"{{.paginate}}", paginate,
		"&", "&amp;",
	)
	replace := replacer.Replace(string(menu.GetTemplateXML()))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Detail(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	req.SetMsisdn(c.Get("User-MSISDN"))

	menu, err := h.menuService.GetBySlug(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"{{.slug}}", req.GetSlug(),
		"{{.title}}", req.GetTitle(),
		"&", "&amp;",
	)
	replace := replacer.Replace(string(menu.GetTemplateXML()))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Select(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	req.SetMsisdn(c.Get("User-MSISDN"))

	// if menu unavailable
	if !h.menuService.IsSlug(req.GetSlug()) {

		data, err := os.ReadFile(PATH_VIEWS + "/xml/404.xml")
		if err != nil {
			return c.Status(fiber.StatusBadGateway).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(string(data))
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	service, err := h.serviceService.GetByPackage(req.GetCategory(), req.GetPackage())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	menu, err := os.ReadFile(PATH_VIEWS + "/xml/confirm.xml")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"{{.slug}}", req.GetSlug(),
		"{{.code}}", service.GetCode(),
		"{{.service}}", service.GetName(),
		"{{.price}}", service.GetPriceToString(),
		"&", "&amp;",
	)
	replace := replacer.Replace(string(menu))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Pagination(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	req.SetMsisdn(c.Get("User-MSISDN"))

	// if menu unavailable
	if !h.menuService.IsSlug(req.GetSlug()) {
		menu, err := h.menuService.GetBySlug("404")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(string(menu.GetTemplateXML()))
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	// if service unavailable
	if !h.serviceService.IsService(req.GetCode()) {
		menu, err := h.menuService.GetBySlug("404")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(menu.GetTemplateXML())
		return c.Status(fiber.StatusOK).SendString(replace)
	}
	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"&", "&amp;",
	)
	replace := replacer.Replace(string(""))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Buy(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml; charset=utf-8")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			&model.WebResponse{
				Error:      true,
				StatusCode: fiber.StatusBadRequest,
				Message:    err.Error(),
			},
		)
	}

	req.SetMsisdn(c.Get("User-MSISDN"))

	// if menu unavailable
	if !h.menuService.IsSlug(req.GetSlug()) {
		menu, err := h.menuService.GetBySlug("404")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(menu.GetTemplateXML())
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	// if service unavailable
	if !h.serviceService.IsService(req.GetCode()) {
		menu, err := h.menuService.GetBySlug("404")
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString(err.Error())
		}
		replacer := strings.NewReplacer(
			"{{.url}}", APP_URL,
			"{{.version}}", API_VERSION,
			"&", "&amp;",
		)
		replace := replacer.Replace(menu.GetTemplateXML())
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	service, err := h.serviceService.Get(req.GetCode())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	menuSuccess, err := h.menuService.GetBySlug("success")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	menuFailed, err := h.menuService.GetBySlug("failed")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	replacer := strings.NewReplacer(
		"{{.url}}", APP_URL,
		"{{.version}}", API_VERSION,
		"{{.slug}}", req.GetSlug(),
		"{{.category}}", req.GetCategory(),
		"{{.package}}", req.GetPackage(),
		"{{.service}}", service.GetName(),
		"{{.price}}", service.GetPriceToString(),
		"&", "&amp;",
	)

	if req.IsYes() {
		h.subscriptionService.Save(
			&entity.Subscription{
				ServiceID:   service.GetId(),
				Category:    service.GetCategory(),
				Msisdn:      req.GetMsisdn(),
				Channel:     CHANNEL_USSD,
				LatestTrxId: "",
				IsActive:    true,
			},
		)
		replace := replacer.Replace(menuSuccess.GetTemplateXML())
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	replace := replacer.Replace(menuFailed.GetTemplateXML())
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) LiveMatchs() string {
	livematchs, err := h.fixtureService.GetAll()
	if err != nil {
		log.Println(err.Error())
	}

	var liveMatchsData []string
	for _, s := range livematchs {
		row := s.Home.GetName() + " - " + s.Away.GetName()
		liveMatchsData = append(liveMatchsData, row)
	}
	liveMatchsString := strings.Join(liveMatchsData, "\n")
	return liveMatchsString
}

func (h *IncomingHandler) Schedules() string {
	schedules, err := h.fixtureService.GetAllScheduleUSSD()
	if err != nil {
		log.Println(err.Error())
	}

	var schedulesData []string
	for _, s := range schedules {
		row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
		schedulesData = append(schedulesData, row)
	}
	schedulesString := strings.Join(schedulesData, "\n")
	return schedulesString
}

func (h *IncomingHandler) FlashNews(page int) string {
	news, err := h.newsService.GetAllUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var newsData []string
	for _, s := range news {
		row := `<a href="` + APP_URL + `/` + API_VERSION + `/ussd/q/detail?slug=flash-news&amp;title=` + s.GetTitleQueryEscape() + `">` + s.GetTitleLimited(20) + "</a>"
		newsData = append(newsData, row)
	}
	newsString := strings.Join(newsData, "\n")
	return newsString
}

func (h *IncomingHandler) ChampionLeagues() string {

	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(1)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}

	return fixturesString
}

func (h *IncomingHandler) PremierLeagues() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(124)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}
	return fixturesString
}

func (h *IncomingHandler) LaLigas() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(297)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}
	return fixturesString
}

func (h *IncomingHandler) Ligue1s() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(236)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}
	return fixturesString
}

func (h *IncomingHandler) SerieA() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(295)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}
	return fixturesString
}

func (h *IncomingHandler) Bundesliguas() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(384)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}
	return fixturesString
}

func (h *IncomingHandler) PrimeiraLigas() string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(441)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := s.Home.GetName() + " - " + s.Away.GetName() + " (" + s.GetFixtureDateToString() + ")"
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}

	return fixturesString
}
