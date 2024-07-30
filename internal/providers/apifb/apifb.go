package apifb

import "github.com/idprm/go-football-alert/internal/utils"

var (
	URL_FOOTBALL string = utils.GetEnv("URL_FOOTBALL")
)

type ApiFb struct {
}

func NewApiFb() *ApiFb {
	return &ApiFb{}
}
