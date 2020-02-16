package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter9/longRunningTaskV1/models"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

// JobServer holds handler functions
type JobServer struct {
	Queue   amqp.Queue
	Channel *amqp.Channel
	Conn    *amqp.Connection
}

func (s *JobServer) asyncDBHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()
	queryParams := r.URL.Query()

	// Ex: client_time: 1569174071
	unixTime, err := strconv.ParseInt(queryParams.Get("client_time"), 10, 64)
	clientTime := time.Unix(unixTime, 0)
	handleError(err, "Error while converting client time")

	jsonBody, err := json.Marshal(models.Job{ID: jobID,
		Type:      "A",
		ExtraData: models.Log{ClientTime: clientTime},
	})
	handleError(err, "JSON body creation failed")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncCallbackHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()

	jsonBody, err := json.Marshal(models.Job{ID: jobID,
		Type:      "B",
		ExtraData: "",
	})
	handleError(err, "JSON body creation failed")

	if s.publish(jsonBody) == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) asyncMailHandler(w http.ResponseWriter, r *http.Request) {
	jobID, err := uuid.NewRandom()

	jsonBody, err := json.Marshal(models.Job{ID: jobID,
		Type:      "C",
		ExtraData: "",
	})
	handleError(err, "JSON body creation failed")

	err = s.publish(jsonBody)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *JobServer) publish(jsonBody []byte) error {
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonBody,
	}
	err := s.Channel.Publish(
		"",        // exchange
		queueName, // routing key(Queue)
		false,     // mandatory
		false,     // immediate
		message,
	)

	handleError(err, "Error while generating JobID")
	return err
}
