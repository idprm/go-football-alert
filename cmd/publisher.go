package cmd

import (
	"time"

	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
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
