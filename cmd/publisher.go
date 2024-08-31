package cmd

import (
	"encoding/json"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/gorm"
)

var publisherScrapingCmd = &cobra.Command{
	Use:   "pub_scraping",
	Short: "Publisher Scraping Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_SCRAPING, timeNow) {

				// scheduleService.Update(false, ACT_CSV)

				go func() {
					scraping(db)
				}()
			}

			if scheduleService.IsUnlocked(ACT_SCRAPING, timeNow) {
				// scheduleService.Update(true, ACT_CSV)
			}

			time.Sleep(timeDuration * time.Minute)
		}
	},
}

func scraping(db *gorm.DB) {

}

var publisherRenewalCmd = &cobra.Command{
	Use:   "pub_renewal",
	Short: "Renewal CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect pgsql
		 */
		db, err := connectDb()
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

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RENEWAL_EXCHANGE, true, RMQ_RENEWAL_QUEUE)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			timeNow := time.Now().Format("15:04")

			scheduleRepo := repository.NewScheduleRepository(db)
			scheduleService := services.NewScheduleService(scheduleRepo)

			if scheduleService.IsUnlocked(ACT_RENEWAL, timeNow) {

				scheduleService.Update(
					&entity.Schedule{
						Name:       ACT_RENEWAL,
						IsUnlocked: false,
					},
				)

				go func() {
					populateRenewal(db, rmq)
				}()
			}

			time.Sleep(timeDuration * time.Minute)

		}
	},
}

func populateRenewal(db *gorm.DB, queue rmqp.AMQP) {
	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	subs := subscriptionService.Renewal()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.ServiceID = s.ServiceID
		sub.Msisdn = s.Msisdn
		sub.Channel = s.Channel
		sub.LatestKeyword = s.LatestKeyword
		sub.LatestSubject = s.LatestSubject
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		queue.IntegratePublish(RMQ_RENEWAL_EXCHANGE, RMQ_RENEWAL_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}
