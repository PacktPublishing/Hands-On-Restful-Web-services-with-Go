package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTask/models"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// Server holds handler functions
type Server struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

func (s *Server) asyncDBHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()
	// Ex: 2012-11-01T22:08:41+00:00
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	clientTime := time.Unix(unixTime, 0)
	handleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Log{Job: models.Job{UUID: jobID}, ClientTime: clientTime})
	handleError(err, "JSON body creation failed")

	message := amqp.Publishing{
		ContentType: "text/json",
		Body:        jsonBody,
	}
	err = s.Channel.Publish(
		"",        // exchange
		queueName, // routing key(Queue)
		false,     // mandatory
		false,     // immediate
		message,
	)

	handleError(err, "Error while generating JobID")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Db Job: %s", jobID)
}

func (s *Server) asyncMailHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	handleError(err, "Error while generating JobID")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mail Job: %s", jobID)
}

func (s *Server) asyncCallbackHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	handleError(err, "Error while generating JobID")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Mail Callback: %s", jobID)
}
