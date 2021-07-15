package infra

import (
	"context"
	"fmt"
	"log"
	"mail-callbacks/config"
	"mail-callbacks/domain"
	"strings"

	"github.com/segmentio/kafka-go"
)

type produceMail struct {
	topic string
}

func NewKafkaProducer(topic string) *produceMail {
	return &produceMail{
		topic: topic,
	}
}

func (s *produceMail) ProduceMessage(m *domain.Message) error {
	log.Print("Message received! Id: ", m.MessageId())

	b, err := m.ToBytes()
	if err != nil {
		return fmt.Errorf("cannot convert struct to bytes: %v", err)
	}
	kafkaBrokers := strings.Split(config.GetConfig().KafkaBrokers(), ",")

	writer := &kafka.Writer{
		Addr:  kafka.TCP(kafkaBrokers...),
		Topic: s.topic,
	}

	pErr := writer.WriteMessages(context.Background(), kafka.Message{
		Value: b,
	})
	if pErr != nil {
		return fmt.Errorf("cannot write a message: %v", pErr)
	}

	return nil
}
