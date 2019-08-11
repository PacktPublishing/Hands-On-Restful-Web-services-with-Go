package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTask/models"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

const queueName string = "jobQueue"
const hostString string = "127.0.0.1:8000"

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func startWorkers(connection *amqp.Connection) {
	/*
		Run workers for the given job types
	*/
	var jobTypes = [...]models.JobType{{
		Name: "A",
		Type: "database",
	}, {
		Name: "B",
		Type: "callback",
	}, {
		Name: "C",
		Type: "mail",
	},
	}

	for _, jobType := range jobTypes {
		go func(jobType models.JobType) {
			workerProcess := Worker{
				name: jobType.Name,
				conn: connection,
			}
			workerProcess.run()
		}(jobType)
	}
}

func init() {
	/*
		Initiate workers on launch
	*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Dialing failed to RabbitMQ broker")
	defer conn.Close()

	startWorkers(conn)
}

func getServer(name string) Server {
	/*
		Creates a server object and initiates
		the Channel and Queue details to publish messages
	*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Dialing failed to RabbitMQ broker")
	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Fetching channel failed")
	defer channel.Close()

	jobQueue, err := channel.QueueDeclare(
		"jobQueue", // Name of the queue
		false,      // Message is persisted or not
		false,      // Delete message when unused
		false,      // Exclusive
		false,      // No Waiting time
		nil,        // Extra args
	)
	handleError(err, "Job queue creation failed")
	return Server{Channel: channel, Queue: jobQueue}
}

func main() {
	router := mux.NewRouter()
	server := getServer(queueName)
	// Attach handlers
	router.HandleFunc("/job/database", server.asyncDBHandler)
	router.HandleFunc("/job/mail", server.asyncMailHandler)
	router.HandleFunc("/job/callback", server.asyncMailHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         hostString,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
