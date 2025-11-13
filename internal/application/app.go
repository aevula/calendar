package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	controllers "github.com/aevula/interview-hustlers-calendar/internal/controllers/http"
	"github.com/aevula/interview-hustlers-calendar/internal/databases"
	"github.com/aevula/interview-hustlers-calendar/internal/databases/postgres"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
)

type App struct {
	cfg    config.Config
	logger *logging.Logger
	db     databases.Db
	Server *http.Server
}

func New(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Config() config.Config {
	return app.cfg
}

func (app *App) Logger() *logging.Logger {
	return app.logger
}

func (app *App) Db() databases.Db {
	return app.db
}

func (app *App) Init(ctx context.Context) {
	initCtx, cancel := context.WithTimeout(ctx, app.cfg.InitTimeout)
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
		app.logger.Sync()

		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Error(err.Error())
		}
	}()

	return errCh
}

func (app *App) Shutdown(ctx context.Context) {
	defer app.logger.Sync()

	shutCtx, cancel := context.WithTimeout(ctx, app.cfg.ShutTimeout)
	defer cancel()

	app.logStoping()
	defer app.logStoped()

	app.closeDb(shutCtx)
	app.closeServer(shutCtx)
}

func (app *App) initLogger(ctx context.Context) {
	app.logger = logging.MustLoad(app.cfg)
}

func (app *App) initDb(ctx context.Context) {
	db, err := postgres.New(ctx, app.cfg)
	if err != nil {
		app.logger.Fatal(err.Error())
	}

	app.db = db
}

func (app *App) initServer(ctx context.Context) {
	router := controllers.NewRouter(app)

	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  app.cfg.Server.ReadTimeout,
		WriteTimeout: app.cfg.Server.WriteTimeout,
		IdleTimeout:  app.cfg.Server.IdleTimeout,
		// ErrorLog:     app.logger,
	}
}

func (app *App) closeDb(ctx context.Context) {
	err := app.db.Close(ctx)
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *App) closeServer(ctx context.Context) {
	err := app.Server.Shutdown(ctx)
	if err != nil {
		app.logger.Error(err.Error())
	}
}

func (app *App) logIniting() {
	app.logger.Info(
		"Starting ...",
		app.logger.Int("pid", os.Getpid()),
		app.logger.String("env", app.cfg.Env),
		app.logger.String("log_level", app.logger.Level().String()),
	)
}

func (app *App) logInited() {}

func (app *App) logStarting() {}

func (app *App) logStarted() {
	app.logger.Info("Started", app.logger.String("addr", app.Server.Addr))
}

func (app *App) logStoping() {
	app.logger.Info("Shutting down ...")
}

func (app *App) logStoped() {
	app.logger.Info("Shut down")
}
