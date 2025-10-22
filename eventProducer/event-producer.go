package eventproducer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

func NewKafkaProducer(brokers, topic string) (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"acks":              "all", // Wait for all replicas
		"retries":           3,
	})
	if err != nil {
		log.Println("kafka error:", err)
		return nil, err
	}
	return &KafkaProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (kp *KafkaProducer) SendEvent(ctx context.Context, eventType string, payload interface{}) error {
	// Match the field names your consumer expects
	event := map[string]interface{}{
		"event_type": eventType,        // Changed from "type" to "event_type"
		"timestamp":  time.Now().UTC(), // Added timestamp
		"payload":    payload,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Parsing error in kafka:", err)
		return err
	}

	log.Printf("📤 Sending event to Kafka - Type: %s, Payload: %s\n", eventType, string(eventBytes))

	// Use delivery channel to wait for confirmation
	deliveryChan := make(chan kafka.Event)
	err = kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kp.topic,
			Partition: kafka.PartitionAny,
		},
		Value: eventBytes,
	}, deliveryChan)

	if err != nil {
		log.Println("Producer error:", err)
		return err
	}

	// Wait for delivery report
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("❌ Delivery failed: %v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Printf("✅ Event '%s' delivered to topic %s [partition %d, offset %d]\n",
		eventType, kp.topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

	return nil
}

func (kp *KafkaProducer) Close() {
	log.Println("Flushing producer...")
	kp.producer.Flush(5000) // Wait up to 5 seconds for pending messages
	kp.producer.Close()
	log.Println("Producer closed")
}
