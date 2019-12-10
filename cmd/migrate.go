package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations to the database",
	Long:  `Apply the database migrations from ./schema/postgres directory`,
	Run: func(cmd *cobra.Command, args []string) {
		apply()
	},
}

func apply() {
	url := viper.GetString("database.url")
	log.Infof("Rollup migrations...\n")
	log.Infof("Database URL: %s", url)
	m, err := migrate.New("file://schema/postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}