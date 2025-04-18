package model

import (
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type WebResponse struct {
	Error       bool   `json:"error,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	Message     string `json:"message,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type SummaryResponse struct {
	StartDate      string      `json:"start_date"`
	EndDate        string      `json:"end_date"`
	TotalActiveSub int         `json:"total_active_sub"`
	TotalSub       int         `json:"total_sub"`
	TotalUnsub     int         `json:"total_unsub"`
	TotalRenewal   int         `json:"total_renewal"`
	TotalRevenue   float64     `json:"total_revenue"`
	Results        interface{} `json:"results"`
}

type ApiFbResponse struct {
	League  *LeagueResponse    `json:"league"`
	Country *CountryResponse   `json:"country"`
	Seasson *[]SeassonResponse `json:"seasons"`
}

type CountryResponse struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type SeassonResponse struct {
	Year    int    `json:"year"`
	Start   string `json:"start"`
	End     string `json:"end"`
	Current bool   `json:"current"`
}

type CoverageResponse struct {
	Standings   bool `json:"standings"`
	Players     bool `json:"players"`
	TopScorers  bool `json:"top_scorers"`
	TopAssists  bool `json:"top_assists"`
	TopCards    bool `json:"top_cards"`
	Injuries    bool `json:"injuries"`
	Predictions bool `json:"predictions"`
	Odds        bool `json:"odds"`
}

type LeagueResult struct {
	Results  int                 `json:"results"`
	Response []ResponseLeagueAPI `json:"response"`
}

type TeamResult struct {
	Results  int               `json:"results"`
	Response []ResponseTeamAPI `json:"response"`
}

type FixtureResult struct {
	Results  int                   `json:"results"`
	Response []ResponseFixturesAPI `json:"response"`
}

type PredictionResult struct {
	Results  int                     `json:"results"`
	Response []ResponsePredictionAPI `json:"response"`
}

type StandingResult struct {
	Results  int                   `json:"results"`
	Response []ResponseStandingAPI `json:"response"`
}

type LineupResult struct {
	Results  int                 `json:"results"`
	Response []ResponseLineupAPI `json:"response"`
}

type ResponseLeagueAPI struct {
	League  LeagueResp  `json:"league"`
	Country CountryResp `json:"country"`
}

type ResponseTeamAPI struct {
	Team TeamResp `json:"team"`
}

type ResponseFixturesAPI struct {
	Fixtures FixturesResponse `json:"fixture"`
	Teams    TeamResponse     `json:"teams"`
	League   LeagueResponse   `json:"league"`
	Goals    GoalResponse     `json:"goals"`
}

type ResponsePredictionAPI struct {
	Prediction PredictionResponse `json:"predictions"`
}

type ResponseStandingAPI struct {
	League struct {
		Name     string               `json:"name"`
		Country  string               `json:"country"`
		Standing [][]StandingResponse `json:"standings"`
	} `json:"league"`
}

type ResponseLineupAPI struct {
	Team struct {
		PrimaryID int    `json:"id"`
		Name      string `json:"name"`
	} `json:"team"`
	Formation string `json:"formation"`
}

type StandingResponse struct {
	Rank int `json:"rank"`
	Team struct {
		PrimaryID int    `json:"id"`
		Name      string `json:"name"`
	} `json:"team"`
	Points      int    `json:"points"`
	GoalsDiff   int    `json:"goalsDiff"`
	Group       string `json:"group"`
	Form        string `json:"form"`
	Status      string `json:"status"`
	Description string `json:"description"`
	All         struct {
		Played int `json:"played"`
		Win    int `json:"win"`
		Draw   int `json:"draw"`
		Lose   int `json:"lose"`
	} `json:"all"`
	UpdateAt string `json:"update"`
}

type FixturesResponse struct {
	ID        int    `json:"id"`
	TimeZone  string `json:"timezone"`
	Date      string `json:"date"`
	Timestamp int    `json:"timestamp"`
	Status    struct {
		Long    string `json:"long"`
		Short   string `json:"short"`
		Elapsed int    `json:"elapsed"`
	} `json:"status"`
}

type TeamResponse struct {
	Home HomeResponse `json:"home"`
	Away AwayResponse `json:"away"`
}

type HomeResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type AwayResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type LeagueResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Country string `json:"country,omitempty"`
	Logo    string `json:"logo,omitempty"`
	Flag    string `json:"flag,omitempty"`
	Season  int    `json:"season,omitempty"`
	Round   string `json:"round,omitempty"`
}

type GoalResponse struct {
	Home int `json:"home,omitempty"`
	Away int `json:"away,omitempty"`
}

type LeagueResp struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Logo string `json:"logo,omitempty"`
}

func (m *LeagueResp) GetNameWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, m.Name)
	return result
}

type CountryResp struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type PredictionResponse struct {
	Winner struct {
		PrimaryID int    `json:"id"`
		Name      string `json:"name"`
		Comment   string `json:"comment"`
	} `json:"winner"`
	Advice  string `json:"advice"`
	Percent struct {
		Home string `json:"home"`
		Draw string `json:"draw"`
		Away string `json:"away"`
	} `json:"percent"`
}

type TeamResp struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Logo    string `json:"logo,omitempty"`
	Founded int    `json:"founded,omitempty"`
	Country string `json:"country,omitempty"`
}

func (m *TeamResp) GetNameWithoutAccents() string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, m.Name)
	return result
}

type MaxfootRSSResponse struct {
	Channel struct {
		Item []MaxfootItem `xml:"item"`
	} `xml:"channel"`
}

type MaxfootItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type MadeInFootRSSResponse struct {
	Channel struct {
		Item []MadeInFootItem `xml:"item"`
	} `xml:"channel"`
}

type MadeInFootItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type AfricaTopSportsRSSResponse struct {
	Channel struct {
		Item []AfricaTopSportsItem `xml:"item"`
	} `xml:"channel"`
}

type AfricaTopSportsItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type FootMercatoSitemapResponse struct {
	Url struct {
		News []FootMercatoItem `xml:"news"`
	} `xml:"url"`
}

type FootMercatoItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"publication_date"`
	Keywords    string `xml:"keywords"`
}

type RmcSportRSSResponse struct {
	Channel struct {
		Item []MaxfootItem `xml:"item"`
	} `xml:"channel"`
}

type RmcSportItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type MobimiumNewsResponse struct {
	Url struct {
		News []MobimiumNewsItemResponse `xml:"news"`
	} `xml:"url"`
}

type MobimiumNewsItemResponse struct {
	Title           string `xml:"title"`
	PublicationDate string `xml:"publication_date"`
}

type MobimiumNewsRSSResponse struct {
	Channel struct {
		Item []MobimiumNewsItemRSS `xml:"item"`
	} `xml:"channel"`
}

type MobimiumNewsItemRSS struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (m *MobimiumNewsItemResponse) GetTitle() string {
	return m.Title
}

func (m *MobimiumNewsItemResponse) GetPublicationDate() string {
	return m.PublicationDate
}

type UnsubRequest struct {
	ID int `validate:"required" json:"id"`
}

func (m *UnsubRequest) GetId() int {
	return m.ID
}

type RevenueResponse struct {
	CreatedAt string  `json:"created_at"`
	Subject   string  `json:"subject"`
	Status    string  `json:"status"`
	Total     int     `json:"total"`
	Revenue   float64 `json:"revenue"`
}

type SummaryRevenueResponse struct {
	StartDate string      `json:"start_date"`
	EndDate   string      `json:"end_date"`
	Results   interface{} `json:"results"`
}
