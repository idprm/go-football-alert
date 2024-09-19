package model

type WebResponse struct {
	Error       bool   `json:"error,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
	Message     string `json:"message,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
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

type FixturesResponse struct {
	ID        int    `json:"id"`
	TimeZone  string `json:"timezone"`
	Date      string `json:"date"`
	Timestamp int    `json:"timestamp"`
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

type CountryResp struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Logo string `json:"logo,omitempty"`
}

type TeamResp struct {
	ID      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Logo    string `json:"logo,omitempty"`
	Founded int    `json:"founded,omitempty"`
	Country string `json:"country,omitempty"`
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
