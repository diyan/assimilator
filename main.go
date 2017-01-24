package main

import (
	"os"

	"github.com/diyan/assimilator/web"
)

// TODO implement web.GetApp(), cron.GetApp(), smtp.GetApp() function
//   see https://docs.sentry.io/server/cli/run/

func main() {
	e := web.GetApp()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	e.Logger.Fatal(e.Start(port))
}
