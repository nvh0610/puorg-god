package kafka_manager

import (
	"context"
	"crawler_go_facebook/pkg/configs"
	"crawler_go_facebook/pkg/logger"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type KafkaProducer struct {
	*kafka.Writer
	isReady bool
}

func NewKafkaProducer() (*KafkaProducer, error) {
	var mechanism sasl.Mechanism
	var err error

	opts := &configs.KafkaConfig{}
	opts.LoadProducerEnvs()

	logger.InfoF("Connect to kafka writer [%v] ...", opts.Addr)

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

	return &KafkaProducer{
		Writer: &kafka.Writer{
			Addr:                   kafka.TCP(opts.Addr...),
			Balancer:               &kafka.LeastBytes{},
			BatchSize:              opts.BatchSize,
			BatchTimeout:           time.Millisecond * 5,
			AllowAutoTopicCreation: true,
			Transport:              &kafka.Transport{SASL: mechanism},
			RequiredAcks:           kafka.RequiredAcks(opts.RequiredAck),
			Async:                  opts.Async,
			ErrorLogger:            kafka.LoggerFunc(KafkaErrorF),
		},
		isReady: opts.IsReady,
	}, err
}

func (p *KafkaProducer) sendMsgs(topic string, data []MessageKeyValueStr) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error
	var msgs []kafka.Message

	for _, msg := range data {
		msgs = append(msgs, kafka.Message{
			Topic: topic,
			Key:   []byte(msg.Key),
			Value: []byte(msg.Value),
		})
	}

	if err = p.WriteMessages(ctx, msgs...); err != nil {
		return err
	}

	return nil
}

func (p *KafkaProducer) sendByteMessages(topic string, data ...MessageKeyValueByte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var err error
	var messages []kafka.Message

	for _, msg := range data {
		messages = append(messages, kafka.Message{
			Topic: topic,
			Key:   msg.Key,
			Value: msg.Value,
		})
	}

	if err = p.WriteMessages(ctx, messages...); err != nil {
		return err
	}
	return nil
}

func (p *KafkaProducer) WriteBatchMessages(topic string, data []MessageKeyValueByte) error {
	err := p.sendByteMessages(topic, data...)
	return err
}

func (p *KafkaProducer) WriteMessageNoKey(topic string, values []byte) error {
	data := MessageKeyValueByte{
		Value: values,
		Key:   nil,
	}
	err := p.sendByteMessages(topic, data)
	return err
}

func (p *KafkaProducer) WriteMessage(topic string, key []byte, values []byte) error {
	data := MessageKeyValueByte{
		Value: values,
		Key:   key,
	}
	err := p.sendByteMessages(topic, data)
	return err
}

func (p *KafkaProducer) IsReady() bool {
	return p.isReady
}
