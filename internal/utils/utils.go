package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateTrxId() string {
	id := uuid.New()
	return id.String()
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func EscapeChar(res []byte) []byte {
	response := string(res)
	r := strings.NewReplacer("&lt;", "<", "&gt;", ">")
	result := r.Replace(response)
	return []byte(result)
}

func IsSMSSuccess(v string) bool {
	return strings.HasPrefix(v, "Success")
}

func FormatFR(t time.Time) string {
	return fmt.Sprintf("%s %d %s.",
		days[t.Weekday()], t.Day(), months[t.Month()-1],
	)
}

func FormatFROnlyMonth(t time.Time) string {
	return months[t.Month()-1]
}

var days = [...]string{
	"Dimanche", "Lundi", "Mardi", "Mercredi", "Jeudi", "Vendredi", "Samedi"}

var months = [...]string{
	"Janvier", "Fevrier", "Mars", "Avril", "Mai", "Juin",
	"Juillet", "Aout", "Septembre", "Octobre", "Novembre", "Decembre",
}
