package main

import (
	"log"
	"time"
	"os"
	"github.com/streadway/amqp"
)

func main() {
	// Get RabbitMQ host from environment
	rabbitHost := os.Getenv("RABBITMQ_HOST")
	if rabbitHost == "" {
		rabbitHost = "localhost"
	}

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@" + rabbitHost + ":5672/")
	for err != nil {
		log.Println("Failed to connect to RabbitMQ, retrying in 5s...")
		time.Sleep(5 * time.Sleep)
		conn, err = amqp.Dial("amqp://guest:guest@" + rabbitHost + ":5672/")
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	// Declare the same queue as the publisher
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)

	// Consume messages
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received a message: %s", d.Body)
			// Simulate heavy work
			time.Sleep(2 * time.Second)
			log.Printf(" [v] Done processing task")
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}