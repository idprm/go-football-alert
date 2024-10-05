package cmd

import (
	"fmt"
	"sync"

	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/spf13/cobra"
	loggerDb "gorm.io/gorm/logger"
)

var consumerUSSDCmd = &cobra.Command{
	Use:   "ussd",
	Short: "Consumer USSD Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_USSD_EXCHANGE, true, RMQ_USSD_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_USSD_QUEUE, RMQ_USSD_EXCHANGE, RMQ_USSD_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.USSD(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerSMSCmd = &cobra.Command{
	Use:   "sms",
	Short: "Consumer SMS Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_SMS_EXCHANGE, true, RMQ_SMS_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_SMS_QUEUE, RMQ_SMS_EXCHANGE, RMQ_SMS_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.SMS(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerMTCmd = &cobra.Command{
	Use:   "mt",
	Short: "Consumer MT Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_MT_QUEUE, RMQ_MT_EXCHANGE, RMQ_MT_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.MT(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerRenewalCmd = &cobra.Command{
	Use:   "renewal",
	Short: "Consumer Renewal Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RENEWAL_EXCHANGE, true, RMQ_RENEWAL_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_RENEWAL_QUEUE, RMQ_RENEWAL_EXCHANGE, RMQ_RENEWAL_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Renewal(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerRetryCmd = &cobra.Command{
	Use:   "retry",
	Short: "Consumer Retry Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RETRY_EXCHANGE, true, RMQ_RETRY_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_RETRY_QUEUE, RMQ_RETRY_EXCHANGE, RMQ_RETRY_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Retry(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerPredictionCmd = &cobra.Command{
	Use:   "prediction",
	Short: "Consumer Prediction Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_PREDICTION_EXCHANGE, true, RMQ_PREDICTION_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_PREDICTION_QUEUE, RMQ_PREDICTION_EXCHANGE, RMQ_PREDICTION_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Prediction(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerCreditGoalCmd = &cobra.Command{
	Use:   "credit_goal",
	Short: "Consumer Credit Goal Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_CREDIT_EXCHANGE, true, RMQ_CREDIT_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_CREDIT_QUEUE, RMQ_CREDIT_EXCHANGE, RMQ_CREDIT_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.CreditGoal(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerNewsCmd = &cobra.Command{
	Use:   "news",
	Short: "Consumer News Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_NEWS_EXCHANGE, true, RMQ_NEWS_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_NEWS_QUEUE, RMQ_NEWS_EXCHANGE, RMQ_NEWS_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.News(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerFollowCompetitionCmd = &cobra.Command{
	Use:   "follow_competition",
	Short: "Consumer Follow Competition Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_FOLLOW_COMPETITION_EXCHANGE, true, RMQ_FOLLOW_COMPETITION_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_FOLLOW_COMPETITION_QUEUE, RMQ_FOLLOW_COMPETITION_EXCHANGE, RMQ_FOLLOW_COMPETITION_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.FollowCompetition(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerFollowTeamCmd = &cobra.Command{
	Use:   "follow_team",
	Short: "Consumer Follow Team Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq, err := connectRabbitMq()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = loggerDb.Default.LogMode(loggerDb.Info)

		/**
		 * SETUP LOG
		 */
		logger := logger.NewLogger()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_FOLLOW_TEAM_EXCHANGE, true, RMQ_FOLLOW_TEAM_QUEUE)
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MT_EXCHANGE, true, RMQ_MT_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_FOLLOW_TEAM_QUEUE, RMQ_FOLLOW_TEAM_EXCHANGE, RMQ_FOLLOW_TEAM_QUEUE)
		if errSub != nil {
			panic(errSub)
		}

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, rmq, logger)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.FollowTeam(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}
