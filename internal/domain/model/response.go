package model

type ApiFbResponse struct {
	League  *LeagueResponse    `json:"league"`
	Country *CountryResponse   `json:"country"`
	Seasson *[]SeassonResponse `json:"seasons"`
}

type LeagueResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
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

type FixturesResponse struct {
	ID                 int          `json:"id"`
	TimeZone           string       `json:"timezone"`
	Date               string       `json:"date"`
	Timestamp          int          `json:"timestamp"`
	Events             bool         `json:"events"`
	Lineups            bool         `json:"lineups"`
	StatisticsFixtures bool         `json:"statistics_fixtures"`
	StatisticsPlayers  bool         `json:"statistics_players"`
	Teams              TeamResponse `json:"teams"`
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
