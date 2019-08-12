package main

import (
	"log"

	"github.com/streadway/amqp"
)

// Worker does the job. It holds a connection
type Worker struct {
	name string
	conn *amqp.Connection
}

func (w *Worker) run() {
	log.Printf("Worker %s has booted up", w.name)
	channel, err := w.conn.Channel()
	handleError(err, "Fetching channel failed")
	defer channel.Close()

	jobQueue, err := channel.QueueDeclare(
		queueName, // Name of the queue
		false,     // Message is persisted or not
		false,     // Delete message when unused
		false,     // Exclusive
		false,     // No Waiting time
		nil,       // Extra args
	)
	handleError(err, "Job queue fetch failed")

	messages, err := channel.Consume(
		jobQueue.Name, // queue
		"",            // consumer
		true,          // auto-acknowledge
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	go func() {
		for message := range messages {
			log.Printf("Worker %s: Received a message from the queue: %s", w.name, message.Body)
		}
	}()

	defer w.conn.Close()
	wait := make(chan bool)
	<-wait
}
