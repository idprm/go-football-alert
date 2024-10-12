package main

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/cmd"
)

func main() {

	cmd.Execute()
	log.Println(time.Now())
}
