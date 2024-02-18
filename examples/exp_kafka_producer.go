package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	ctx := context.Background()

	endpoints := []string{
		//"localhost:9092",
		"localhost:9093",
		//"localhost:9094",
	}

	// Creation of a missing topic before publishing a message.
	// Note! it was the default behaviour up to the version v0.4.30.

	w := &kafka.Writer{
		Addr: kafka.TCP(endpoints...),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true, // v0.4.30以上版本支持此参数
	}

	err := w.WriteMessages(
		ctx,
		// NOTE: Each Message has Topic defined, otherwise an error is returned.
		kafka.Message{
			Topic: "topic-A",
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Topic: "topic-B",
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Topic: "topic-C",
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err = w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
