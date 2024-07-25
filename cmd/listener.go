package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/domain/repository"
	"github.com/idprm/go-football-alert/internal/handler"
	"github.com/idprm/go-football-alert/internal/services"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Listener Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// connect db
		db, err := connectDb()
		if err != nil {
			panic(err)
		}

		// DEBUG ON CONSOLE
		db.Logger = logger.Default.LogMode(logger.Info)

		// TODO: Add migrations
		db.AutoMigrate(
			&entity.Ussd{},
		)

		r := routeUrlListener(db)
		log.Fatal(r.Listen(":" + APP_PORT))

	},
}

func routeUrlListener(db *gorm.DB) *fiber.App {
	app := fiber.New()

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	app.Static(PATH_STATIC, path+"/public")
	app.Use(cors.New())

	ussdRepo := repository.NewUssdRepository(db)
	ussdService := services.NewUssdService(ussdRepo)
	ussdHandler := handler.NewUssdHandler(ussdService)

	v1 := app.Group(API_VERSION)

	// callback
	ussd := v1.Group("ussd")
	ussd.Post("/callback", ussdHandler.Callback)
	ussd.Post("/event", ussdHandler.Event)

	return app
}
