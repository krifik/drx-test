package main

import (
	"flag"

	"github.com/krifik/test-drx/app"
	_ "github.com/krifik/test-drx/docs"
	"github.com/krifik/test-drx/exception"
)

// @title Mangojek API Docs
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	flag.Parse()
	if arg := flag.Arg(0); arg != "" {
		app.InitializeDB()
		return
	}
	app := app.InitializedApp()
	// Start App
	err := app.Listen(":2000")
	exception.PanicIfNeeded(err)
}
