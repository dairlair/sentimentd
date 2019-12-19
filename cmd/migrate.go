package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Manage migrations",
	Long:  `Allows to apply or rollback migrations`,
}

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply migrations to the database",
	Long:  `Apply the database migrations from ./schema/postgres directory`,
	Run: func(cmd *cobra.Command, args []string) {
		execute(func (m *migrate.Migrate) error {
			return m.Up()
		})
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "Rollback migrations in the database",
	Run: func(cmd *cobra.Command, args []string) {
		execute(func (m *migrate.Migrate) error {
			return m.Down()
		})
	},
}

func execute(f func (*migrate.Migrate) error) {
	url := viper.GetString("database.url")
	m, err := migrate.New("file://schema/postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := f(m); err != nil {
		log.Fatal(err)
	}
}