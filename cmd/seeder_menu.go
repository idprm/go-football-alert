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
    <a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match">Live Match</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=flash-news&title=Flash+News">Flash News</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Champ. Mali</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=sms-alerte&title=SMS+Alerte">SMS Alerte</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot&title=Kit+Foot">Kit Foot</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Foot Europe</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-afrique&title=Foot+Afrique">Foot Afrique</a>
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
		<a href="{{.url}}/{{.version}}/ussd/buy?slug={{.slug}}&code={{.code}}&title={{.title}}" key="1">Oui</a>
        <br/>
		<a href="{{.url}}/{{.version}}/ussd/q?slug=flash-news" key="0">Ecran Précédent</a>
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
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}">{{.title}}</a>
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
    <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a>
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
    <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
  </page>
</pages>
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Live match",
		Slug:      "lm",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Live Match">
    <page>
        Live Match<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-live-match">Live Match</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-schedule">Schedule</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-lineup">Line Up</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=lm-display">Display Live match</a>
        <br />
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Live match",
		Slug:      "lm-live-match",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Live Match">
	<page>
		Live Match<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a>
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
		{{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a>
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
		{{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a>
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
		{{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a>
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "LIVEMATCH",
		Name:      "Display Live match",
		Slug:      "lm-display-livematch",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Display Live match">
	<page>
	Display Live match<br/>
	<br/>
	{{.paginate}}
    <a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match" key="0">Ecran Précédent</a>
    <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "FLASHNEWS",
		Name:      "Flash News",
		Slug:      "flash-news",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Flash News">
	<page>
		Flash News {{.date}}<br/>
		{{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "CREDITGOAL",
		Name:      "Crédit Goal",
		Slug:      "credit-goal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		Crédit Goal {{.date}}<br/>
        {{.data}}
		<br/>
		{{.paginate}}
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
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-results">Results</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-standings">Standings</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-schedule">Schedule</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-team">Team</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-credit-score">Crédit Score</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-credit-goal">Crédit Goal</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-sms-alerte">SMS Alerte</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali-sms-alerte-equipe">SMS Alerte Equipe</a>
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Results",
		Slug:      "champ-mali-results",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Results">
	<page>
		Results {{.date}}<br/>
        {{.data}}
		<br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Standings",
		Slug:      "champ-mali-standings",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Standings">
	<page>
		Standings {{.date}}<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
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
<pages descr="Schedule">
	<page>
		Schedule {{.date}}<br/>
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
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
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
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
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali" key="0">Ecran Précédent</a>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali Crédit Goal",
		Slug:      "champ-mali-credit-goal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		Crédit Goal {{.date}}<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Prev</a>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali SMS Alerte",
		Slug:      "champ-mali-sms-alerte",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="SMS Alerte">
	<page>
		SMS Alerte {{.date}}<br/>
        {{.data}}
        <br />
    	<a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "Champ. Mali SMS Alerte Equipe",
		Slug:      "champ-mali-sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		SMS Alerte Equipe {{.date}}<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a>
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
        {{.data}}
		<br />
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
        {{.data}}
		<br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}" key="0">Ecran Précédent</a>
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
        {{.data}}
		<br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=prediction&title=Prédiction">Prev</a>
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
        {{.data}}
		<br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=prediction&title=Prédiction">Prev</a>
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "SMS Alerte",
		Slug:      "sms-alerte",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="SMS Alerte">
    <page>
        SMS Alerte {{.date}}<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot">Kit Foot</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe">Europe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-afrique">Afrique</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-international">Foot International</a>
        <br />
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
		{{.data}}
        <br />
		{{.paginate}}
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
		{{.data}}
        <br />
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot" key="0">Ecran Précédent</a>
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
        <br />
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot" key="0">Ecran Précédent</a>
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
		{{.data}}
		<br />
		{{.paginate}}
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
        {{.data}}
		<br />
		{{.paginate}}
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
		{{.data}}
		<br />
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "SMSALERTE",
		Name:      "SMS Alerte Equipe",
		Slug:      "sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="SMS Alerte Equipe">
	<page>
		SMS Alerte Equipe {{.date}}<br/>
        {{.data}}
		<br />
        <a href="{{.url}}/{{.version}}/ussd/" key="00">Accueil</a>
	</page>
</pages>		
		`,
	},
}
