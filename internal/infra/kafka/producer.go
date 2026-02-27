package kafka

import (
	"context"
	"encoding/json"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
	adapter "github.com/aevula/interview-hustlers-calendar/internal/infra/kafka/kgo"
)

type Producer interface {
	ProduceSync(ctx context.Context, topic string, data RecordValue) error
}

type producer struct {
	client adapter.Client
}

func NewProducer(cfg config.Config) (Producer, error) {
	cl, err := adapter.NewClient(cfg.Consumer.Seeds)
	return &producer{client: cl}, err
}

func (c *producer) ProduceSync(ctx context.Context, topic string, data RecordValue) error {
	record, err := toAdapterRecord(topic, data)
	if err != nil {
		return err
	}

	return c.client.ProduceSync(ctx, record).Error()
}

func toAdapterRecord(topic string, data RecordValue) (adapter.Record, error) {
	var zero adapter.Record

	value, err := json.Marshal(data)
	if err != nil {
		return zero, err
	}

	return adapter.NewRecord(topic, value), nil
}
