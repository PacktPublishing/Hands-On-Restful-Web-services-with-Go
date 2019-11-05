package main

import (
	"log"
	"time"

	"encoding/json"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTaskV1/models"
	"github.com/streadway/amqp"
)

// Workers does the job. It holds a connection
type Workers struct {
	conn *amqp.Connection
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
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

			job := models.Job{}
			err = json.Unmarshal(message.Body, &job)

			log.Printf("Workers received a message from the queue: %s", job)
			handleError(err, "Unable to load queue message")

			switch job.Type {
			case "A":
				w.dbWork(job)
			case "B":
				w.callbackWork(job)
			case "C":
				w.emailWork(job)
			}
		}
	}()

	defer w.conn.Close()
	wait := make(chan bool)
	<-wait // Run long-running worker
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string]interface{})
	log.Printf("Worker %s: extracting data..., JOB: %s", job.Type, result)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: saving data to database..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) callbackWork(job models.Job) {
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to the given callback..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) emailWork(job models.Job) {
	log.Printf("Worker %s: sending the email..., JOB: %s", job.Type, job.ID)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
}
