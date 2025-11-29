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
	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
	usersService "github.com/aevula/interview-hustlers-calendar/internal/services/users"
)

type Server interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

type server struct {
	app BaseApp

	logger logging.Logger
	server *http.Server

	usersController  users.UsersController
	eventsController events.EventsController
}

func NewServer(ctx context.Context, cfg config.Config) Server {
	server := &server{app: New(cfg)}
	server.app.Init(ctx)

	server.logger = server.app.Logger().With(server.app.Logger().String("tag", "server"))
	server.logger.Info(
		"Starting ...",
		server.logger.Int("pid", os.Getpid()),
		server.logger.String("env", server.app.Cfg().Env),
		server.logger.String("log_level", server.logger.Level().String()),
	)

	server.initControllers()
	server.initServer()

	return server
}

func (server *server) Start(ctx context.Context) {
	go func() {
		server.logger.Info("Started", server.logger.String("addr", server.server.Addr))
		server.logger.Sync()

		if err := server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			server.logger.Error(err.Error())
		}
	}()
}

func (server *server) Stop(ctx context.Context) {
	server.logger.Info("Shutting down ...")
	defer server.logger.Info("Shut down")

	server.closeServer(ctx)

	server.app.Stop(ctx)
}

func (server *server) initControllers() {
	tasksService := tasksService.NewTasksService(server.app.TasksRepo())
	usersService := usersService.NewUsersService(server.app.UsersRepo())
	eventsService := eventsService.NewEventsService(server.app.EventsRepo(), tasksService)

	server.usersController = users.NewUsersController(usersService, server.logger)
	server.eventsController = events.NewEventsController(eventsService, server.logger)
}

func (server *server) initServer() {
	router := handlers.NewRouter()
	router.BuildUsers(server.usersController)
	router.BuildEvents(server.eventsController)

	server.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", server.app.Cfg().Server.Port),
		Handler:      router.Done(),
		ReadTimeout:  server.app.Cfg().Server.ReadTimeout,
		WriteTimeout: server.app.Cfg().Server.WriteTimeout,
		IdleTimeout:  server.app.Cfg().Server.IdleTimeout,
		// ErrorLog:     server.app.Logger,
	}
}

func (server *server) closeServer(ctx context.Context) {
	err := server.server.Shutdown(ctx)
	if err != nil {
		server.logger.Error(err.Error())
	}
}
