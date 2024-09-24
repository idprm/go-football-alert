package apifb

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/idprm/go-football-alert/internal/utils"
)

var (
	API_FOOTBALL_URL  string = utils.GetEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = utils.GetEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = utils.GetEnv("API_FOOTBALL_HOST")
)

type ApiFb struct {
}

func NewApiFb() *ApiFb {
	return &ApiFb{}
}

func (p *ApiFb) GetLeagues() ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/leagues", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("season", strconv.Itoa(time.Now().Year()))

	req.URL.RawQuery = q.Encode()

	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetTeams(leagueId int) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/teams", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("league", strconv.Itoa(leagueId))
	// q.Add("season", strconv.Itoa(time.Now().Year()))
	q.Add("season", "2024")

	req.URL.RawQuery = q.Encode()

	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetFixtures(leagueId int) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/fixtures", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("season", strconv.Itoa(time.Now().Year()))
	q.Add("league", strconv.Itoa(leagueId))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetPredictions(fixtureId int) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/predictions", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("fixture", strconv.Itoa(fixtureId))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetStandings(leagueId int) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/standings", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("season", strconv.Itoa(time.Now().Year()))
	q.Add("league", strconv.Itoa(leagueId))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetFixtureByLeague(leagueId int) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/fixtures", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("season", strconv.Itoa(time.Now().Year()))
	// primary league
	q.Add("league", strconv.Itoa(leagueId))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetFixtureByRounds(leagueId, seasonId int, current bool) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/fixtures/rounds", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("league", strconv.Itoa(leagueId))
	q.Add("season", strconv.Itoa(seasonId))
	q.Add("current", strconv.FormatBool(current))
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

func (p *ApiFb) GetFixtureByHeadToHead(h2h string) ([]byte, error) {
	req, err := http.NewRequest("GET", API_FOOTBALL_URL+"/fixtures/headtohead", nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	q := req.URL.Query()
	q.Add("h2h", h2h)
	q.Add("date", "")
	q.Add("league", "")
	q.Add("season", "")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("x-rapidapi-key", API_FOOTBALL_KEY)
	req.Header.Set("x-rapidapi-host", API_FOOTBALL_HOST)

	tr := &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}

// func (p *ApiFb) GetOdds() ([]byte, error) {

// }
