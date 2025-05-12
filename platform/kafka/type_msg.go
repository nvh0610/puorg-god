package kafka_manager

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type MessageKeyValueStr struct {
	Key   string
	Value string
}

type MessageKeyValueByte struct {
	Key   []byte
	Value []byte
}

type ConsumerMessageByte struct {
	Partition int
	Offset    int64
	Key       []byte
	Value     []byte
	KeySer    interface{}
	ValueMap  map[string]interface{}
	ValueSer  interface{}
	kafkaMsg  *kafka.Message
	ctx       *context.Context
}
