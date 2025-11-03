package application

import (
	"context"
	"os"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/aevula/interview-hustlers-calendar/internal/databases/postgres"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
)

type App struct {
	Cfg    config.Config
	Logger *logging.Logger
	Db     databases.Db
}

func New(cfg config.Config) *App {
	return &App{Cfg: cfg}
}

func (app *App) Init(ctx context.Context) {
	initCtx, cancel := context.WithTimeout(ctx, app.Cfg.InitTimeout)
	defer cancel()

	app.initLogger(initCtx)
	app.logIniting()

	app.initDb(initCtx)

	app.logInited()
}

func (app *App) Run(ctx context.Context) {
	app.logStarting()
	defer app.logStarted()

	for {
	}
}

func (app *App) Shutdown(ctx context.Context) {
	defer app.Logger.Sync()

	app.logStoping()
	defer app.logStoped()

	app.closeDb(ctx)
}

func (app *App) initLogger(ctx context.Context) {
	app.Logger = logging.MustLoad(app.Cfg)
}

func (app *App) initDb(ctx context.Context) {
	pg, err := postgres.New(ctx, app.Cfg)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	app.Db = databases.Db(pg)
}

func (app *App) closeDb(ctx context.Context) {
	err := app.Db.Close(ctx)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}

func (app *App) logIniting() {
	app.Logger.Info(
		"Starting ...",
		app.Logger.Int("pid", os.Getpid()),
		app.Logger.String("env", app.Cfg.Env),
		app.Logger.String("log_level", app.Logger.Level().String()),
	)
}

func (app *App) logInited() {}

func (app *App) logStarting() {}

func (app *App) logStarted() {
	app.Logger.Info("Started")
}

func (app *App) logStoping() {
	app.Logger.Info("Shutting down ...")
}

func (app *App) logStoped() {
	app.Logger.Info("Shut down")
}
