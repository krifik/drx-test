package app

import (
	"fmt"
	"log"
	"os"

	"github.com/krifik/test-drx/config"
	"github.com/krifik/test-drx/module"
	"github.com/krifik/test-drx/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/urfave/cli"
)

func InitializedApp() *fiber.App {
	configuration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configuration)
	productModule := module.NewProductModule(database)
	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	// Setup Routing
	routes.Route(app, productModule)
	// Start the server on port 3000
	log.Fatal(app.Listen(":" + configuration.Get("SERVICE_PORT")))
	return app

}

func InitializeDB() {
	configration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configration)

	cmdApp := cli.NewApp()

	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(cli *cli.Context) error {
				// migration function
				config.NewRunMigration(database)
				fmt.Println("================ migrated successfully ================")
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(cli *cli.Context) error {
				// seeding function
				config.NewRunSeed(database)
				fmt.Println("================ seeded successfully ================")
				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
