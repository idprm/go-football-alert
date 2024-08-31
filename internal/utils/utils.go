package utils

import (
	"log"
	"os"
	"strings"

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

func EscapeChar(res []byte) []byte {
	response := string(res)
	r := strings.NewReplacer("&lt;", "<", "&gt;", ">")
	result := r.Replace(response)
	return []byte(result)
}

func IsSMSSuccess(v string) bool {
	return strings.HasPrefix(v, "Success")
}
