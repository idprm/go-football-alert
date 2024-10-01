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
<pages>
  <page>
    Orange Football Club, votre choix:<br/>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=lm&title=Live+Match">Live Match</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=flash-news&title=Flash+News">Flash News</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=credit-goal&title=Crédit+Goal">Crédit Goal</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Champ. Mali</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=prediction&title=Prédiction">Prédiction</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=sms-alerte&title=SMS+Alerte">SMS Alerte</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot&title=Kit+Foot">Kit Foot</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Foot Europe</a>
    <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-afrique&title=Foot+Afrique">Foot Afrique</a>
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
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
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
<pages descr="Confirm">
	<page nav="stop">
		Confirm<br/>
        To signup to {{.service}} charging {{.price}} per SMS. Please reply with YES to confirm or NO to decline
        <a href="{{.url}}/{{.version}}/ussd/buy?slug={{.slug}}&code={{.code}}&action=yes">Yes</a>
        <a href="{{.url}}/{{.version}}/ussd/buy?slug={{.slug}}&code={{.code}}&action=no">No</a>
        <br/>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
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
<pages descr="Success">
	<page>
		Success<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}">{{.title}}</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
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
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
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
    <a href="{{.url}}/{{.version}}/ussd/q?slug={{.slug}}">Dos</a>
	<a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
  </page>
</pages>
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Live match",
		Slug:      "lm",
		IsConfirm: true,
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
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
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
    <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "creditgoal",
		Name:      "Crédit Goal",
		Slug:      "credit-goal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		Crédit Goal<br/>
        {{.data}}
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Champ. Mali",
		Slug:      "champ-mali",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Champ. Mali">
    <page>
        Champ. Mali<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-results">Results</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-standings">Standings</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-schedule">Schedule</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-team">Team</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-credit-score">Crédit Score</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-credit-goal">Crédit Goal</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-sms-alerte">SMS Alerte</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-sms-alerte-equipe">SMS Alerte Equipe</a>
		<br/>
		{{.paginate}}
		<a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Results",
		Slug:      "champ-results",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Results">
	<page>
		Results<br/>
        {{.data}}
		<br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>			
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Standings",
		Slug:      "champ-standings",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Standings">
	<page>
		Standings<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Schedule",
		Slug:      "champ-schedule",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Schedule">
	<page>
		Schedule<br/>
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>	
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Team",
		Slug:      "champ-team",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Team">
	<page>
		Team<br/>
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Crédit Score",
		Slug:      "champ-credit-score",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Score">
	<page>
		Crédit Score<br/>
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">DOs</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "Crédit Goal",
		Slug:      "champ-creditgoal",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		Crédit Goal<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Prev</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "SMS Alerte",
		Slug:      "champ-sms-alerte",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		SMS Alerte<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Prev</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-team",
		Name:      "SMS Alerte Equipe",
		Slug:      "champ-sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Crédit Goal">
	<page>
		SMS Alerte Equipe<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-mali&title=Champ.+Mali">Prev</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "prediction",
		Name:      "Prédiction",
		Slug:      "prediction",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Prédiction">
	<page>
		Prédiction<br/>
        {{.data}}
		<form action="{{ .url }}/v1/ussd/q/?slug={{.slug}}">
			<entry kind="digits" var="package">
				<prompt></prompt>
			</entry>
		</form>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "SMS Alerte",
		Slug:      "sms-alerte",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="SMS Alerte">
    <page>
        SMS Alerte<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-results">Kit Foot</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-standings">Europe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-schedule">Afrique</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-team">SMS Alerte Equipe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=champ-team">Foot International</a>
        <br />
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Kit Foot",
		Slug:      "sms-kit-foot",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Kit Foot">
    <page>
        Kit Foot<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot-champ-mali">Alerte Champ. Mali + Equipe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot-premier-league">Alerte Premier League + Equipe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot-la-liga">Alerte La Liga + Equipe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot-ligue-1">Alerte Ligue 1 + Equipe</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=kit-foot-serie-a">Alerte Serie A + Equipe</a>
        <br />
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Europe",
		Slug:      "sms-foot-europe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Afrique",
		Slug:      "sms-foot-afrique",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "SMS Alerte Equipe",
		Slug:      "sms-alerte-equipe",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Foot International",
		Slug:      "sms-foot-international",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Kit Foot",
		Slug:      "kit-foot",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte Champ. Mali + Equipe",
		Slug:      "kit-foot-champ",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte Premier League + Equipe",
		Slug:      "kit-foot-premier-league",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte La Liga + Equipe",
		Slug:      "kit-foot-la-liga",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte Ligue 1 + Equipe",
		Slug:      "kit-foot-ligue-1",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte Serie A + Equipe",
		Slug:      "kit-foot-serie-a",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Alerte Bundesligue + Equipe",
		Slug:      "kit-foot-bundesligue",
		IsConfirm: true,
		IsActive:  true,
		TemplateXML: `
		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Foot Europe",
		Slug:      "foot-europe",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Foot Europe">
    <page>
        Foot Europe<br/>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-champion-league">Champion League</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-premier-league">Premier League</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-la-liga">La Liga</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-ligue-1">Ligue 1</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-l-europa">L. Europa</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-serie-a">Serie A</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-bundesligua">Bundesligua</a>
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe-champ-portugal">Champ Portugal</a>
        <br />
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Champion League",
		Slug:      "foot-europe-champion-league",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Champion League">
    <page>
        Champion League<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Premier League",
		Slug:      "foot-europe-premier-league",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Premier League">
    <page>
        Premier League<br/>
        {{.data}}
        <br />
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "La Liga",
		Slug:      "foot-europe-la-liga",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="La Liga">
    <page>
        La Liga<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Ligue 1",
		Slug:      "foot-europe-ligue-1",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Ligue 1">
    <page>
        Ligue 1<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "L. Europa",
		Slug:      "foot-europe-l-europa",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="L. Europa">
    <page>
        L. Europa<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
    </page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Serie A",
		Slug:      "foot-europe-serie-a",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Serie A">
	<page>
		Serie A<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Bundesligua",
		Slug:      "foot-europe-bundesligua",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Bundesligua">
	<page>
		Bundesligua<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Champ Portugal",
		Slug:      "foot-europe-champ-portugal",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Champ Portugal">
	<page>
		Champ Portugal<br/>
        {{.data}}
        <br/>
		{{.paginate}}
        <a href="{{.url}}/{{.version}}/ussd/q?slug=foot-europe&title=Foot+Europe">Dos</a>
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>		
		`,
	},
	{
		Category:  "follow-competition",
		Name:      "Foot Afrique",
		Slug:      "foot-africa",
		IsConfirm: false,
		IsActive:  true,
		TemplateXML: `
<?xml version="1.0" encoding="ISO-8859-1"?>
<!DOCTYPE pages SYSTEM "cellflash-1.3.dtd">
<pages descr="Foot Africa">
	<page>
		Foot Africa<br/>
        {{.data}}
        <br />
        <a href="{{.url}}/{{.version}}/ussd/">Accueil</a>
	</page>
</pages>			
		`,
	},
}
