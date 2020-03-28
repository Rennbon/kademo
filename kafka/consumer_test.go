package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestConsmer(t *testing.T) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"www.rennbon.online:19092", "www.rennbon.online:19093", "www.rennbon.online:19094"},
		GroupID:  "t1",
		Topic:    "test",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	for i := 0; i < 20; i++ {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
	r.Close()
	select {}
}

func TestProducer(t *testing.T) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"www.rennbon.online:19092", "www.rennbon.online:19093", "www.rennbon.online:19094"},
		Topic:    "test",
		Balancer: &kafka.Hash{},
		Async:    true,
	})
	for i := 200; i < 2000; i++ {
		err := w.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(fmt.Sprintf("key%v", i)),
			Value: []byte(fmt.Sprintf("val%v", i)),
		})
		fmt.Println(err, i)
		if err != nil {
			t.Error(err)
			return
		}
	}
	w.Close()
}
