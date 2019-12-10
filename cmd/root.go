package cmd

import (
	"github.com/dairlair/sentimentd/pkg/app"
	"github.com/dairlair/sentimentd/pkg/infrastructure/helpers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var (
	rootCmd = &cobra.Command{
		Use:   "sentimentd",
		Short: "short",
		Long:  `long`,
	}
	application *app.App
	console helpers.Console
)

func init() {
	configureViper()
	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}
	// Config is read, lest create application...
	config := app.Config{
		Database: struct{ URL string }{
			URL: viper.GetString("database.url"),
		},
	}
	application = app.NewApp(config)
	console = helpers.NewConsole(os.Stdout)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func configureViper() {
	viper.AutomaticEnv()
	viper.SetConfigName("sentimentd")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/sentimentd")
	viper.AddConfigPath("$HOME/.sentimentd")
	viper.AddConfigPath("./")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("database.url", "postgres://sentimentd:sentimentd@sentimentd:5432/sentimentd?sslmode=disable")
}
