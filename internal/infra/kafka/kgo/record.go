package kgo

type Record interface {
	Topic() string
	Value() []byte
}

type record struct {
	topic string
	value []byte
}

func NewRecord(topic string, value []byte) Record {
	return &record{
		topic: topic,
		value: value,
	}
}

func (r *record) Topic() string {
	return r.topic
}

func (r *record) Value() []byte {
	return r.value
}
