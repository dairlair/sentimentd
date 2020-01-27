package cmd

import (
	"github.com/dairlair/sentimentd/pkg/application"
	"github.com/dairlair/sentimentd/pkg/interface/cli"
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
	app *application.App
)

func init() {
	configureViper()
	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}
	// Config is read, lest create application...
	config := application.Config{
		Database: struct{ URL string }{
			URL: viper.GetString("database.url"),
		},
	}
	app = application.NewApp(config)
	app.Init()

	cmdFactory := cli.NewCommandsRunner(app, os.Stdin, os.Stdout, os.Stderr)
	rootCmd.AddCommand(cmdFactory.NewCmdPredict())
	rootCmd.AddCommand(cmdFactory.NewCmdBrain())
	rootCmd.AddCommand(cmdFactory.NewCmdTrain())
	rootCmd.AddCommand(cmdFactory.NewCmdListen())
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
