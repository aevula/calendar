package kgo

import "github.com/twmb/franz-go/pkg/kgo"

type PollIter interface {
	Done() bool
	Next() Record
}

type pollIter struct {
	iter *kgo.FetchesRecordIter
}

func NewPollIter(raw *kgo.FetchesRecordIter) PollIter {
	return &pollIter{iter: raw}
}

func (pi *pollIter) Done() bool {
	return pi.iter.Done()
}

func (pi *pollIter) Next() Record {
	rec := pi.iter.Next()
	return NewRecord(rec.Topic, rec.Value)
}
