package model

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

type FixtureResult struct {
	Results  int                   `json:"results"`
	Response []ResponseFixturesAPI `json:"response"`
}

type ResponseFixturesAPI struct {
	Fixtures FixturesResponse `json:"fixture"`
	Teams    TeamResponse     `json:"teams"`
	League   LeagueResponse   `json:"league"`
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
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
	Round   string `json:"round"`
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
