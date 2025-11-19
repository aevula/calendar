package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/events"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/handlers"
	"github.com/aevula/interview-hustlers-calendar/internal/controllers/http/users"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	eventsService "github.com/aevula/interview-hustlers-calendar/internal/services/events"
	usersService "github.com/aevula/interview-hustlers-calendar/internal/services/users"
)

type Server interface {
	Init(ctx context.Context)
	Run(ctx context.Context) (done <-chan struct{})
	Stop(ctx context.Context)
}

type server struct {
	app BaseApp

	logger           logging.Logger
	Server           *http.Server
	UsersController  users.UsersController
	EventsController events.EventsController
}

func NewServer(cfg config.Config) Server {
	return &server{app: new(cfg)}
}

func (server *server) Init(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, server.app.Cfg().Server.InitTimeout)
	defer cancel()

	server.app.Init(ctx)
	server.logger = server.app.Logger().With(server.app.Logger().String("tag", "server"))

	server.logStarting()

	server.initControllers()
	server.initServer()
}

func (server *server) Run(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	go func() {
		defer close(done)

		server.logStarted()
		server.logger.Sync()

		if err := server.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.logger.Error(err.Error())
		}
	}()

	return done
}

func (server *server) Stop(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, server.app.Cfg().Server.ShutTimeout)
	defer cancel()

	server.logStoping()
	defer server.logStoped()

	server.closeServer(ctx)

	server.app.Stop(ctx)
}

func (server *server) initControllers() {
	usersService := usersService.NewUsersService(server.app.UsersRepo())
	eventsService := eventsService.NewEventsService(server.app.EventsRepo())

	server.UsersController = users.NewUsersController(usersService, server.logger)
	server.EventsController = events.NewEventsController(eventsService, server.logger)
}

func (server *server) initServer() {
	router := handlers.NewRouter()
	router.BuildUsers(server.UsersController)
	router.BuildEvents(server.EventsController)

	server.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", server.app.Cfg().Server.Port),
		Handler:      router.Done(),
		ReadTimeout:  server.app.Cfg().Server.ReadTimeout,
		WriteTimeout: server.app.Cfg().Server.WriteTimeout,
		IdleTimeout:  server.app.Cfg().Server.IdleTimeout,
		// ErrorLog:     server.app.Logger,
	}
}

func (server *server) closeServer(ctx context.Context) {
	err := server.Server.Shutdown(ctx)
	if err != nil {
		server.logger.Error(err.Error())
	}
}

func (server *server) logStarting() {
	server.logger.Info(
		"Starting ...",
		server.logger.Int("pid", os.Getpid()),
		server.logger.String("env", server.app.Cfg().Env),
		server.logger.String("log_level", server.logger.Level().String()),
	)
}

func (server *server) logStarted() {
	server.logger.Info("Started", server.logger.String("addr", server.Server.Addr))
}

func (server *server) logStoping() {
	server.logger.Info("Shutting down ...")
}

func (server *server) logStoped() {
	server.logger.Info("Shut down")
}
