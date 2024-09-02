package cmd

import (
	"log"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"gorm.io/gorm"
)

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
