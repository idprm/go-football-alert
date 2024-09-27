package handler

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/model"
)

const (
	ACT_MENU    string = "MENU"
	ACT_NO_MENU string = "NO_MENU"
	ACT_CONFIRM string = "CONFIRM"
	ACT_REG     string = "REG"
)

var (
	USSD_TITLE    string = "Orange Football Club, votre choix:"
	USSD_MENU_404 string = "Menu not found"
)

func (h *IncomingHandler) IsService(code string) bool {
	return h.serviceService.IsService(code)
}

func (h *IncomingHandler) getService(code string) (*entity.Service, error) {
	return h.serviceService.Get(code)
}

func (h *IncomingHandler) ChooseMenu(req *model.UssdRequest) {
}

func (h *IncomingHandler) Schedule() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Schedule", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Lineup() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Lineup", KeyPress: "0"},
	}
}

func (h *IncomingHandler) MatchStats() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Match Stats", KeyPress: "0"},
	}
}

func (h *IncomingHandler) DisplayLiveMatch() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Display Live Match", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampResults() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Results", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampStandings() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Standings", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSchedule() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Schedule", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampTeam() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Team", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampCreditScore() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Credit Score", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampCreditGoal() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Credit Goal", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSMSAlerte() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ SMS Alerte", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampSMSAlerteEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ SMS Alerte Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Europe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Europe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Afrique() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Afrique", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SMSAlerteEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No SMS Alerte Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) FootInternational() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Foot International", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteChampMaliEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Champ Mali Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlertePremierLeagueEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Premier League Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteLaLigaEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte LaLiga Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteLigue1Equipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Ligue 1 Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteSerieAEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Serie A Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) AlerteBundesligueEquipe() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Alerte Bundesligue Equipe", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampionLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champion League", KeyPress: "0"},
	}
}

func (h *IncomingHandler) PremierLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Premier League", KeyPress: "0"},
	}
}

func (h *IncomingHandler) LaLiga() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No La Liga", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Ligue1() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Ligue 1", KeyPress: "0"},
	}
}

func (h *IncomingHandler) LEuropa() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No L Europa", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SerieA() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Serie A", KeyPress: "0"},
	}
}

func (h *IncomingHandler) Bundesligua() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Bundesligua", KeyPress: "0"},
	}
}

func (h *IncomingHandler) ChampPortugal() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Champ Portugal", KeyPress: "0"},
	}
}

func (h *IncomingHandler) SaudiLeague() []*entity.Menu {
	return []*entity.Menu{
		{ID: 0, Name: "No Saudi League", KeyPress: "0"},
	}
}
