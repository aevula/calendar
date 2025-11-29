package application

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/dispatcher"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/registry"
	eventsRegistry "github.com/aevula/interview-hustlers-calendar/internal/jobs/registry/events"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/queues"
	"github.com/aevula/interview-hustlers-calendar/internal/scheduler"
	tasksService "github.com/aevula/interview-hustlers-calendar/internal/services/tasks"
	"github.com/aevula/interview-hustlers-calendar/internal/workers"
)

type Backjobs interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

type backjobs struct {
	app BaseApp

	logger logging.Logger

	queues     []queues.Queue
	dispatcher dispatcher.JobsDispatcher
	scheduler  scheduler.Scheduler
	pools      []workers.Pool
}

func NewBackjobs(ctx context.Context, cfg config.Config) Backjobs {
	bkg := &backjobs{app: New(cfg)}
	bkg.app.Init(ctx)

	bkg.logger = bkg.app.Logger().With(bkg.app.Logger().String("tag", "backjobs"))
	bkg.logger.Info(
		"Starting ...",
		bkg.logger.Int("pid", os.Getpid()),
		bkg.logger.String("env", bkg.app.Cfg().Env),
		bkg.logger.String("log_level", bkg.logger.Level().String()),
	)

	bkg.initQueues()
	bkg.initJobsDispatcher()
	bkg.initScheduler()
	bkg.initPools()

	return bkg
}

func (bkg *backjobs) Start(ctx context.Context) {
	defer bkg.logger.Sync()
	defer bkg.logger.Info("Started")

	for i := range bkg.queues {
		bkg.queues[i].Start()
	}

	bkg.scheduler.Start()

	for i := range bkg.pools {
		bkg.pools[i].Start()
	}
}

func (bkg *backjobs) Stop(ctx context.Context) {
	bkg.logger.Info("Shutting down ...")
	defer bkg.logger.Info("Shut down")

	var wg sync.WaitGroup
	defer wg.Wait()

	for i := range bkg.queues {
		wg.Go(func() { bkg.queues[i].Stop(ctx) })
	}

	wg.Go(func() { bkg.scheduler.Stop(ctx) })

	for i := range bkg.pools {
		wg.Go(func() { bkg.pools[i].Stop(ctx) })
	}
}

func (bkg *backjobs) initQueues() {
	bkg.queues = queues.Queues(bkg.app.Cfg())
}

func (bkg *backjobs) initJobsDispatcher() {
	producer, err := kafka.NewProducer(bkg.app.Cfg())
	if err != nil {
		bkg.logger.Fatal(err.Error())
	}

	reg := registry.NewJobFactoryRegistry()
	reg.Register("delete_old_events", eventsRegistry.NewDeleteOldEvents(bkg.app.EventsRepo()))
	reg.Register("send_events_notifications", eventsRegistry.NewSendEventsNotifications(bkg.app.EventsRepo(), producer))

	opts := dispatcher.DispatcherOptions{RefreshRate: 1 * time.Minute} // cfg

	bkg.dispatcher = dispatcher.NewJobsDispatcher(bkg.queues, reg, bkg.app.TasksRepo(), opts)
}

func (bkg *backjobs) initScheduler() {
	tasksService := tasksService.NewTasksService(bkg.app.TasksRepo())
	tasks := scheduler.PeriodicTasks(bkg.app.Cfg())
	bkg.scheduler = scheduler.New(tasksService, tasks)
}

func (bkg *backjobs) initPools() {
	bkg.pools = make([]workers.Pool, 0, len(bkg.queues))

	for i := range bkg.queues {
		opts := workers.WorkerOptions{}
		bkg.pools = append(bkg.pools, workers.NewPool(1, bkg.queues[i], opts))
	}
}
