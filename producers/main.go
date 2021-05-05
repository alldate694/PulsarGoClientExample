package main

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})

	_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte("hello"),
	})

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")

}
