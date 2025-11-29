package kgo

import (
	"github.com/twmb/franz-go/pkg/kgo"
)

type ProduceResult interface {
	Error() error
}

type produceResult struct {
	err error
}

func NewProduceResult(raw kgo.ProduceResults) ProduceResult {
	return &produceResult{err: raw.FirstErr()}
}

func (r *produceResult) Error() error {
	return r.err
}
