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
	//
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
			Category:   "CREDIT-GOAL",
			Name:       "Credit Goal Daily",
			Code:       "CG1",
			Package:    "daily",
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
			Category:   "PREDICTION",
			Name:       "Prediction & Win Daily",
			Code:       "PW1",
			Package:    "daily",
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
			Category:   "FOLLOW-TEAM",
			Name:       "Follow Team Daily",
			Code:       "FT1",
			Package:    "daily",
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
			Category:   "FOLLOW-COMPETITION",
			Name:       "Follow Competition Daily",
			Code:       "FC",
			Package:    "daily",
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
			Category:   "CREDIT-GOAL",
			Name:       "Credit Goal Weekly",
			Code:       "CG7",
			Package:    "weekly",
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
			Category:   "PREDICTION",
			Name:       "Prediction & Win Weekly",
			Code:       "PW7",
			Package:    "weekly",
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
			Category:   "FOLLOW-TEAM",
			Name:       "Follow Team Weekly",
			Code:       "FT7",
			Package:    "weekly",
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
			Category:   "FOLLOW-COMPETITION",
			Name:       "Follow Competition Weekly",
			Code:       "FC",
			Package:    "weekly",
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
			Category:   "CREDIT-GOAL",
			Name:       "Credit Goal Monthly",
			Code:       "CG30",
			Package:    "monthly",
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
			Category:   "PREDICTION",
			Name:       "Prediction & Win Monthly",
			Code:       "PW30",
			Package:    "monthly",
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
			Category:   "FOLLOW-TEAM",
			Name:       "Follow Team Monthly",
			Code:       "FT30",
			Package:    "monthly",
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
			Category:   "FOLLOW-COMPETITION",
			Name:       "Follow Competition Monthly",
			Code:       "FC30",
			Package:    "monthly",
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
			Name:     SMS_FOLLOW_COMPETITION_INVALID_SUB,
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
			KeyPress: "#101#36#1#5",
			ParentID: 1,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       7,
			Name:     "Flash News",
			KeyPress: "#101#36#2",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       8,
			Name:     "Crédit Goal",
			KeyPress: "#101#36#3",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       9,
			Name:     "Champ. Mali",
			KeyPress: "#101#36#4",
			ParentID: 0,
			Child:    0,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       10,
			Name:     "Results",
			KeyPress: "#101#36#4#1",
			ParentID: 9,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       11,
			Name:     "Standings",
			KeyPress: "#101#36#4#2",
			ParentID: 9,
			Child:    2,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       12,
			Name:     "Schedule",
			KeyPress: "#101#36#4#3",
			ParentID: 9,
			Child:    3,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       13,
			Name:     "Team",
			KeyPress: "#101#36#4#4",
			ParentID: 9,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       14,
			Name:     "Crédit Score",
			KeyPress: "#101#36#4#5",
			ParentID: 9,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       15,
			Name:     "Crédit Goal",
			KeyPress: "#101#36#4#6",
			ParentID: 9,
			Child:    6,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       16,
			Name:     "SMS Alerte",
			KeyPress: "#101#36#4#7",
			ParentID: 9,
			Child:    7,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       17,
			Name:     "SMS Alerte Equipe",
			KeyPress: "#101#36#4#8",
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
			KeyPress: "#101#36#6#1",
			ParentID: 19,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       21,
			Name:     "Europe",
			KeyPress: "#101#36#6#2",
			ParentID: 19,
			Child:    2,
			Action:   "foot-europe",
			IsActive: true,
		},
		{
			ID:       22,
			Name:     "Afrique",
			KeyPress: "#101#36#6#3",
			ParentID: 19,
			Child:    3,
			Action:   "foot-afrique",
			IsActive: true,
		},
		{
			ID:       23,
			Name:     "SMS Alerte Equipe",
			KeyPress: "#101#36#6#4",
			ParentID: 19,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       24,
			Name:     "Foot International",
			KeyPress: "#101#36#6#5",
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
			KeyPress: "#101#36#7#1",
			ParentID: 25,
			Child:    1,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       27,
			Name:     "Alerte Premier League + Equipe",
			KeyPress: "#101#36#7#2",
			ParentID: 25,
			Child:    2,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       28,
			Name:     "Alerte La Liga + Equipe",
			KeyPress: "#101#36#7#3",
			ParentID: 25,
			Child:    3,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       29,
			Name:     "Alerte Ligue 1 + Equipe",
			KeyPress: "#101#36#7#4",
			ParentID: 25,
			Child:    4,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       30,
			Name:     "Alerte Serie A + Equipe",
			KeyPress: "#101#36#7#5",
			ParentID: 25,
			Child:    5,
			Action:   "",
			IsActive: true,
		},
		{
			ID:       31,
			Name:     "Alerte Bundesligue + Equipe",
			KeyPress: "#101#36#7#6",
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
			KeyPress: "#101#36#8#1",
			ParentID: 32,
			Child:    1,
			Action:   "champion-league",
			IsActive: true,
		},
		{
			ID:       34,
			Name:     "Premier League",
			KeyPress: "#101#36#8#2",
			ParentID: 32,
			Child:    2,
			Action:   "premier-league",
			IsActive: true,
		},
		{
			ID:       35,
			Name:     "La Liga",
			KeyPress: "#101#36#8#3",
			ParentID: 32,
			Child:    3,
			Action:   "la-liga",
			IsActive: true,
		},
		{
			ID:       36,
			Name:     "Ligue 1",
			KeyPress: "#101#36#8#4",
			ParentID: 32,
			Child:    4,
			Action:   "ligue-1",
			IsActive: true,
		},
		{
			ID:       37,
			Name:     "L. Europa",
			KeyPress: "#101#36#8#5",
			ParentID: 32,
			Child:    5,
			Action:   "l-europa",
			IsActive: true,
		},
		{
			ID:       38,
			Name:     "Serie A",
			KeyPress: "#101#36#8#6",
			ParentID: 32,
			Child:    6,
			Action:   "serie-a",
			IsActive: true,
		},
		{
			ID:       39,
			Name:     "Bundesligua",
			KeyPress: "#101#36#8#7",
			ParentID: 32,
			Child:    7,
			Action:   "bundesligua",
			IsActive: true,
		},
		{
			ID:       40,
			Name:     "Champ Portugal",
			KeyPress: "#101#36#8#8",
			ParentID: 32,
			Child:    8,
			Action:   "champ-portugal",
			IsActive: true,
		},
		{
			ID:       41,
			Name:     "Saudi League",
			KeyPress: "#101#36#8#9",
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
