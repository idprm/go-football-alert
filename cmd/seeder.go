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
			Category:   "creditgoal",
			Name:       "Credit Goal Daily",
			Code:       "CG1",
			Package:    "jour",
			Price:      100,
			RewardGoal: 10,
			RenewalDay: 1,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         2,
			CountryID:  1,
			Category:   "prediction",
			Name:       "Prediction & Win Daily",
			Code:       "PW1",
			Package:    "jour",
			Price:      50,
			RewardGoal: 10,
			RenewalDay: 1,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         3,
			CountryID:  1,
			Category:   "follow-team",
			Name:       "Follow Team Daily",
			Code:       "FT1",
			Package:    "jour",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 1,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         4,
			CountryID:  1,
			Category:   "follow-competition",
			Name:       "Follow Competition Daily",
			Code:       "FC",
			Package:    "jour",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 1,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         5,
			CountryID:  1,
			Category:   "creditgoal",
			Name:       "Credit Goal Weekly",
			Code:       "CG7",
			Package:    "semaine",
			Price:      100,
			RewardGoal: 10,
			RenewalDay: 7,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         6,
			CountryID:  1,
			Category:   "prediction",
			Name:       "Prediction & Win Weekly",
			Code:       "PW7",
			Package:    "semaine",
			Price:      50,
			RewardGoal: 10,
			RenewalDay: 7,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         7,
			CountryID:  1,
			Category:   "follow-team",
			Name:       "Follow Team Weekly",
			Code:       "FT7",
			Package:    "semaine",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 7,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         8,
			CountryID:  1,
			Category:   "follow-competition",
			Name:       "Follow Competition Weekly",
			Code:       "FC",
			Package:    "semaine",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 7,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         9,
			CountryID:  1,
			Category:   "creditgoal",
			Name:       "Credit Goal Monthly",
			Code:       "CG30",
			Package:    "mois",
			Price:      100,
			RewardGoal: 10,
			RenewalDay: 30,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         10,
			CountryID:  1,
			Category:   "prediction",
			Name:       "Prediction & Win Monthly",
			Code:       "PW30",
			Package:    "mois",
			Price:      50,
			RewardGoal: 10,
			RenewalDay: 30,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         11,
			CountryID:  1,
			Category:   "follow-team",
			Name:       "Follow Team Monthly",
			Code:       "FT30",
			Package:    "mois",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 30,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
		{
			ID:         12,
			CountryID:  1,
			Category:   "follow-competition",
			Name:       "Follow Competition Monthly",
			Code:       "FC30",
			Package:    "mois",
			Price:      200,
			RewardGoal: 0,
			RenewalDay: 30,
			TrialDay:   0,
			UrlTelco:   "http://172.17.111.40:8080/services/OrangeService.OrangeServiceHttpSoap11Endpoint/",
			UserTelco:  "ESERV",
			PassTelco:  "WS001",
			UrlMT:      "http://10.106.0.3:4445/cgi-bin/sendsms?smsc=MOBIMIUM&username={username}&password={password}&from={from}&to={to}&text={text}",
			UserMT:     "send",
			PassMT:     "sendsms",
			ScSubMT:    "37405",
			ScUnsubMT:  "37484",
			ShortCode:  "????",
			UssdCode:   "#101#36#",
			IsActive:   true,
		},
	}

	var contents = []entity.Content{
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CREDIT_GOAL_SUB,
			Value:    "Credit Goal: Votre participation a ete enregistree. Si {team} marque des buts lors du prochain match, vous recevrez {reward_per_goal}{currency} de credit par but. {price}{curency}/souscription",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CREDIT_GOAL_ALREADY_SUB,
			Value:    "Credit Goal: Desole, vous avez deja une equipe favorite pour le prochain match de {league}. Vous pourrez choisir une autre equipe pour la journee suivante au {sc}.",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CREDIT_GOAL_UNVALID_SUB,
			Value:    "Credit Goal: Desole, votre choix n est pas valide. Renvoyez par SMS au {sc} votre equipe favorite pour le prochain match. {price}{currency}/souscription",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CREDIT_GOAL_MATCH_END_PAYOUT,
			Value:    "Credit Goal: Felicitations! Le score final du match {home} - {away} est {score}. Votre compte va etre credite de {price}{currency} [CFA]",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CREDIT_GOAL_MATCH_INCENTIVE,
			Value:    "Credit Goal: la prochaine journee de {league} va commencer. Soutenez votre equipe et gagnez du credit. Envoyez au {sc} votre equipe. {price}{currency}/souscription",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_SUB,
			Value:    "Prono: Faites vos pronostics pour {home} - {away}. Envoyez au {sc}: {league}[12] plus [1] pour {home}, [2] pour {away} ou [3] pour un nul. {price}{currency}/SMS",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_SUB_BET_WIN,
			Value:    "Prono: Votre pronostic a ete enregistre avec une victoire de: {winner}. Votre compte a ete debite de {price}{currency} [CFA]. Bonne Chance!",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_SUB_BET_DRAW,
			Value:    "Prono: Votre pronostic est un match nul pour le match {home} - {away}. Votre compte a ete debite de {price}{currency} [CFA]. Bonne Chance!",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_UNVALID_SUB,
			Value:    "Prono: Desole, votre choix n est pas valide. Renvoyez votre choix au {sc} ou le nom de votre competition. {price}{currency}/SMS",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_SUB_REJECT_MATCH_END,
			Value:    "Prono: Desole, cette competition est terminee. Envoyez le nom de votre competition au {sc} pour recevoir les invitations a pronostiquer. {price}{currency}/SMS",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_SUB_REJECT_MATCH_STARTED,
			Value:    "Prono: Desole, vous ne pouvez plus faire de pronostics sur ce match. SMS Confirmation Retentez votre chance Free of charge SMS Text Template: G21(NEW) pour le prochain match {price}{currency}/SMS",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_MATCH_END_WINNER_AIRTIME,
			Value:    "Prono: Bravo! Bon pronostic pour le match {home} - {away}, vous avez gagne {reward}{currency} de credit de com par tirage au sort. Votre compte sera credite sous 48h.",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_MATCH_END_WINNER_LOTERY,
			Value:    "Bravo! Vous avez fait un bon pronostic pour le match {home} - {away}, vous venez de gagner un lot, aller à la boutique Orange pour les obtenir.",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_MATCH_END_LUCKY_LOSER,
			Value:    "Prono: Vous avez fait un bon pronostic, mais vous n avez pas ete tire au sort pour gagner {reward}{currency}. Retentez votre chance pour le prochain match au {sc}. {price}{currency}/SMS",
		},
		{
			Category: ACT_NOTIFICATION,
			Channel:  ACT_SMS,
			Name:     SMS_PREDICT_MATCH_END_LOSER_NOTIF,
			Value:    "Prono: Desole! Votre pronostic n a pas ete le bon pour le match {home}-{away}. Retentez votre chance pour les prochains matchs .Bonne chance!",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_TEAM_SUB,
			Value:    "SMS Alerte: Vous avez souscrit avec succes pour suivre votre equipe favorite! Jusqu au {date}/{month} vous recevrez toutes les infos en direct. {price}{currency}/souscription",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_TEAM_UNVALID_SUB,
			Value:    "SMS Alerte: Desole, votre choix n est pas valide. Envoyez de nouveau au {sc} le nom de votre equipe pour obtenir toutes les infos. {price}{currency}/souscription",
		},
		{
			Category: ACT_NOTIFICATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_TEAM_EXPIRE_SUB,
			Value:    "SMS Alerte: Desole mais votre abonnement pour suivre votre equipe vient de se terminer. Pour le renouveler, envoyer au {sc} le nom de l equipe a suivre. {price}{currency}",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_COMPETITION_SUB,
			Value:    "SMS Alerte: Vous avez souscrit avec succes pour suivre {league}! Jusqu au {day}/{month} vous recevrez toutes les infos en direct. {price}{currency}/souscription",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_COMPETITION_UNVALID_SUB,
			Value:    "SMS Alerte: Desole, mais votre choix n est pas valide. Renvoyez au {sc} la competition que vous voulez suivre en direct. {price}{currency}/SMS",
		},
		{
			Category: ACT_NOTIFICATION,
			Channel:  ACT_SMS,
			Name:     SMS_FOLLOW_COMPETITION_EXPIRE_SUB,
			Value:    "SMS Alerte: Desole mais votre abonnement pour suivre votre competition vient de se terminer. Pour le renouveler, envoyer au {sc} le nom de l equipe a suivre. {price}{currency}/SMS",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_CONFIRMATION,
			Value:    "Info Services FFC: [#157#] USSD /[981] SMS Alerte Equipe /[9009] - SMS Alerte Competition /[###] SMS Alerte a la demande /[944] Credit Goal /[985] Pronostic",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_INFO,
			Value:    "Info Services FFC: [#157#] USSD /[981] SMS Alerte Equipe /[9009] - SMS Alerte Competition /[###] SMS Alerte a la demande /[944] Credit Goal /[985] Pronostic",
		},
		{
			Category: ACT_CONFIRMATION,
			Channel:  ACT_SMS,
			Name:     SMS_STOP,
			Value:    "Services FFC: Vous venez d envoyer STOP. Orange vous remercie d avoir utiliser nos services! Visitez www.orange.cm pour en savoir plus sur Football Fan Club",
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
			Name:       ACT_FOLLOW_COMPETITION,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
		{
			CountryID:  1,
			Name:       ACT_FOLLOW_TEAM,
			PublishAt:  time.Now(),
			UnlockedAt: time.Now(),
			IsUnlocked: false,
		},
	}

	/**
	1.  #101#36#1 : Live match (Confirm Message)
	2.  #101#36#2 : Flash News (Confirm Message)
	3.  #101#36#3 : Crédit Goal (Confirm Message)
	4.  #101#36#4 : Champ. Mali (Confirm Message)
	5.  #101#36#5 : Prédiction (Confirm Message)
	6.  #101#36#6 : SMS Alerte (Confirm Message)
	7.  #101#36#7 : Kit Foot (Confirm Message)
	8.  #101#36#8 : Foot Europe (Free Access)
	9.  #101#36#9 : Suiv
	**/

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
