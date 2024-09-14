package handler

import (
	"encoding/json"
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
	API_FOOTBALL_URL  string = utils.GetEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = utils.GetEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = utils.GetEnv("API_FOOTBALL_HOST")
)

var (
	RMQ_DATA_TYPE       string = "application/json"
	RMQ_MO_EXCHANGE     string = "E_MO"
	RMQ_MO_QUEUE        string = "Q_MO"
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
)

type IncomingHandler struct {
	rmq                 rmqp.AMQP
	logger              *logger.Logger
	menuService         services.IMenuService
	ussdService         services.IUssdService
	leagueService       services.ILeagueService
	seasonService       services.ISeasonService
	teamService         services.ITeamService
	fixtureService      services.IFixtureService
	homeService         services.IHomeService
	awayService         services.IAwayService
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
	seasonService services.ISeasonService,
	teamService services.ITeamService,
	fixtureService services.IFixtureService,
	homeService services.IHomeService,
	awayService services.IAwayService,
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
		seasonService:       seasonService,
		teamService:         teamService,
		fixtureService:      fixtureService,
		homeService:         homeService,
		awayService:         awayService,
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

func (h *IncomingHandler) USSD(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *IncomingHandler) Sub(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *IncomingHandler) UnSub(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OK"})
}

func (h *IncomingHandler) Callback(c *fiber.Ctx) error {
	req := new(model.UssdRequest)

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

	if req.IsMain() {
		main, err := h.menuService.GetAll()
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		var mainData []string
		for i, s := range main {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			mainData = append(mainData, row)
		}

		justString := strings.Join(mainData, "\n")
		return c.Status(fiber.StatusOK).SendString(justString)

	} else {
		if !h.menuService.IsKeyPress(req.GetText()) {
			return c.Status(fiber.StatusNotFound).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusNotFound,
					Message:    "menu_not_found",
				},
			)
		}
		menu, err := h.menuService.GetMenuByKeyPress(req.GetText())
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		submenus, err := h.menuService.GetMenuByParentId(menu.GetId())
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(
				&model.WebResponse{
					Error:      true,
					StatusCode: fiber.StatusBadGateway,
					Message:    err.Error(),
				},
			)
		}

		prevs := []entity.Menu{
			{ID: 0, Name: "Préc", KeyPress: "0"},
			{ID: 98, Name: "Accueil", KeyPress: "98"},
		}

		var subData []string
		subData = append(subData, menu.GetName())
		for i, s := range submenus {
			idx := strconv.Itoa(i + 1)
			row := idx + ". " + s.Name
			subData = append(subData, row)
		}
		if req.IsLevel() {
			var row string
			/**
			 **	for loop data
			 **/
			if req.IsLiveMatch() {
				row = h.LiveMatch()
			}

			if req.IsSchedule() {
				row = h.Schedule()
			}

			if req.IsLineup() {
				row = h.Lineup()
			}

			if req.IsMatchStats() {
				row = h.MatchStats()
			}

			if req.IsDisplayLiveMatch() {

			}

			if req.IsFlashNews() {

			}

			if req.IsCreditGoal() {

			}

			if req.IsChampResults() {

			}

			if req.IsChampStandings() {

			}

			if req.IsChampTeam() {

			}

			if req.IsChampCreditScore() {

			}

			if req.IsChampCreditGoal() {

			}

			if req.IsChampSMSAlerte() {

			}

			if req.IsChampSMSAlerteEquipe() {

			}

			if req.IsPrediction() {

			}

			if req.IsKitFoot() {

			}

			if req.IsEurope() {

			}

			if req.IsAfrique() {

			}

			if req.IsSMSAlerteEquipe() {

			}

			if req.IsFootInternational() {

			}

			if req.IsAlerteChampMaliEquipe() {

			}

			if req.IsAlertePremierLeagueEquipe() {

			}

			if req.IsAlertePremierLeagueEquipe() {

			}

			if req.IsAlerteLaLigaEquipe() {

			}

			if req.IsAlerteLigue1Equipe() {

			}

			if req.IsAlerteSerieAEquipe() {

			}

			if req.IsAlerteBundesligueEquipe() {

			}

			if req.IsChampionLeague() {

			}

			if req.IsPremierLeague() {

			}

			if req.IsLaLiga() {

			}

			if req.IsLigue1() {

			}

			if req.IsSerieA() {

			}

			if req.IsBundesligua() {

			}

			if req.IsChampPortugal() {

			}

			if req.IsSaudiLeague() {

			}

			row = "0. Suiv"
			subData = append(subData, row)
		} else {
			for _, p := range prevs {
				row := p.KeyPress + ". " + p.Name
				subData = append(subData, row)
			}
		}
		justString := strings.Join(subData, "\n")
		return c.Status(fiber.StatusOK).SendString(justString)
	}
}

func (h *IncomingHandler) LiveMatch() string {
	return "No Live Match"
}

func (h *IncomingHandler) Schedule() string {
	return "No Schedule"
}

func (h *IncomingHandler) Lineup() string {
	return "No Lineup"
}

func (h *IncomingHandler) MatchStats() string {
	return "No Match Stats"
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
