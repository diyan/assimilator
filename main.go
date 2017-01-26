package main

import (
	"fmt"
	"os"

	"github.com/diyan/assimilator/migrations"
	"github.com/diyan/assimilator/web"
)

// TODO implement web.GetApp(), cron.GetApp(), smtp.GetApp() function
//   see https://docs.sentry.io/server/cli/run/
func main() {
	fmt.Println("Upgrading database schema, please wait...")
	migrations.UpgradeDB()
	fmt.Print("Database is up to date. Starting web app...")
	e := web.GetApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	e.Logger.Fatal(e.Start(port))
}
