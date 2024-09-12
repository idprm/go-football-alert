package cmd

import (
	"fmt"
	"sync"

	"github.com/idprm/go-football-alert/internal/logger"
	"github.com/spf13/cobra"
	loggerDb "gorm.io/gorm/logger"
)

var consumerMOCmd = &cobra.Command{
	Use:   "mo",
	Short: "Consumer MO Service CLI",
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
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_MO_EXCHANGE, true, RMQ_MO_QUEUE)

		messagesData, errSub := rmq.Subscribe(1, false, RMQ_MO_QUEUE, RMQ_MO_EXCHANGE, RMQ_MO_QUEUE)
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
				p.MO(&wg, d.Body)
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
