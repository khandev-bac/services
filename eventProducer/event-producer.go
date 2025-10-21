package eventproducer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(brokers, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Println("kafka error: ", err)
		return nil, err
	}
	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}
func (kp *KafkaProducer) SendEvent(ctx context.Context, eventType string, payload interface{}) error {
	event := map[string]interface{}{
		"type":    eventType,
		"payload": payload,
	}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Parsing error in kafka: ", err)
		return err
	}
	err = kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kp.topic,
			Partition: kafka.PartitionAny,
		},
		Value: eventBytes,
	}, nil)
	if err != nil {
		log.Println("Producer error: ", err)
		return err
	}
	e := <-kp.producer.Events()
	m, ok := e.(*kafka.Message)
	if ok && m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}
	log.Printf("Event %s sent to Kafka topic %s\n", eventType, kp.topic)
	return nil
}

func (kp *KafkaProducer) Close() {
	kp.producer.Close()
}
