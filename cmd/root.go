package cmd

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/db/migrations"
	"github.com/diyan/assimilator/log"
	"github.com/diyan/assimilator/web"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	// TODO consider PersistentFlags instead of just Flags
	RootCmd.Flags().IntP("port", "p", 3000, "Port to run the server on")
	RootCmd.Flags().StringP("db-url", "d", "postgres://postgres@localhost/sentry?sslmode=disable", "Url to connect to PostgreSQL database")
	// TODO fix order displayed in cli help
	RootCmd.Flags().String("initial-team", "ACME-Team", "Create an initial team inside Sentry DB with the given name")
	RootCmd.Flags().String("initial-project", "ACME", "Create an initial project for the above team (owner for both is the created admin")
	RootCmd.Flags().String("initial-key", "763a78a695424ed687cf8b7dc26d3161:763a78a695424ed687cf8b7dc26d3161", "Set a key for the above project so you can set DSN in your client, e.g. public:secret")
	RootCmd.Flags().String("initial-platform", "python", "Indicates a platform for the above initial project")

	viper.SetEnvPrefix("assimilator")
	viper.AutomaticEnv()
	// ASSIMILATOR_PORT env var is for user, PORT is for codegangsta/gin tool
	viper.BindEnv("port", "PORT")
	viper.BindPFlag("port", RootCmd.Flags().Lookup("port"))
	viper.BindPFlag("db_url", RootCmd.Flags().Lookup("db-url"))
	viper.BindPFlag("initial_team", RootCmd.Flags().Lookup("initial-team"))
	viper.BindPFlag("initial_project", RootCmd.Flags().Lookup("initial-project"))
	viper.BindPFlag("initial_key", RootCmd.Flags().Lookup("initial-key"))
	viper.BindPFlag("initial_platform", RootCmd.Flags().Lookup("initial-platform"))
}

var RootCmd = &cobra.Command{
	Use:   "assimilator",
	Short: "Assimilator is an attempt to port minimum valuable subset of Sentry from Python to the Golang",
	Long: `Assimilator is an attempt to port minimum valuable subset of Sentry from Python to the Golang.
        Put a multiline 
        rationale here`,
	RunE: func(cmd *cobra.Command, args []string) error {
		config := conf.Config{}
		if err := viper.Unmarshal(&config); err != nil {
			return errors.Wrap(err, "can not load configuration")
		}
		log.Init(config)
		logrus.Info("Upgrading database schema, please wait...")
		if err := migrations.UpgradeDB(config.DatabaseURL); err != nil {
			return err
		}
		logrus.Info("Database is up to date. Starting web app...")
		// TODO implement web.NewApp(), cron.NewApp(), smtp.NewApp() function
		//   see https://docs.sentry.io/server/cli/run/
		e := web.NewApp(config)
		if err := e.Start(fmt.Sprintf(":%d", config.Port)); err != nil {
			e.Logger.Fatal(err)
		}
		return nil
	},
}
