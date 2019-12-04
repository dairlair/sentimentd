package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply migrations to the database",
	Long:  `Apply the database migrations from ./schema/postgres directory`,
	Run: func(cmd *cobra.Command, args []string) {
		apply()
	},
}

func apply() {
	fmt.Println("Rollup migrations...")
	m, err := migrate.New(
		"file://schema/postgres",
		"postgres://sentimentd:sentimentd@localhost:5432/sentimentd?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}