package application

import (
	"context"
	"os"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	"github.com/aevula/interview-hustlers-calendar/internal/infra/kafka"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/consumer/registry"
	"github.com/aevula/interview-hustlers-calendar/internal/jobs/consumer/registry/events"
	"github.com/aevula/interview-hustlers-calendar/internal/logging"
)

type Consumer interface {
	Start(ctx context.Context)
	Stop(ctx context.Context)
}

type consumer struct {
	app BaseApp

	logger logging.Logger

	consumer kafka.Consumer
	handlers registry.JobFactoryRegistry
	stop     context.CancelFunc
}

func NewConsumer(ctx context.Context, cfg config.Config) Consumer {
	cns := &consumer{app: New(cfg)}
	cns.app.Init(ctx)

	cns.logger = cns.app.Logger().With(cns.app.Logger().String("tag", "scheduler"))
	cns.logger.Info(
		"Starting ...",
		cns.logger.Int("pid", os.Getpid()),
		cns.logger.String("env", cns.app.Cfg().Env),
		cns.logger.String("log_level", cns.logger.Level().String()),
	)

	cns.initKafka()
	cns.initHandlers()

	return cns
}

func (cns *consumer) Start(ctx context.Context) {
	defer cns.logger.Sync()
	defer cns.logger.Info("Started")

	go func() {
		ctx, stop := context.WithCancel(context.Background())
		cns.stop = stop

		for {
			select {
			case <-ctx.Done():
				return
			default:
				cns.consumer.PollAndHandle(ctx, cns.handleRecord)
			}
		}
	}()
}

func (cns *consumer) Stop(ctx context.Context) {
	cns.logger.Info("Shutting down ...")
	defer cns.logger.Info("Shut down")

	cns.stop()
}

func (cns *consumer) initKafka() {
	cons, err := kafka.NewConsumer(cns.app.Cfg())
	if err != nil {
		cns.logger.Fatal(err.Error())
	}

	cns.consumer = cons
}

func (cns *consumer) initHandlers() {
	cns.handlers = registry.NewJobFactoryRegistry()
	cns.handlers.Register("event_notifications", events.NewSendEventNotification(cns.app.EventsRepo()))
}

func (cns *consumer) handleRecord(ctx context.Context, record kafka.Record) {
	factory, ok := cns.handlers.Get(record.Topic)
	if !ok {
		cns.logger.Error("job not found", cns.logger.String("topic", record.Topic))
	}

	job, err := factory.Create(record)
	if err != nil {
		cns.logger.Error(err.Error(), cns.logger.String("topic", record.Topic))
	}

	job.Run(ctx)
}
