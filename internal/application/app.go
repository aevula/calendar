package application

import (
	"context"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/aevula/interview-hustlers-calendar/internal/databases/postgres"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/repository"
)

type BaseApp interface {
	Init(ctx context.Context)
	Stop(ctx context.Context)

	Cfg() config.Config
	Logger() logging.Logger
	DB() databases.DB

	UsersRepo() repository.UserRepository
	EventsRepo() repository.EventRepository
}

type app struct {
	cfg    config.Config
	logger logging.Logger
	db     databases.DB

	usersRepo  repository.UserRepository
	eventsRepo repository.EventRepository
}

func new(cfg config.Config) BaseApp {
	return &app{cfg: cfg}
}

func (app *app) Init(ctx context.Context) {
	app.initLogger()

	app.initDb(ctx)

	app.initRepos()
}

func (app *app) Stop(ctx context.Context) {
	defer app.logger.Sync()

	app.closeDb(ctx)
}

func (app *app) Cfg() config.Config {
	return app.cfg
}

func (app *app) Logger() logging.Logger {
	return app.logger
}

func (app *app) DB() databases.DB {
	return app.db
}

func (app *app) UsersRepo() repository.UserRepository {
	return app.usersRepo
}

func (app *app) EventsRepo() repository.EventRepository {
	return app.eventsRepo
}

func (app *app) initLogger() {
	app.logger = logging.MustLoad(app.cfg)
}

func (app *app) initDb(ctx context.Context) {
	db, err := postgres.New(ctx, app.cfg)
	if err != nil {
		app.logger.Fatal(err.Error())
	}

	if err = db.Ping(ctx); err != nil {
		app.logger.Fatal(err.Error())
	}

	app.db = db
}

func (app *app) initRepos() {
	app.usersRepo = repository.NewUserRepository(app.db)
	app.eventsRepo = repository.NewEventRepository(app.db)
}

func (app *app) closeDb(ctx context.Context) {
	err := app.db.Close(ctx)
	if err != nil {
		app.logger.Error(err.Error())
	}
}
