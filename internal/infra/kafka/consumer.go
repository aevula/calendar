package kafka

import (
	"context"
	"encoding/json"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	adapter "github.com/aevula/interview-hustlers-calendar/internal/infra/kafka/kgo"
)

type ConsumedRecordHandler = func(ctx context.Context, record Record)

type PollIter interface {
	Done() bool
	Next() Record
}

type Consumer interface {
	PollAndHandle(ctx context.Context, handler ConsumedRecordHandler) error
}

type consumer struct {
	client adapter.Client
}

func NewConsumer(cfg config.Config) (Consumer, error) {
	cl, err := adapter.NewClient(cfg.Consumer.Seeds, cfg.Consumer.Topics...)
	return &consumer{client: cl}, err
}

func (c *consumer) PollAndHandle(ctx context.Context, handler ConsumedRecordHandler) error {
	iter, err := c.client.PollIter(ctx)
	if err != nil {
		return err
	}

	for iter.Done() {
		record, err := fromAdapterRecord(iter.Next())
		if err != nil {
			continue
		}

		handler(ctx, record)
	}

	return nil
}

func fromAdapterRecord(raw adapter.Record) (Record, error) {
	var zero Record

	value := make(RecordValue, 0)
	err := json.Unmarshal(raw.Value(), &value)
	if err != nil {
		return zero, nil
	}

	return Record{Topic: raw.Topic(), Value: value}, nil
}
