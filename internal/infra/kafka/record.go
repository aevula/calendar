package kafka

type Topic = string
type RecordValue = map[string]any

type Record struct {
	Topic Topic
	Value RecordValue
}
