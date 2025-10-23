package eventproducer

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/services/db-models"
	"github.com/services/utils/common"
)

type kafkaconsumer struct {
	consumer *kafka.Consumer
	db       *db.Queries
	topic    string
}

func NewKafkaConsumer(broker, groupID, topic string, db *db.Queries) (*kafkaconsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
	})
	if err != nil {
		log.Println("Kafka consumer error: ", err)
		return nil, err
	}
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Println("failed to subscribe producer topic :", err)
		return nil, err
	}
	return &kafkaconsumer{
		consumer: c,
		topic:    topic,
		db:       db,
	}, nil
}

func (kc *kafkaconsumer) Start(ctx context.Context) {
	log.Println("Kafka consumer started for topic: ", kc.topic)
	for {
		select {
		case <-ctx.Done():
			log.Println("Kafka consumer shutting down")
			kc.consumer.Close()
			return
		default:
			msg, err := kc.consumer.ReadMessage(time.Second)
			if err != nil {
				continue
			}
			kc.handleUserDelete(msg)
		}
	}
}
func (kc *kafkaconsumer) handleUserDelete(msg *kafka.Message) {
	var event common.KafkaEvent
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		log.Println("Failed to unmarshal event:", err)
		return
	}
	var user common.KafkaDeleteEvent
	if err := json.Unmarshal(event.Payload, &user); err != nil {
		log.Println("Failed to unmarshal payload:", err)
		return
	}
	if err := kc.db.DeleteUser(context.Background(), user.UserId); err != nil {
		log.Println("Failed to delete user in Auth DB:", err)
		return
	}
	log.Println("✅ User deleted in Auth DB:", user.UserId)
}
