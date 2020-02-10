package application

import (
	. "github.com/dairlair/sentimentd/pkg/domain/repository"
	. "github.com/dairlair/sentimentd/pkg/domain/service"
	"github.com/dairlair/sentimentd/pkg/domain/service/predictor"
	"github.com/dairlair/sentimentd/pkg/domain/service/tokenizer"
	"github.com/dairlair/sentimentd/pkg/domain/service/training"
	"github.com/dairlair/sentimentd/pkg/infrastructure/db"
	"github.com/dairlair/sentimentd/pkg/infrastructure/repository"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/url"
	"time"
)

// Config contains a configuration which is required for Apps
type Config struct {
	Log struct {
		Level  string
		Format string
	}
	Database struct {
		URL               string
		ConnectionTimeout time.Duration
		Automigrate       bool
		MigrationsPath    string
	}
}

// App is used to contains all services together and control their life
type App struct {
	db              *gorm.DB
	config          *Config
	brainRepository BrainRepositoryInterface
	classRepository ClassRepositoryInterface
	trainingService *training.Service
	predictor       *predictor.Predictor
}

// NewApp creates new App
func NewApp(config Config) *App {
	app := &App{
		config: &config,
	}
	return app
}

// Init runs initialization for all application components.
// Requires database and other external services are available.
func (app *App) Init() {
	if app.config.Database.Automigrate {
		app.Migrate()
	}
	databaseURL, err := url.Parse(app.config.Database.URL)
	if err != nil {
		log.Fatalf("Can not parse database URL. %s", err)
	}
	app.db = db.CreateDBConnection(databaseURL, app.config.Database.ConnectionTimeout)
	app.brainRepository = repository.NewBrainRepository(app.db)
	app.classRepository = repository.NewClassRepository(app.db)
	tokenRepository := repository.NewTokenRepository(app.db)
	defaultTokenizer := tokenizer.NewTokenizer()
	tokenService := NewTokenService(tokenRepository)
	classService := NewClassService(app.classRepository)
	resultsRepository := repository.NewResultRepository(app.db)
	app.trainingService = training.NewTrainingService(
		classService,
		resultsRepository,
		&defaultTokenizer,
		tokenService,
	)
	app.predictor = predictor.NewPredictor(&defaultTokenizer, tokenRepository, resultsRepository)
}

// Destroy closes all connections
func (app *App) Destroy() {
	err := app.db.Close()
	if err != nil {
		log.Errorf("Database connection closing failed. %s", err)
	}
}
