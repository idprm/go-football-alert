package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	URL_MT            string = utils.GetEnv("URL_MT")
	USER_MT           string = utils.GetEnv("USER_MT")
	PASS_MT           string = utils.GetEnv("PASS_MT")
)

var (
	RMQ_DATA_TYPE       string = "application/json"
	RMQ_USSD_EXCHANGE   string = "E_USSD"
	RMQ_USSD_QUEUE      string = "Q_USSD"
	RMQ_SMS_EXCHANGE    string = "E_SMS"
	RMQ_SMS_QUEUE       string = "Q_SMS"
	RMQ_MT_EXCHANGE     string = "E_MT"
	RMQ_MT_QUEUE        string = "Q_MT"
	RMQ_NEWS_EXCHANGE   string = "E_NEWS"
	RMQ_NEWS_QUEUE      string = "Q_NEWS"
	MT_SMS_ALERTE       string = "SMS_ALERTE"
	MT_CREDIT_GOAL      string = "CREDIT_GOAL"
	MT_PREDICTION       string = "PREDICTION"
	MT_FIRSTPUSH        string = "FIRSTPUSH"
	MT_RENEWAL          string = "RENEWAL"
	MT_NEWS             string = "NEWS"
	MT_UNSUB            string = "UNSUB"
	STATUS_SUCCESS      string = "SUCCESS"
	STATUS_FAILED       string = "FAILED"
	SUBJECT_FIRSTPUSH   string = "FIRSTPUSH"
	SUBJECT_DAILYPUSH   string = "DAILYPUSH"
	SUBJECT_FREEPUSH    string = "FREE"
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

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	errors := ValidateStruct(*req)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	l.WithFields(logrus.Fields{"request": req}).Info("MO")
	h.rmq.IntegratePublish(
		RMQ_SMS_EXCHANGE,
		RMQ_SMS_QUEUE,
		RMQ_DATA_TYPE, "", string(jsonData),
	)

	return c.Status(fiber.StatusOK).SendString("OK")
}

func (h *IncomingHandler) Main(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml")
	menu, err := h.menuService.GetBySlug("home")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	replacer := strings.NewReplacer(
		"{{.url}}", c.BaseURL(),
		"{{.version}}", API_VERSION,
		"&", "&amp;",
	)
	replace := replacer.Replace(menu.GetTemplateXML())
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Menu(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	// set msisdn
	req.SetMsisdn(c.Get("User-MSISDN"))

	// if menu or page not found
	if !h.menuService.IsSlug(req.GetSlug()) {
		return c.Status(fiber.StatusOK).SendString(h.PageNotFound(c.BaseURL()))
	}

	menu, err := h.menuService.GetBySlug(req.GetSlug())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	// if is_confirm = true
	if menu.IsConfirm {

		var data string

		if req.IsLmLiveMatch() {
			data = h.LiveMatchs(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsLmLiveMatch() {
			data = h.Schedules(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsFlashNews() {
			data = h.FlashNews(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsCreditGoal() {
			data = h.CreditGoal(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampResults() {
			data = h.ChampResults(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampStandings() {
			data = h.ChampStandings(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampSchedules() {
			data = h.ChampSchedules(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampTeam() {
			data = h.ChampTeam(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampCreditScore() {
			data = h.ChampCreditScore(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsChampCreditGoal() {
			data = h.ChampCreditGoal(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "champ-sms-alerte" {
			data = h.ChampSMSAlerte(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "champ-sms-alerte-equipe" {
			data = h.ChampSMSAlerteEquipe(c.BaseURL(), req.GetPage()+1)
		}

		if req.IsPrediction() {
			data = h.Prediction(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "sms-alerte" {
			data = h.SMSAlerte(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "sms-alerte-equipe" {
			data = h.SMSAlerteEquipe(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "kit-foot" {
			data = h.KitFoot(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "kit-foot-by-league" {
			data = h.KitFootByLeague(c.BaseURL(), req.GetLeagueId(), req.GetPage()+1)
		}

		if req.GetSlug() == "foot-europe" {
			data = h.FootEurope(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "foot-afrique" {
			data = h.FootAfrique(c.BaseURL(), req.GetPage()+1)
		}

		if req.GetSlug() == "foot-international" {
			data = h.FootInternational(c.BaseURL(), req.GetPage()+1)
		}

		leagueId := strconv.Itoa(req.GetLeagueId())
		teamId := strconv.Itoa(req.GetTeamId())
		page := strconv.Itoa(req.GetPage() + 1)

		paginate := `<a href="` + c.BaseURL() +
			`/` + API_VERSION +
			`/ussd/q?slug=` + req.GetSlug() +
			`&amp;title=` + req.GetTitleQueryEscape() +
			`&amp;league_id=` + leagueId +
			`&amp;team_id=` + teamId +
			`&amp;page=` + page +
			`">Suiv</a>`

		replacer := strings.NewReplacer(
			"{{.url}}", c.BaseURL(),
			"{{.version}}", API_VERSION,
			"{{.date}}", utils.FormatFR(time.Now()),
			"{{.data}}", data,
			"{{.title}}", req.GetTitleWithoutAccents(),
			"{{.paginate}}", paginate,
			"&", "&amp;",
		)
		replace := replacer.Replace(string(menu.GetTemplateXML()))
		return c.Status(fiber.StatusOK).SendString(replace)
	}

	var data string

	leagueId := strconv.Itoa(req.GetLeagueId())
	teamId := strconv.Itoa(req.GetTeamId())
	page := strconv.Itoa(req.GetPage() + 1)

	paginate := `<a href="` + c.BaseURL() +
		`/` + API_VERSION +
		`/ussd/q?slug=` + req.GetSlug() +
		`&amp;title=` + req.GetTitleQueryEscape() +
		`&amp;league_id=` + leagueId +
		`&amp;team_id=` + teamId +
		`&amp;page=` + page +
		`">Suiv</a>`
	replacer := strings.NewReplacer(
		"{{.url}}", c.BaseURL(),
		"{{.version}}", API_VERSION,
		"{{.date}}", utils.FormatFR(time.Now()),
		"{{.data}}", data,
		"{{.title}}", req.GetTitleWithoutAccents(),
		"{{.paginate}}", paginate,
		"&", "&amp;",
	)
	replace := replacer.Replace(string(menu.GetTemplateXML()))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Detail(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}
	// setter msisdn
	req.SetMsisdn(c.Get("User-MSISDN"))

	log.Println(req.GetMsisdn())

	// check if msisdn not found
	if !req.IsMsisdn() {
		return c.Status(fiber.StatusOK).SendString(h.MsisdnNotFound(c.BaseURL()))
	}

	// get menu
	menu, err := h.menuService.GetBySlug(req.GetSlug())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	// check sub is active
	if !h.subscriptionService.IsActiveSubscriptionByCategory(menu.GetCategory(), req.GetMsisdn()) {
		return c.Status(fiber.StatusOK).SendString(h.IsActiveSubscription(c.BaseURL(), req.GetSlug(), menu.GetCategory()))
	}

	replacer := strings.NewReplacer(
		"{{.url}}", c.BaseURL(),
		"{{.version}}", API_VERSION,
		"{{.date}}", utils.FormatFR(time.Now()),
		"{{.slug}}", req.GetSlug(),
		"{{.title}}", req.GetTitleWithoutAccents(),
		"&", "&amp;",
	)
	replace := replacer.Replace(string(menu.GetTemplateXML()))
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Confirm(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml")
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// setter msisdn
	req.SetMsisdn(c.Get("User-MSISDN"))

	// check if msisdn not found
	if !req.IsMsisdn() {
		return c.Status(fiber.StatusOK).SendString(h.MsisdnNotFound(c.BaseURL()))
	}

	// if menu or page not found
	if !h.menuService.IsSlug(req.GetSlug()) {
		return c.Status(fiber.StatusOK).SendString(h.PageNotFound(c.BaseURL()))
	}

	// if service unavailable
	if !h.serviceService.IsService(req.GetCode()) {
		return c.Status(fiber.StatusOK).SendString(h.PageNotFound(c.BaseURL()))
	}

	service, err := h.serviceService.Get(req.GetCode())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	menu, err := h.menuService.GetBySlug("confirm-sms-alerte")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	replacer := strings.NewReplacer(
		"{{.url}}", c.BaseURL(),
		"{{.version}}", API_VERSION,
		"{{.slug}}", req.GetSlug(),
		"{{.title}}", service.GetName(),
		"{{.category}}", req.GetCategory(),
		"{{.package}}", req.GetPackage(),
		"{{.code}}", service.GetCode(),
		"{{.price}}", service.GetPriceToString(),
		"&", "&amp;",
	)

	replace := replacer.Replace(menu.GetTemplateXML())
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) Buy(c *fiber.Ctx) error {
	c.Set("Content-type", "application/xml")
	l := h.logger.Init("mo", true)
	req := new(model.UssdRequest)

	err := c.QueryParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// setter msisdn
	req.SetMsisdn(c.Get("User-MSISDN"))

	// check if msisdn not found
	if !req.IsMsisdn() {
		return c.Status(fiber.StatusOK).SendString(h.MsisdnNotFound(c.BaseURL()))
	}

	// if menu or page not found
	if !h.menuService.IsSlug("success") {
		return c.Status(fiber.StatusOK).SendString(h.PageNotFound(c.BaseURL()))
	}

	// if service unavailable
	if !h.serviceService.IsService(req.GetCode()) {
		return c.Status(fiber.StatusOK).SendString(h.PageNotFound(c.BaseURL()))
	}

	menu, err := h.menuService.GetBySlug("success")
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	service, err := h.serviceService.Get(req.GetCode())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).SendString(err.Error())
	}

	l.WithFields(logrus.Fields{"request": req}).Info("MO")
	h.rmq.IntegratePublish(
		RMQ_USSD_EXCHANGE,
		RMQ_USSD_QUEUE,
		RMQ_DATA_TYPE, "", string(jsonData),
	)

	replacer := strings.NewReplacer(
		"{{.url}}", c.BaseURL(),
		"{{.version}}", API_VERSION,
		"{{.service}}", service.GetName(),
		"&", "&amp;",
	)
	replace := replacer.Replace(menu.GetTemplateXML())
	return c.Status(fiber.StatusOK).SendString(replace)
}

func (h *IncomingHandler) PageNotFound(baseUrl string) string {
	menu, err := h.menuService.GetBySlug("404")
	if err != nil {
		return err.Error()
	}
	replacer := strings.NewReplacer(
		"{{.url}}", baseUrl,
		"{{.version}}", API_VERSION,
		"&", "&amp;",
	)
	return replacer.Replace(menu.GetTemplateXML())
}

func (h *IncomingHandler) MsisdnNotFound(baseUrl string) string {
	// check if msisdn not found
	menu, err := h.menuService.GetBySlug("msisdn_not_found")
	if err != nil {
		log.Println(err.Error())
	}
	replacer := strings.NewReplacer(
		"{{.url}}", baseUrl,
		"{{.version}}", API_VERSION,
		"&", "&amp;",
	)
	return replacer.Replace(menu.GetTemplateXML())
}

func (h *IncomingHandler) IsActiveSubscription(baseUrl, slug, category string) string {
	services, err := h.serviceService.GetAllByCategory(category)
	if err != nil {
		return err.Error()
	}

	menu, err := h.menuService.GetBySlug("package")
	if err != nil {
		return err.Error()
	}
	// package
	var servicesData []string
	for _, s := range services {
		row := `<a href="` +
			baseUrl + `/` +
			API_VERSION +
			`/ussd/confirm?slug=` + slug +
			`&code=` + s.Code +
			`&category=` + category +
			`&package=` + s.GetPackage() + `">` +
			s.GetName() + " (" + s.GetPriceToString() + ")" +
			"</a>"
		servicesData = append(servicesData, row)
	}
	servicesString := strings.Join(servicesData, "\n")

	replacer := strings.NewReplacer(
		"{{.url}}", baseUrl,
		"{{.version}}", API_VERSION,
		"{{.title}}", "S'abonner",
		"{{.data}}", servicesString,
		"&", "&amp;",
	)
	return replacer.Replace(string(menu.GetTemplateXML()))
}

/**
** Menu service
**/
func (h *IncomingHandler) LiveMatchs(baseUrl string, page int) string {
	livematchs, err := h.fixtureService.GetAllLiveMatchUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var liveMatchsData []string
	var liveMatchsString string
	if len(livematchs) > 0 {
		for _, s := range livematchs {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION + `/ussd/q/detail?slug=lm-live-match&amp;title=` +
				s.GetFixtureNameQueryEscape() + `">` +
				s.GetFixtureName() +
				`</a>`
			liveMatchsData = append(liveMatchsData, row)
		}
		liveMatchsString = strings.Join(liveMatchsData, "\n")
	} else {
		liveMatchsString = "No data"
	}
	return liveMatchsString
}

func (h *IncomingHandler) Schedules(baseUrl string, page int) string {
	schedules, err := h.fixtureService.GetAllScheduleUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var schedulesData []string
	var schedulesString string
	if len(schedules) > 0 {
		for _, s := range schedules {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q/detail?slug=lm-schedule&amp;title=` +
				s.GetFixtureNameQueryEscape() + `">` +
				s.GetFixtureName() + `</a>`
			schedulesData = append(schedulesData, row)
		}
		schedulesString = strings.Join(schedulesData, "\n")
	} else {
		schedulesString = "No data"
	}
	return schedulesString
}

func (h *IncomingHandler) FlashNews(baseUrl string, page int) string {
	news, err := h.newsService.GetAllUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var newsData []string
	var newsString string
	if len(news) > 0 {
		for _, s := range news {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION + `/ussd/q/detail?slug=flash-news&amp;category=SMSALERTE&amp;title=` +
				s.GetTitleQueryEscape() + `">` +
				s.GetTitleLimited(20) +
				`</a>`
			newsData = append(newsData, row)
		}
		newsString = strings.Join(newsData, "\n")
	} else {
		newsString = "No data"
	}
	return newsString
}

func (h *IncomingHandler) SMSAlerte(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampResults(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampStandings(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampSchedules(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampTeam(baseUrl string, page int) string {
	teams, err := h.teamService.GetAllTeamUSSD(1, page)
	if err != nil {
		log.Println(err.Error())
	}
	var teamsData []string
	var teamsString string

	if len(teams) > 0 {
		for _, s := range teams {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q/detail?slug=champ-mali&amp;title=` +
				s.Team.GetNameQueryEscape() + `">` +
				s.Team.GetName() +
				`</a>`
			teamsData = append(teamsData, row)
		}
		teamsString = strings.Join(teamsData, "\n")
	} else {
		teamsString = "No match"
	}

	return teamsString
}

func (h *IncomingHandler) ChampCreditScore(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampCreditGoal(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampSMSAlerte(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampSMSAlerteEquipe(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) CreditGoal(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) Prediction(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) KitFoot(baseUrl string, page int) string {
	leagues, err := h.leagueService.GetAllUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var leaguesData []string
	var leagueString string
	if len(leagues) > 0 {
		for _, s := range leagues {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/buy?slug=confirm-sms-alerte&amp;code=FC30&amp;league_id=` +
				s.GetIdToString() +
				`&amp;title=` + s.GetNameQueryEscape() +
				`">Alerte ` + s.GetNameWithoutAccents() +
				`</a>`
			leaguesData = append(leaguesData, row)
		}
		leagueString = strings.Join(leaguesData, "\n")
	} else {
		leagueString = "No data"
	}
	return leagueString
}

func (h *IncomingHandler) KitFootByLeague(baseUrl string, leagueId, page int) string {
	teams, err := h.teamService.GetAllTeamUSSD(leagueId, page)
	if err != nil {
		log.Println(err.Error())
	}

	var teamsData []string
	var teamsString string
	if len(teams) > 0 {
		for _, s := range teams {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q/detail?slug=kit-foot-by-team&amp;team_id=` +
				s.Team.GetIdToString() + `&amp;title=` +
				s.Team.GetNameQueryEscape() + `">` +
				s.Team.GetName() +
				`</a>`
			teamsData = append(teamsData, row)
		}
		teamsString = strings.Join(teamsData, "\n")
	} else {
		teamsString = "No data"
	}
	return teamsString
}

func (h *IncomingHandler) KitFootByTeam(baseUrl string, teamId, page int) string {
	return ""
}

func (h *IncomingHandler) FootEurope(baseUrl string, page int) string {
	leagues, err := h.leagueService.GetAllEuropeUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var leaguesData []string
	var leagueString string
	if len(leagues) > 0 {
		for _, s := range leagues {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q?slug=kit-foot-by-league&amp;league_id=` + s.GetIdToString() +
				`&amp;title=` + s.GetNameQueryEscape() +
				`">Alerte ` + s.GetNameWithoutAccents() +
				`</a>`
			leaguesData = append(leaguesData, row)
		}
		leagueString = strings.Join(leaguesData, "\n")
	} else {
		leagueString = "No data"
	}
	return leagueString
}

func (h *IncomingHandler) FootAfrique(baseUrl string, page int) string {
	leagues, err := h.leagueService.GetAllAfriqueUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var leaguesData []string
	var leagueString string
	if len(leagues) > 0 {
		for _, s := range leagues {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q?slug=kit-foot-by-league&amp;league_id=` + s.GetIdToString() +
				`&amp;title=` + s.GetNameQueryEscape() +
				`">Alerte ` + s.GetNameWithoutAccents() +
				`</a>`
			leaguesData = append(leaguesData, row)
		}
		leagueString = strings.Join(leaguesData, "\n")
	} else {
		leagueString = "No data"
	}
	return leagueString
}

func (h *IncomingHandler) FootInternational(baseUrl string, page int) string {
	leagues, err := h.leagueService.GetAllWorldUSSD(page)
	if err != nil {
		log.Println(err.Error())
	}

	var leaguesData []string
	var leagueString string
	if len(leagues) > 0 {
		for _, s := range leagues {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q?slug=kit-foot-by-league&amp;league_id=` + s.GetIdToString() +
				`&amp;title=` + s.GetNameQueryEscape() +
				`">Alerte ` + s.GetNameWithoutAccents() +
				`</a>`
			leaguesData = append(leaguesData, row)
		}
		leagueString = strings.Join(leaguesData, "\n")
	} else {
		leagueString = "No data"
	}
	return leagueString
}

func (h *IncomingHandler) SMSAlerteEquipe(baseUrl string, page int) string {
	return ""
}

func (h *IncomingHandler) ChampionLeagues(baseUrl string, leagueId, page int) string {
	fixtures, err := h.fixtureService.GetAllByLeagueIdUSSD(leagueId, page)
	if err != nil {
		log.Println(err.Error())
	}

	var fixturesData []string
	var fixturesString string
	if len(fixtures) > 0 {
		for _, s := range fixtures {
			row := `<a href="` +
				baseUrl + `/` +
				API_VERSION +
				`/ussd/q/detail?slug=foot-europe-champion-league&amp;title=` +
				s.GetFixtureNameQueryEscape() + `">` +
				s.GetFixtureName() +
				`</a>`
			fixturesData = append(fixturesData, row)
		}
		fixturesString = strings.Join(fixturesData, "\n")
	} else {
		fixturesString = "No match"
	}

	return fixturesString
}
