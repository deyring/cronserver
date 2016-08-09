package cronserver

import "time"

type JobStatus struct {
	ID       int
	Interval uint64
	Unit     string
	JobFunc  string
	AtTime   string
	LastRun  time.Time
	NextRun  time.Time
	Period   time.Duration
}

type StartJobRequest struct {
	JobID int `json:"job_id"`
}

type Health struct {
	BuildVersion string
	BuildDate    string
	ServiceName  string
	Status       string
}
