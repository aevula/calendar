package application

import (
	"context"
	"os"
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
	"github.com/aevula/interview-hustlers-calendar/internal/workers"
)

type Scheduler interface {
	Init(ctx context.Context)
	Run(ctx context.Context) (done <-chan struct{})
	Stop(ctx context.Context)
}

type scheduler struct {
	app BaseApp

	logger  logging.Logger
	workers []workers.Worker
	ready   chan workers.Worker
}

func NewScheduler(cfg config.Config) Scheduler {
	return &scheduler{app: new(cfg)}
}

func (sched *scheduler) Init(ctx context.Context) {
	sched.app.Init(ctx)

	sched.logger = sched.app.Logger().With(sched.app.Logger().String("tag", "scheduler"))
	sched.logger.Info(
		"Starting ...",
		sched.logger.Int("pid", os.Getpid()),
		sched.logger.String("env", sched.app.Cfg().Env),
		sched.logger.String("log_level", sched.logger.Level().String()),
	)

	sched.initWorkers()
}

func (sched *scheduler) Run(ctx context.Context) <-chan struct{} {
	done := make(chan struct{})

	for i := range sched.workers {
		sched.workers[i].Run(ctx)
	}

	ticker := time.NewTicker(sched.app.Cfg().Scheduler.RefreshRate)

	go func() {
		defer close(done)
		defer ticker.Stop()

		sched.logger.Info("Started")
		sched.logger.Sync()

		for {
			select {
			case <-ticker.C:
				sched.tick(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()

	return done
}

func (sched *scheduler) Stop(ctx context.Context) {
	sched.logger.Info("Shutting down ...")
	defer sched.logger.Info("Shut down")

	sched.stopWorkers(ctx)

	sched.app.Stop(ctx)
}

func (sched *scheduler) initWorkers() {
	sched.ready = make(chan workers.Worker, sched.app.Cfg().Scheduler.WorkersCount)
	sched.workers = make([]workers.Worker, 0, sched.app.Cfg().Scheduler.WorkersCount)

	for range sched.app.Cfg().Scheduler.WorkersCount {
		sched.workers = append(sched.workers, workers.New(sched.ready))
	}
}

func (sched *scheduler) tick(ctx context.Context) {
	job := jobs.NewDeleteOldEvents(sched.app.EventsRepo(), sched.logger)

	select {
	case worker := <-sched.ready:
		worker.Add(job)
	case <-ctx.Done():
	default:
	}
}

func (sched *scheduler) stopWorkers(ctx context.Context) {
	for i := range sched.workers {
		if err := sched.workers[i].Stop(ctx); err != nil {
			sched.logger.Error(err.Error())
		}
	}
}
