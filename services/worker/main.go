package main

import (
    "log"
    "os"
    "time"
    "github.com/streadway/amqp" 
)

func main() {
	rabbitHost := os.Getenv("RABBITMQ_HOST")
	if rabbitHost == "" {
		rabbitHost = "rabbitmq"
	}

	var conn *amqp.Connection
	var err error

	// Wait for RabbitMQ (10 attempts, 5 seconds apart)
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial("amqp://guest:guest@" + rabbitHost + ":5672/")
		if err == nil {
			break
		}
		log.Printf("Connecting to RabbitMQ... attempt %d", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Fatal: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Channel fail: %v", err)
	}
	defer ch.Close()

	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	log.Printf(" [*] Ready to process messages.")
	
	for d := range msgs {
		log.Printf(" [x] RECEIVED: %s", string(d.Body))
		time.Sleep(2 * time.Second)
		log.Printf(" [v] DONE")
	}
}