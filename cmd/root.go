package cmd

import (
	"database/sql"
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
	PATH_LOG          string = getEnv("PATH_LOG")
	PATH_IMAGE        string = getEnv("PATH_IMAGE")
	API_FOOTBALL_URL  string = getEnv("API_FOOTBALL_URL")
	API_FOOTBALL_KEY  string = getEnv("API_FOOTBALL_KEY")
	API_FOOTBALL_HOST string = getEnv("API_FOOTBALL_HOST")
	URL_MT            string = getEnv("URL_MT")
	USER_MT           string = getEnv("USER_MT")
	PASS_MT           string = getEnv("PASS_MT")
)

const (
	RMQ_EXCHANGE_TYPE                         string = "direct"
	RMQ_DATA_TYPE                             string = "application/json"
	RMQ_USSD_EXCHANGE                         string = "E_USSD"
	RMQ_USSD_QUEUE                            string = "Q_USSD"
	RMQ_SMS_EXCHANGE                          string = "E_SMS"
	RMQ_SMS_QUEUE                             string = "Q_SMS"
	RMQ_MO_EXCHANGE                           string = "E_MO"
	RMQ_MO_QUEUE                              string = "Q_MO"
	RMQ_NEWS_EXCHANGE                         string = "E_NEWS"
	RMQ_NEWS_QUEUE                            string = "Q_NEWS"
	RMQ_SMS_ALERTE_EXCHANGE                   string = "E_SMS_ALERTE"
	RMQ_SMS_ALERTE_QUEUE                      string = "Q_SMS_ALERTE"
	RMQ_SMS_ACTU_EXCHANGE                     string = "E_SMS_ACTU"
	RMQ_SMS_ACTU_QUEUE                        string = "Q_SMS_ACTU"
	RMQ_SMS_PRONO_EXCHANGE                    string = "E_SMS_PRONO"
	RMQ_SMS_PRONO_QUEUE                       string = "Q_SMS_PRONO"
	RMQ_PREDICT_WIN_EXCHANGE                  string = "E_PREDICT_WIN"
	RMQ_PREDICT_WIN_QUEUE                     string = "Q_PREDICT_WIN"
	RMQ_CREDIT_GOAL_EXCHANGE                  string = "E_CREDIT_GOAL"
	RMQ_CREDIT_GOAL_QUEUE                     string = "Q_CREDIT_GOAL"
	RMQ_CREDIT_SCORE_EXCHANGE                 string = "E_CREDIT_SCORE"
	RMQ_CREDIT_SCORE_QUEUE                    string = "Q_CREDIT_SCORE"
	RMQ_REMINDER_48H_BEFORE_CHARGING_EXCHANGE string = "E_REMINDER_48H_BEFORE_CHARGING"
	RMQ_REMINDER_48H_BEFORE_CHARGING_QUEUE    string = "Q_REMINDER_48H_BEFORE_CHARGING"
	RMQ_REMINDER_AFTER_TRIAL_ENDS_EXCHANGE    string = "E_REMINDER_AFTER_TRIAL_ENDS"
	RMQ_REMINDER_AFTER_TRIAL_ENDS_QUEUE       string = "Q_REMINDER_AFTER_TRIAL_ENDS"
	RMQ_RENEWAL_EXCHANGE                      string = "E_RENEWAL"
	RMQ_RENEWAL_QUEUE                         string = "Q_RENEWAL"
	RMQ_RETRY_EXCHANGE                        string = "E_RETRY"
	RMQ_RETRY_QUEUE                           string = "Q_RETRY"
	RMQ_RETRY_UP_EXCHANGE                     string = "E_RETRY_UNDERPAYMENT"
	RMQ_RETRY_UP_QUEUE                        string = "Q_RETRY_UNDERPAYMENT"
	RMQ_NOTIF_EXCHANGE                        string = "E_NOTIF"
	RMQ_NOTIF_QUEUE                           string = "Q_NOTIF"
	RMQ_MT_EXCHANGE                           string = "E_MT"
	RMQ_MT_QUEUE                              string = "Q_MT"
	RMQ_PB_MO_EXCHANGE                        string = "E_POSTBACK_MO"
	RMQ_PB_MO_QUEUE                           string = "Q_POSTBACK_MO"
	ACT_USSD                                  string = "USSD"
	ACT_SMS                                   string = "SMS"
	ACT_CONFIRMATION                          string = "CONFIRMATION"
	ACT_NOTIFICATION                          string = "NOTIFICATION"
	ACT_MO                                    string = "MO"
	ACT_FIRSTPUSH                             string = "FIRSTPUSH"
	ACT_RENEWAL                               string = "RENEWAL"
	ACT_RETRY                                 string = "RETRY"
	ACT_SMS_ALERTE                            string = "SMS_ALERTE"
	ACT_CREDIT_GOAL                           string = "CREDIT_GOAL"
	ACT_PRONOSTIC                             string = "PRONOSTIC"
	ACT_PREDICT_WIN                           string = "PREDICT_WIN"
	ACT_SUB                                   string = "SUB"
	ACT_UNSUB                                 string = "UNSUB"
	ACT_USER_LOSES                            string = "USER_LOSES"
	ACT_PREDICTION                            string = "PREDICTION"
	ACT_SCRAPING                              string = "SCRAPING"
)

const (
	SMS_CREDIT_GOAL_SUB                  string = "CREDIT_GOAL_SUB"
	SMS_CREDIT_GOAL_ALREADY_SUB          string = "CREDIT_GOAL_ALREADY_SUB"
	SMS_CREDIT_GOAL_UNVALID_SUB          string = "CREDIT_GOAL_UNVALID_SUB"
	SMS_CREDIT_GOAL_MATCH_END_PAYOUT     string = "CREDIT_GOAL_MATCH_END_PAYOUT"
	SMS_CREDIT_GOAL_MATCH_INCENTIVE      string = "CREDIT_GOAL_MATCH_INCENTIVE"
	SMS_PREDICT_SUB                      string = "PREDICT_SUB"
	SMS_PREDICT_SUB_BET_WIN              string = "PREDICT_SUB_BET_WIN"
	SMS_PREDICT_SUB_BET_DRAW             string = "PREDICT_SUB_BET_DRAW"
	SMS_PREDICT_UNVALID_SUB              string = "PREDICT_UNVALID_SUB"
	SMS_PREDICT_SUB_REJECT_MATCH_END     string = "PREDICT_SUB_REJECT_MATCH_END"
	SMS_PREDICT_SUB_REJECT_MATCH_STARTED string = "PREDICT_SUB_REJECT_MATCH_STARTED"
	SMS_PREDICT_MATCH_END_WINNER_AIRTIME string = "PREDICT_MATCH_END_WINNER_AIRTIME"
	SMS_PREDICT_MATCH_END_WINNER_LOTERY  string = "PREDICT_MATCH_END_WINNER_LOTERY"
	SMS_PREDICT_MATCH_END_LUCKY_LOSER    string = "PREDICT_MATCH_END_LUCKY_LOSER"
	SMS_PREDICT_MATCH_END_LOSER_NOTIF    string = "PREDICT_MATCH_END_LOSER_NOTIF"
	SMS_FOLLOW_TEAM_SUB                  string = "FOLLOW_TEAM_SUB"
	SMS_FOLLOW_TEAM_ALREADY_SUB          string = "FOLLOW_TEAM_ALREADY_SUB"
	SMS_FOLLOW_TEAM_UNVALID_SUB          string = "FOLLOW_TEAM_UNVALID_SUB"
	SMS_FOLLOW_TEAM_EXPIRE_SUB           string = "FOLLOW_TEAM_EXPIRE_SUB"
	SMS_FOLLOW_TEAM_STOP                 string = "FOLLOW_TEAM_STOP"
	SMS_FOLLOW_COMPETITION_SUB           string = "FOLLOW_COMPETITION_SUB"
	SMS_FOLLOW_COMPETITION_ALREADY_SUB   string = "FOLLOW_COMPETITION_ALREADY_SUB"
	SMS_FOLLOW_COMPETITION_UNVALID_SUB   string = "FOLLOW_COMPETITION_UNVALID_SUB"
	SMS_FOLLOW_COMPETITION_EXPIRE_SUB    string = "FOLLOW_COMPETITION_EXPIRE_SUB"
	SMS_FOLLOW_COMPETITION_STOP          string = "FOLLOW_COMPETITION_STOP"
	SMS_FOLLOW_UNVALID_SUB               string = "FOLLOW_UNVALID_SUB"
	SMS_LIVE_MATCH_SUB                   string = "LIVE_MATCH_SUB"
	SMS_LIVE_MATCH_UNVALID_SUB           string = "LIVE_MATCH_UNVALID_SUB"
	SMS_FLASH_NEWS_SUB                   string = "FLASH_NEWS_SUB"
	SMS_PRONOSTIC_SAFE_SUB               string = "PRONOSTIC_SAFE_SUB"
	SMS_PRONOSTIC_COMBINED_SUB           string = "PRONOSTIC_COMBINED_SUB"
	SMS_PRONOSTIC_VIP_SUB                string = "PRONOSTIC_VIP_SUB"
	SMS_PRONOSTIC_SAFE_ALREADY_SUB       string = "PRONOSTIC_SAFE_ALREADY_SUB"
	SMS_PRONOSTIC_COMBINED_ALREADY_SUB   string = "PRONOSTIC_COMBINED_ALREADY_SUB"
	SMS_PRONOSTIC_VIP_ALREADY_SUB        string = "PRONOSTIC_VIP_ALREADY_SUB"
	SMS_PRONOSTIC_STOP                   string = "PRONOSTIC_STOP"
	SMS_PRONOSTIC_UNVALID_SUB            string = "PRONOSTIC_UNVALID_SUB"
	SMS_REMINDER_48H_BEFORE_CHARGING     string = "REMINDER_48H_BEFORE_CHARGING"
	SMS_REMINDER_AFTER_TRIAL_ENDS        string = "REMINDER_AFTER_TRIAL_ENDS"
	SMS_SUCCESS_CHARGING                 string = "SUCCESS_CHARGING"
	SMS_INFO                             string = "INFO"
	SMS_STOP                             string = "STOP"
	SMS_OTP                              string = "OTP"
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
	rootCmd.AddCommand(consumerUSSDCmd)
	rootCmd.AddCommand(consumerSMSCmd)
	rootCmd.AddCommand(consumerMOCmd)
	rootCmd.AddCommand(consumerNewsCmd)
	rootCmd.AddCommand(consumerSMSAlerteCmd)
	rootCmd.AddCommand(consumerSMSActuCmd)
	rootCmd.AddCommand(consumerSMSPronoCmd)
	rootCmd.AddCommand(consumerCreditGoalCmd)
	rootCmd.AddCommand(consumerCreditScoreCmd)
	rootCmd.AddCommand(consumerPredictWinCmd)
	rootCmd.AddCommand(consumerReminder48HBeforeChargingCmd)
	rootCmd.AddCommand(consumerReminderAfterTrialEndsCmd)
	rootCmd.AddCommand(consumerRenewalCmd)
	rootCmd.AddCommand(consumerRetryCmd)
	rootCmd.AddCommand(consumerRetryUnderpaymentCmd)
	rootCmd.AddCommand(consumerMTCmd)
	rootCmd.AddCommand(consumerPostbackMOCmd)

	/**
	 * Publisher Scraping service
	 */
	rootCmd.AddCommand(publisherScrapingMasterCmd)
	rootCmd.AddCommand(publisherScrapingFixturesCmd)
	rootCmd.AddCommand(publisherScrapingLiveMatchesCmd)
	rootCmd.AddCommand(publisherScrapingPredictionCmd)
	rootCmd.AddCommand(publisherScrapingNewsCmd)
	rootCmd.AddCommand(publisherReportCmd)
	rootCmd.AddCommand(publisherCreditGoalCmd)
	rootCmd.AddCommand(publisherPredictWinCmd)
	rootCmd.AddCommand(publisherReminder48HBeforeChargingCmd)
	rootCmd.AddCommand(publisherReminderAfterTrialEndsCmd)
	rootCmd.AddCommand(publisherRenewalCmd)
	rootCmd.AddCommand(publisherRetryCmd)
	rootCmd.AddCommand(publisherRetryUnderpaymentCmd)

	// Test command
	rootCmd.AddCommand(consumerTestLeagueCmd)
	rootCmd.AddCommand(consumerTestTeamCmd)
	rootCmd.AddCommand(consumerTestFixtureCmd)
	rootCmd.AddCommand(consumerTestPredictionCmd)
	rootCmd.AddCommand(consumerTestStandingCmd)
	rootCmd.AddCommand(consumerTestLineupCmd)
	rootCmd.AddCommand(consumerTestNewsCmd)
	rootCmd.AddCommand(consumerTestBalanceCmd)
	rootCmd.AddCommand(consumerTestChargeCmd)
	rootCmd.AddCommand(consumerTestChargeFailedCmd)
	rootCmd.AddCommand(consumerTestUpdateFalseCmd)
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

func connectSqlDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", URI_MYSQL)
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
