package app

import (
	"fmt"
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	"github.com/dairlair/sentimentd/pkg/infrastructure"
	"github.com/dairlair/sentimentd/pkg/infrastructure/repository"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/url"
)

type Config struct {
	Database struct {
		URL string
	}
}

type App struct {
	db *gorm.DB
	config *Config
	brainRepository BrainRepositoryInterface
}

func NewApp(config Config) *App {
	fmt.Printf("Config for app: %v\n", config)
	app := &App{
		config: &config,
	}
	app.Init()
	return app
}

func (app *App) Init() {
	databaseURL, err := url.Parse(app.config.Database.URL)
	if err != nil {
		log.Fatalf("Can not parse database URL. %s", err)
	}
	app.db = infrastructure.GetDB(databaseURL)
	app.brainRepository = repository.NewBrainRepository(app.db)
}

func (app *App) Destroy() {
	err := app.db.Close()
	if err != nil {
		log.Errorf("Database connection closing failed. %s", err)
	}
}