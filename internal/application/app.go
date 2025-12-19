package application

import (
	"context"
	"fmt"
	"net/http"
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
	Server *http.Server
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
	app.initServer(initCtx)

	app.logInited()
}

func (app *App) Run(ctx context.Context) <-chan struct{} {
	app.logStarting()
	errCh := make(chan struct{})

	go func() {
		defer close(errCh)

		app.logStarted()
		app.Logger.Sync()

		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Error(err.Error())
		}
	}()

	return errCh
}

func (app *App) Shutdown(ctx context.Context) {
	defer app.Logger.Sync()

	shutCtx, cancel := context.WithTimeout(ctx, app.Cfg.ShutTimeout)
	defer cancel()

	app.logStoping()
	defer app.logStoped()

	app.closeDb(shutCtx)
	app.closeServer(shutCtx)
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

func (app *App) initServer(ctx context.Context) {
	router := newRouter()

	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  app.Cfg.Server.ReadTimeout,
		WriteTimeout: app.Cfg.Server.WriteTimeout,
		IdleTimeout:  app.Cfg.Server.IdleTimeout,
		// ErrorLog:     app.Logger,
	}
}

func (app *App) closeDb(ctx context.Context) {
	err := app.Db.Close(ctx)
	if err != nil {
		app.Logger.Error(err.Error())
	}
}

func (app *App) closeServer(ctx context.Context) {
	err := app.Server.Shutdown(ctx)
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
	app.Logger.Info("Started", app.Logger.String("addr", app.Server.Addr))
}

func (app *App) logStoping() {
	app.Logger.Info("Shutting down ...")
}

func (app *App) logStoped() {
	app.Logger.Info("Shut down")
}
