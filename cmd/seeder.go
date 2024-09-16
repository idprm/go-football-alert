package cmd

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

func seederDB(db *gorm.DB) {
	var country []entity.Country
	var service []entity.Service
	var content []entity.Content
	var schedule []entity.Schedule
	var menu []entity.Menu

	var countries = []entity.Country{
		{
			ID:       1,
			Name:     "MALI",
			Code:     "223",
			TimeZone: "GMT",
			Currency: "CFA",
		},
		{
			ID:       2,
			Name:     "GUINEE",
			Code:     "224",
			TimeZone: "GMT",
			Currency: "GNF",
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
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			UssdCode:   "#101#36#",
		},
	}

	var contents = []entity.Content{
		{
			ServiceID: 1,
			Name:      ACT_FIRSTPUSH,
			Value:     "Bravo! Vous venez de souscrire au pack SMS alert qui vs permet de suivre votre equipe favorite Joueur! Jusqu au {date} vous recevrez toutes les infos. {price}{currency}/souscription",
		},
		{
			ServiceID: 1,
			Name:      ACT_RENEWAL,
			Value:     "Votre souscription au service SMS-Alerte City est activee jusqe au {date}. Prix: {price}{currency}. Orange vous remercie",
		},
		{
			ServiceID: 1,
			Name:      ACT_PREDICTION,
			Value:     "Credit Goal: {home}-{away}! Gagnez {credit}F credit a chaque but de votre equipe si elle gagne le match! Envoyes {home-code} ou {away-code} par SMS au {sdc}. {price}{currency}/sms",
		},
		{
			ServiceID: 1,
			Name:      ACT_SUB,
			Value:     "Votre participation a ete enregistree. Si {team} gagne et marque des buts lors du prochain match, vous recevrez {price}{currency} deb bonus par but. {price}{currency}/souscription",
		},
		{
			ServiceID: 1,
			Name:      ACT_CREDIT_GOAL,
			Value:     "Credit Goal: Felicitations! Le score final du match {home}-{away} est {score}. Votre compte va etre credite dans un delai de 72H de {price}{currency}",
		},
		{
			ServiceID: 1,
			Name:      ACT_USER_LOSES,
			Value:     "Test User Loses",
		},
		{
			ServiceID: 1,
			Name:      ACT_UNSUB,
			Value:     "Test Unsub",
		},
	}

	var schedules = []entity.Schedule{
		{
			CountryID:  1,
			Name:       ACT_RENEWAL,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
		{
			CountryID:  1,
			Name:       ACT_PREDICTION,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
		{
			CountryID:  1,
			Name:       ACT_CREDIT_GOAL,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
		{
			CountryID:  1,
			Name:       ACT_NEWS,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
	}

	/**
	1.  #101#36#1 : Live Match (Confirm Message)
	2.  #101#36#2 : Flash News (Confirm Message)
	3.  #101#36#3 : Crédit Goal (Confirm Message)
	4.  #101#36#4 : Champ. Mali (Confirm Message)
	5.  #101#36#5 : Prédiction (Confirm Message)
	6.  #101#36#6 : SMS Alerte (Confirm Message)
	7.  #101#36#7 : Kit Foot (Confirm Message)
	8.  #101#36#8 : Foot Europe (Free Access)
	9.  #101#36#9 : Suiv

	### Level 2

	1.  #101#36#1#1 : Live Match
	2.  #101#36#1#2 : Schedule
	3.  #101#36#1#3 : Line Up
	4.  #101#36#1#4 : Match Stats
	5.  #101#36#1#5 : Display Live match
	7.  #101#36#4#1 : Results
	8.  #101#36#4#2 : Standings
	9.  #101#36#4#3 : Schedule
	10. #101#36#4#4 : Team
	11. #101#36#4#5 : Crédit Score
	12. #101#36#4#6 : Crédit Goal
	13. #101#36#4#7 : SMS Alerte
	14. #101#36#4#8 : SMS Alerte Equipe
	15. #101#36#6#1 : Kit Foot
	16. #101#36#6#2 : Europe
	17. #101#36#6#3 : Afrique
	18. #101#36#6#4 : SMS Alerte Equipe
	19. #101#36#6#5 : Foot International
	20. #101#36#7#1 : Alerte Champ. Mali + Equipe
	21. #101#36#7#2 : Alerte Premier League + Equipe
	22. #101#36#7#3 : Alerte La Liga + Equipe
	23. #101#36#7#4 : Alerte Ligue 1 + Equipe
	24. #101#36#7#5 : Alerte Serie A + Equipe
	25. #101#36#7#6 : Alerte Bundesligue + Equipe
	26. #101#36#8#1 : Champion League
	27. #101#36#8#2 : Premier League
	28. #101#36#8#3 : La Liga
	29. #101#36#8#4 : Ligue 1
	30. #101#36#8#5 : L. Europa
	31. #101#36#8#6 : Serie A
	32. #101#36#8#7 : Bundesligua
	33. #101#36#8#8 : Champ Portugal
	34. #101#36#8#9 : Saudi League

		**/
	var menus = []entity.Menu{
		{
			ID:       1,
			Name:     "Live Match",
			KeyPress: "#101#36#1",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       2,
			Name:     "Live Match",
			KeyPress: "#101#36#1#1",
			ParentID: 1,
			Child:    1,
			Action:   "live-match",
			IsActive: true,
		},
		{
			ID:       3,
			Name:     "Schedule",
			KeyPress: "#101#36#1#2",
			ParentID: 1,
			Child:    2,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       4,
			Name:     "Line Up",
			KeyPress: "#101#36#1#3",
			ParentID: 1,
			Child:    3,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       5,
			Name:     "Match Stats",
			KeyPress: "#101#36#1#4",
			ParentID: 1,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       6,
			Name:     "Display Live match",
			KeyPress: " #101#36#1#5",
			ParentID: 1,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       7,
			Name:     "Flash News",
			KeyPress: "2",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       8,
			Name:     "Crédit Goal",
			KeyPress: "3",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       9,
			Name:     "Champ. Mali",
			KeyPress: "4",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       10,
			Name:     "Results",
			KeyPress: "4*1",
			ParentID: 9,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       11,
			Name:     "Standings",
			KeyPress: "4*2",
			ParentID: 9,
			Child:    2,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       12,
			Name:     "Schedule",
			KeyPress: "4*3",
			ParentID: 9,
			Child:    3,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       13,
			Name:     "Team",
			KeyPress: "4*4",
			ParentID: 9,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       14,
			Name:     "Crédit Score",
			KeyPress: "4*5",
			ParentID: 9,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       15,
			Name:     "Crédit Goal",
			KeyPress: "4*6",
			ParentID: 9,
			Child:    6,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       16,
			Name:     "SMS Alerte",
			KeyPress: "4*7",
			ParentID: 9,
			Child:    7,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       17,
			Name:     "SMS Alerte Equipe",
			KeyPress: "4*8",
			ParentID: 9,
			Child:    8,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       18,
			Name:     "Prédiction",
			KeyPress: "5",
			ParentID: 0,
			Child:    0,
			Action:   "prediction",
			IsActive: true,
		},
		{
			ID:       19,
			Name:     "SMS Alerte",
			KeyPress: "6",
			ParentID: 0,
			Child:    0,
			Action:   "prediction",
			IsActive: true,
		},
		{
			ID:       20,
			Name:     "Kit Foot",
			KeyPress: "6*1",
			ParentID: 19,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       21,
			Name:     "Europe",
			KeyPress: "6*2",
			ParentID: 19,
			Child:    2,
			Action:   "foot-europe",
			IsActive: true,
		},
		{
			ID:       22,
			Name:     "Afrique",
			KeyPress: "6*3",
			ParentID: 19,
			Child:    3,
			Action:   "foot-afrique",
			IsActive: true,
		},
		{
			ID:       23,
			Name:     "SMS Alerte Equipe",
			KeyPress: "6*4",
			ParentID: 19,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       24,
			Name:     "Foot International",
			KeyPress: "6*5",
			ParentID: 19,
			Child:    5,
			Action:   "foot-international",
			IsActive: true,
		},
		{
			ID:       25,
			Name:     "Kit Foot",
			KeyPress: "7",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       26,
			Name:     "Alerte Champ. Mali + Equipe",
			KeyPress: "7*1",
			ParentID: 25,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       27,
			Name:     "Alerte Premier League + Equipe",
			KeyPress: "7*2",
			ParentID: 25,
			Child:    2,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       28,
			Name:     "Alerte La Liga + Equipe",
			KeyPress: "7*3",
			ParentID: 25,
			Child:    3,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       29,
			Name:     "Alerte Ligue 1 + Equipe",
			KeyPress: "7*4",
			ParentID: 25,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       30,
			Name:     "Alerte Serie A + Equipe",
			KeyPress: "7*5",
			ParentID: 25,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       31,
			Name:     "Alerte Bundesligue + Equipe",
			KeyPress: "7*6",
			ParentID: 25,
			Child:    6,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       32,
			Name:     "Foot Europe",
			KeyPress: "8",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       33,
			Name:     "Champion League",
			KeyPress: "8*1",
			ParentID: 32,
			Child:    1,
			Action:   "champion-league",
			IsActive: true,
		},
		{
			ID:       34,
			Name:     "Premier League",
			KeyPress: "8*2",
			ParentID: 32,
			Child:    2,
			Action:   "premier-league",
			IsActive: true,
		},
		{
			ID:       35,
			Name:     "La Liga",
			KeyPress: "8*3",
			ParentID: 32,
			Child:    3,
			Action:   "la-liga",
			IsActive: true,
		},
		{
			ID:       36,
			Name:     "Ligue 1",
			KeyPress: "8*4",
			ParentID: 32,
			Child:    4,
			Action:   "ligue-1",
			IsActive: true,
		},
		{
			ID:       37,
			Name:     "L. Europa",
			KeyPress: "8*5",
			ParentID: 32,
			Child:    5,
			Action:   "l-europa",
			IsActive: true,
		},
		{
			ID:       38,
			Name:     "Serie A",
			KeyPress: "8*6",
			ParentID: 32,
			Child:    6,
			Action:   "serie-a",
			IsActive: true,
		},
		{
			ID:       39,
			Name:     "Bundesligua",
			KeyPress: "8*7",
			ParentID: 32,
			Child:    7,
			Action:   "bundesligua",
			IsActive: true,
		},
		{
			ID:       40,
			Name:     "Champ Portugal",
			KeyPress: "8*8",
			ParentID: 32,
			Child:    8,
			Action:   "champ-portugal",
			IsActive: true,
		},
		{
			ID:       41,
			Name:     "Saudi League",
			KeyPress: "8*9",
			ParentID: 32,
			Child:    9,
			Action:   "saudi-league",
			IsActive: true,
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

	if db.Find(&schedule).RowsAffected == 0 {
		for i, _ := range schedules {
			db.Model(&entity.Schedule{}).Create(&schedules[i])
		}
		log.Println("schedules migrated")
	}

	if db.Find(&menu).RowsAffected == 0 {
		for i, _ := range menus {
			db.Model(&entity.Menu{}).Create(&menus[i])
		}
		log.Println("menus migrated")
	}
}
