package cmd

import (
	"fmt"
	stan "github.com/nats-io/go-nats-streaming"
	"os"
	"strings"

	"github.com/dairlair/sentimentd/pkg/application"
	"github.com/dairlair/sentimentd/pkg/infrastructure/nats"
	"github.com/dairlair/sentimentd/pkg/interface/cli"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	queueCreator := func() stan.Conn { return getNATSStreaming("listen") }
	rootCmd.AddCommand(cmdFactory.NewCmdListen(queueCreator))
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

func getNATSStreaming(configPath string) stan.Conn {
	client, err := nats.NewStreaming(viper.GetString, configPath)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to NATS Streaming: %s", err))
	}

	return client
}
