package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	APP_URL           string = getEnv("APP_URL")
	APP_PORT          string = getEnv("APP_PORT")
	APP_TZ            string = getEnv("APP_TZ")
	API_VERSION       string = getEnv("API_VERSION")
	URI_MYSQL         string = getEnv("URI_MYSQL")
	URI_REDIS         string = getEnv("URI_REDIS")
	AUTH_SECRET       string = getEnv("AUTH_SECRET")
	PATH_STATIC       string = getEnv("PATH_STATIC")
	PATH_IMAGE        string = getEnv("PATH_IMAGE")
	API_FOOTBALL_URL  string = getEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = getEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = getEnv("API_FOOTBALL_HOST")
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func init() {
	// setup timezone
	loc, _ := time.LoadLocation(APP_TZ)
	time.Local = loc
	/**
	 * Listener service
	 */
	rootCmd.AddCommand(listenerCmd)
	/**
	 * Scraper service
	 */
	rootCmd.AddCommand(scraperCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

// Connect to gorm mysql
func connectDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(URI_MYSQL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Connect to redis
// func connectRedis() (*redis.Client, error) {
// 	opts, err := redis.ParseURL(URI_REDIS)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return redis.NewClient(opts), nil
// }
