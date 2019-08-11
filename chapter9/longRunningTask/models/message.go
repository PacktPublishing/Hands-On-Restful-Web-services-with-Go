package models

import (
	"time"

	"github.com/google/uuid"
)

// Job represents UUID of a Job
type Job struct {
	UUID uuid.UUID `json:"uuid"`
	Type string    `json:"type"`
}

// Log logs a given server time and is for Worker-A
type Log struct {
	ClientTime time.Time `json:"server_time"`
	Job
}

// CallBack is for worker-B
type CallBack struct {
	CallBackURL string `json:"callback_url"`
	Job
}

// Mail is for worker-C
type Mail struct {
	EmailAddress string `json:"email_address"`
	Job
}

// JobType is a struct that holds name and message of a worker
type JobType struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
