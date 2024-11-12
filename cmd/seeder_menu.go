package cmd

import (
	"github.com/idprm/go-football-alert/internal/domain/entity"
)

var menus = []entity.Menu{
	{
		Category:  "none",
		Name:      "Home",
		Slug:      "home",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Orange Football Club">
  <page>
    Orange Football Club, votre choix:<br/>
	<br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Match+en+Direct">Match en Direct</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=flash-news&title=Actu+Flash">Actu Flash</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Champ. Mali</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=sms-alerte&title=SMS+Alerte">Alerte SMS</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot&title=Kit+Foot">Kit Foot</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Foot Europe</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-afrique&title=Foot+Afrique">Foot Afrique</a><br/>
	<a href="{{.url}}/{{.version}}/ussd/q?slug=foot-international&title=Foot+International">Foot International</a>
  </page>
</pages>`,
	},
	{
		Category:  "none",
		Name:      "Package",
		Slug:      "package",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="S'abonner">
	<page>
		{{.title}}<br/>
		<br/>
        {{.data}}
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Confirm",
		Slug:      "confirm",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="{{.title}}">
	<page nav="stop">
		{{.title}} <br/>
        {{.title}} Obtenez en direct toutes les informations sur votre mobile sur la {{.title}}. Confirmez-vous votre inscription aux alertes?<br/>
		Prix : {{.price}} / 1 {{.package}}.
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/buy?slug={{.slug}}&code={{.code}}&title={{.title}}" key="1">Oui</a>
        <br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=flash-news" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Success",
		Slug:      "success",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Succes">
	<page>
		Vous avez souscrit avec succes!<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}">{{.title}}</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Failed",
		Slug:      "failed",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="S'abonner">
	<page>
		{{.title}}<br/>
		<br/>
        {{.data}}
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Detail",
		Slug:      "detail",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages>
  <page descr="{{.title}}">
  	{{.title}}
	<br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a><br/>
	<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
  </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Not found",
		Slug:      "404",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages>
  <page>
  	Menu non trouvé<br/>
	<br/>
    <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
  </page>
</pages>
		`,
	},
	{
		Category:  "none",
		Name:      "Msisdn not found",
		Slug:      "msisdn_not_found",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages>
  <page>
  	Msisdn not found<br/>
	<br/>
    <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
  </page>
</pages>
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Match en Direct",
		Slug:      "lm",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Match en Direct">
    <page>
        Match en Direct<br/>
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-live-match">Match en Direct</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-schedule">Calendrier</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-lineup">Onze de depart</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-display">Statistiques</a><br/>
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Match en Direct",
		Slug:      "lm-live-match",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Match en Direct">
	<page>
		Match en Direct<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a><br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Schedule",
		Slug:      "lm-schedule",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Schedule">
	<page>
		Schedule<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a><br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Line Up",
		Slug:      "lm-line-up",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Lineup">
	<page>
		Lineup<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a><br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "match Stats",
		Slug:      "lm-match-stats",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Match Stats">
	<page>
		Match Stats<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a><br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Display Match en Direct",
		Slug:      "lm-display-livematch",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Display Match en Direct">
	<page>
	Display Match en Direct<br/>
	<br/>
	{{.paginate}}
	<br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a><br/>
    <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "FLASHNEWS",
		Name:      "Actu Flash",
		Slug:      "flash-news",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Actu Flash">
	<page>
		Actu Flash {{.date}}<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "CREDITGOAL",
		Name:      "But Gagnant",
		Slug:      "credit-goal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="But Gagnant">
	<page>
		But Gagnant {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali",
		Slug:      "champ-mali",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Champ. Mali">
    <page>
        Champ. Mali {{.date}}<br/>
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-results">Resultats</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-standings">Classement</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-schedule">Calendrier</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-credit-score">Crédit Score</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-credit-goal">But Gagnant</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-sms-alerte">Alerte SMS</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-sms-alerte-equipe">Alerte SMS Equipe</a><br/>
		<br/>
		{{.paginate}}<br/>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Resultats",
		Slug:      "champ-mali-results",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Resultats">
	<page>
		Resultats {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Classement",
		Slug:      "champ-mali-standings",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Classement">
	<page>
		Classement {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Schedule",
		Slug:      "champ-mali-schedule",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Calendrier">
	<page>
		Calendrier {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>	
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Team",
		Slug:      "champ-mali-team",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Team">
	<page>
		Team {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Crédit Score",
		Slug:      "champ-mali-credit-score",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Score">
	<page>
		Crédit Score {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali But Gagnant",
		Slug:      "champ-mali-credit-goal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="But Gagnant">
	<page>
		But Gagnant {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Prev</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Alerte SMS",
		Slug:      "champ-mali-sms-alerte",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Alerte SMS">
	<page>
		Alerte SMS {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
    	<a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Alerte SMS Equipe",
		Slug:      "champ-mali-sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="But Gagnant">
	<page>
		Alerte SMS Equipe {{.date}}<br/>
		<br/>
        {{.data}}
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "PREDICTION",
		Name:      "Prédiction",
		Slug:      "prediction",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Prédiction">
	<page>
		Prédiction {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "PREDICTION",
		Name:      "Safe of the Day",
		Slug:      "prediction-safe-of-the-day",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Safe of the Day">
	<page>
		Safe of the Day {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "PREDICTION",
		Name:      "Daily combined bets",
		Slug:      "prediction-daily-combined-bets",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Daily combined bets">
	<page>
		Daily combined bets {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=prediction&title=Prédiction">Prev</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "PREDICTION",
		Name:      "VIP Prono",
		Slug:      "prediction-vip-prono",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="VIP Prono">
	<page>
		VIP Prono {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=prediction&title=Prédiction">Prev</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Alerte SMS",
		Slug:      "sms-alerte",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Alerte SMS">
    <page>
        Alerte SMS {{.date}}<br/>
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot">Kit Foot</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe">Europe</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-afrique">Afrique</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-international">Foot International</a><br/>
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Kit Foot",
		Slug:      "kit-foot",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Kit Foot">
    <page>
        Kit Foot {{.date}}<br/>
		<br/>
		{{.data}}
        <br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Kit Foot By League",
		Slug:      "kit-foot-by-league",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="{{.title}}">
    <page>
        {{.title}} {{.date}}<br/>
		<br/>
		{{.data}}
        <br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Kit Foot By Team",
		Slug:      "kit-foot-by-team",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="{{.title}}">
    <page>
        {{.title}} {{.date}}<br/>
		{{.data}}
        <br/>
		{{.paginate}}
		<br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot" key="0">Ecran Précédent</a><br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Foot Europe",
		Slug:      "foot-europe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Foot Europe">
	<page>
		Foot Europe {{.date}}<br/>
		<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Foot Afrique",
		Slug:      "foot-afrique",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Foot Afrique">
	<page>
		Foot Afrique {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
		{{.paginate}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Foot International",
		Slug:      "foot-international",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Foot International">
	<page>
		Foot International {{.date}}<br/>
		<br/>
		{{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Alerte SMS Equipe",
		Slug:      "sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Alerte SMS Equipe">
	<page>
		Alerte SMS Equipe {{.date}}<br/>
		<br/>
        {{.data}}
		<br/>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
}
