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
			Value:     "Test Firstpush",
		},
		{
			ServiceID: 1,
			Name:      ACT_RENEWAL,
			Value:     "Test Renewal",
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
}
