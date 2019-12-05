package app

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/infrastructure/repository"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type Config struct {
	Database struct {
		URL string
	}
}

type App struct {
	brainRepository BrainRepositoryInterface
}

func NewApp(config Config) *App {

	fmt.Printf("Config for app: %v\n", config)
	url, err := url.Parse(config.Database.URL)

	if err != nil {
		log.Fatalf("Can not parse database URL. %s", err)
	}

	fmt.Printf("%v\n", url)

	brainRepository := NewBrainRepository()

	return &App{
		brainRepository: brainRepository,
	}
}