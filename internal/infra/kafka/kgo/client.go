package kgo

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

type Client interface {
	ProduceSync(ctx context.Context, record Record) ProduceResult
	PollIter(ctx context.Context) (PollIter, error)
}

type client struct {
	Client *kgo.Client
}

func NewClient(seeds []string, topics ...string) (Client, error) {
	kgoClient, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumeTopics(topics...),
	)

	if err != nil {
		return &client{Client: kgoClient}, err
	}

	cl := &client{Client: kgoClient}

	return cl, nil
}

func (c *client) ProduceSync(ctx context.Context, record Record) ProduceResult {
	kgoRecord := kgo.Record{Topic: record.Topic(), Value: record.Value()}
	kgoResult := c.Client.ProduceSync(ctx, &kgoRecord)
	return NewProduceResult(kgoResult)
}

func (c *client) PollIter(ctx context.Context) (PollIter, error) {
	fetches := c.Client.PollFetches(ctx)
	if fetches.Errors() != nil {
		return nil, fetches.Err()
	}

	return NewPollIter(fetches.RecordIter()), nil
}
