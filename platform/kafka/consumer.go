package kafka_manager

import (
	"context"
	"crawler_go_facebook/jobs/schedule"
	"crawler_go_facebook/pkg/configs"
	"crawler_go_facebook/pkg/helper"
	"crawler_go_facebook/pkg/logger"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"sync/atomic"
	"time"

	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

const (
	maxAttempts = 10
)

type (
	ReadMessageFn  func(ctx context.Context) (kafka.Message, error)
	FetchMessageFn func(ctx context.Context) (kafka.Message, error)
)

type KafkaConsumer struct {
	reader       *kafka.Reader
	KeySer       Serializer
	ValueSer     Serializer
	speedRead    atomic.Int64
	entryId      int
	readMessage  ReadMessageFn
	fetchMessage FetchMessageFn
}

func KafkaErrorF(msg string, args ...interface{}) {
	logger.ErrorF(msg, args...)
}
func KafkaInfo(msg string, args ...interface{}) {
	logger.InfoF(msg, args...)
}

func newKafkaConsumer(opts *configs.KafkaConfig, topic, groupId string, ValueSer Serializer) (*KafkaConsumer, error) {
	if opts == nil {
		opts = &configs.KafkaConfig{}
		opts.LoadConsumerEnvs()
	}

	logger.InfoF("Connect to kafka [%v] ...", opts.Addr)

	var mechanism sasl.Mechanism
	var err error

	switch opts.AuthType {
	case "SHA256":
		mechanism, err = scram.Mechanism(scram.SHA256, opts.Username, opts.Password)
	case `SHA512`:
		mechanism, err = scram.Mechanism(scram.SHA512, opts.Username, opts.Password)
	case `PLAIN`:
		mechanism = plain.Mechanism{
			Username: opts.Username,
			Password: opts.Password,
		}
	}

	if err != nil {
		return nil, err
	}

	dialer := &kafka.Dialer{
		Timeout:       time.Second * 60,
		Deadline:      time.Time{},
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:               opts.Addr,
		GroupID:               groupId,
		Topic:                 topic,
		Dialer:                dialer,
		OffsetOutOfRangeError: true,
		WatchPartitionChanges: true,
		CommitInterval:        1 * time.Second, //set 0 if using sync reader
		StartOffset:           kafka.FirstOffset,
		MaxAttempts:           maxAttempts,
		ErrorLogger:           kafka.LoggerFunc(KafkaErrorF),
		//Logger:                kafka.LoggerFunc(KafkaInfo),
	})

	return &KafkaConsumer{
		reader:       r,
		KeySer:       &DefaultSerializer{},
		ValueSer:     ValueSer,
		readMessage:  r.ReadMessage,
		fetchMessage: r.FetchMessage,
	}, err
}

func (c *KafkaConsumer) SetReadMessageFn(fn ReadMessageFn) {
	c.readMessage = fn
}

func (c *KafkaConsumer) SetFetchMessage(fn FetchMessageFn) {
	c.fetchMessage = fn
}

func (c *KafkaConsumer) ReadMessage() (*ConsumerMessageByte, error) {
	var msg kafka.Message
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	if msg, err = c.fetchMessage(ctx); err != nil {
		return nil, err
	}

	c.speedRead.Add(1)

	return &ConsumerMessageByte{
		Partition: msg.Partition,
		Offset:    msg.Offset,
		Key:       msg.Key,
		Value:     msg.Value,
		KeySer:    c.KeySer.Decode(msg.Key),
		ValueMap:  helper.FromJsonByteToMap(msg.Value),
		ValueSer:  c.ValueSer.Decode(msg.Value),
		kafkaMsg:  &msg,
	}, err
}

func (c *KafkaConsumer) CommitMessage(msg *ConsumerMessageByte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	err := c.reader.CommitMessages(ctx, *msg.kafkaMsg)
	return err
}

func (c *KafkaConsumer) ReadMessageCommit() (*ConsumerMessageByte, error) {
	var msg *ConsumerMessageByte
	var err error

	if msg, err = c.ReadMessage(); err != nil {
		return nil, err
	}

	err = c.CommitMessage(msg)
	return msg, err
}

func (c *KafkaConsumer) ReadMsgBatch(batchSize int) ([]*ConsumerMessageByte, error) {
	var err error
	var msgBatch []kafka.Message
	var msgBatchSchema []*ConsumerMessageByte

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	for i := 0; i < batchSize; i++ {
		var msg kafka.Message
		if msg, err = c.fetchMessage(ctx); err != nil {
			break
		}

		msgSchema := &ConsumerMessageByte{
			Partition: msg.Partition,
			Offset:    msg.Offset,
			Key:       msg.Key,
			Value:     msg.Value,
			KeySer:    c.KeySer.Decode(msg.Key),
			ValueSer:  c.ValueSer.Decode(msg.Value),
		}

		msgBatch = append(msgBatch, msg)
		msgBatchSchema = append(msgBatchSchema, msgSchema)
	}

	c.speedRead.Add(1)

	if err != nil {
		switch {
		case !errors.Is(err, ctx.Err()):
			return nil, err
		}
	}

	if len(msgBatch) > 0 {
		select {
		case <-ctx.Done():
			ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
		default:
		}

		if err = c.reader.CommitMessages(ctx, msgBatch...); err != nil {
			return nil, err
		}
	}
	return msgBatchSchema, err
}

func (c *KafkaConsumer) Init() {
	entryId, _ := schedule.RegisterScheduler("0 */5 * * * *", func() {
		fmt.Printf("Kafka reader speed %d msg/m\n", c.speedRead.Add(0)/5)
		c.speedRead.Store(0)
	})
	c.entryId = int(entryId)
}
func (c *KafkaConsumer) Close() {
	logger.Info("Close camera-ai application kafka readers")
	if c.entryId > 0 {
		schedule.UnRegisterScheduler(c.entryId)
	}
	c.reader.Close()
}
