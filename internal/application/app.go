package application

import (
	"context"
	"os"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/databases/postgres"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
)

type App struct {
	Cfg    config.Config
	Logger *logging.Logger
	Db     *postgres.Db
}

func New(ctx context.Context) *App {
	app := new(App)
	app.initCfg()

	initCtx, cancel := context.WithTimeout(ctx, app.Cfg.InitTimeout)
	defer cancel()

	app.initLogger(initCtx)
	app.logStarting()

	app.initDb(initCtx)

	app.logStarted()

	return app
}

func (app *App) Run(ctx context.Context) {
	defer app.Logger.Sync()

	for {
	}
}

func (app *App) Shutdown(ctx context.Context, done chan<- struct{}) {
	defer close(done)
	defer app.Logger.Sync()

	app.logStoping()

	app.closeDb(ctx)

	app.logStoped()
}

func (app *App) initCfg() {
	app.Cfg = config.MustLoad()
}

func (app *App) initLogger(ctx context.Context) {
	app.Logger = logging.MustLoad(app.Cfg)
}

func (app *App) initDb(ctx context.Context) {
	db, err := postgres.New(ctx, app.Cfg)
	if err != nil {
		app.Logger.Fatal(err.Error())
	}

	app.Db = db
}

func (app *App) closeDb(ctx context.Context) {
	err := app.Db.Close(ctx)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}

func (app *App) logStarting() {
	app.Logger.Info(
		"Starting ...",
		app.Logger.Int("pid", os.Getpid()),
		app.Logger.String("env", app.Cfg.Env),
		app.Logger.String("log_level", app.Logger.Level().String()),
	)
}

func (app *App) logStarted() {
	app.Logger.Info("Started")
}

func (app *App) logStoping() {
	app.Logger.Info("Shutting down ...")
}

func (app *App) logStoped() {
	app.Logger.Info("Shut down")
}
