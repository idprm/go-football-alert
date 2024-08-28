package cmd

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
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
	RMQ_HOST          string = getEnv("RMQ_HOST")
	RMQ_USER          string = getEnv("RMQ_USER")
	RMQ_PASS          string = getEnv("RMQ_PASS")
	RMQ_PORT          string = getEnv("RMQ_PORT")
	RMQ_VHOST         string = getEnv("RMQ_VHOST")
	RMQ_URL           string = getEnv("RMQ_URL")
	AUTH_SECRET       string = getEnv("AUTH_SECRET")
	PATH_STATIC       string = getEnv("PATH_STATIC")
	PATH_IMAGE        string = getEnv("PATH_IMAGE")
	API_FOOTBALL_URL  string = getEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = getEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = getEnv("API_FOOTBALL_HOST")
)

const (
	RMQ_EXCHANGE_TYPE     string = "direct"
	RMQ_DATA_TYPE         string = "application/json"
	RMQ_MO_EXCHANGE       string = "E_MO"
	RMQ_MO_QUEUE          string = "Q_MO"
	RMQ_RENEWAL_EXCHANGE  string = "E_RENEWAL"
	RMQ_RENEWAL_QUEUE     string = "Q_RENEWAL"
	RMQ_RETRY_EXCHANGE    string = "E_RETRY"
	RMQ_RETRY_QUEUE       string = "Q_RETRY"
	RMQ_NOTIF_EXCHANGE    string = "E_NOTIF"
	RMQ_NOTIF_QUEUE       string = "Q_NOTIF"
	RMQ_POSTBACK_EXCHANGE string = "E_POSTBACK"
	RMQ_POSTBACK_QUEUE    string = "Q_POSTBACK"
	RMQ_TRAFFIC_EXCHANGE  string = "E_TRAFFIC"
	RMQ_TRAFFIC_QUEUE     string = "Q_TRAFFIC"
	ACT_MO                string = "MO"
	ACT_RENEWAL           string = "RENEWAL"
	ACT_RETRY             string = "RETRY"
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
	 * Consumer service
	 */
	rootCmd.AddCommand(consumerMOCmd)

	/**
	 * Publisher Scraping service
	 */
	rootCmd.AddCommand(publisherScrapingCmd)
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

// Connect to rabbitmq
func connectRabbitMq() (rmqp.AMQP, error) {
	var rb rmqp.AMQP
	port, _ := strconv.Atoi(RMQ_PORT)
	rb.SetAmqpURL(RMQ_HOST, port, RMQ_USER, RMQ_PASS, RMQ_VHOST)
	errConn := rb.SetUpConnectionAmqp()
	if errConn != nil {
		return rb, errConn
	}
	return rb, nil
}

// Connect to redis
func connectRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(URI_REDIS)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
