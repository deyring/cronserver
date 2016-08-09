package cronserver

import (
	"net/http"

	"github.com/deyring/gocron"
	"github.com/gin-gonic/gin"
)

var (
	serviceRoot = "/cron"
	s           *gocron.Scheduler
	health      Health
)

// CronServer ...
type CronServer struct {
	Scheduler *gocron.Scheduler
	Webserver *gin.Engine
}

// CreateNewCronServer ...
func CreateNewCronServer(version, date string) (*CronServer, error) {
	health = Health{ServiceName: "cron", BuildDate: date, BuildVersion: version, Status: "OK"}

	s = gocron.NewScheduler()
	server := CronServer{Scheduler: s, Webserver: gin.Default()}

	// Always expose health end point
	server.Webserver.GET(serviceRoot+"/status", handleStatus)
	server.Webserver.GET(serviceRoot+"/health", handleHealth)

	server.Webserver.POST(serviceRoot+"/startJob", handleStartJob)

	// Listen and Server in 0.0.0.0:80
	go func() {
		server.Webserver.Run(":80")
	}()

	return &server, nil
}

func handleStatus(c *gin.Context) {
	var response []JobStatus
	for _, job := range s.Jobs {
		if job != nil {
			status := JobStatus{ID: job.JobID,
				Interval: job.Interval,
				Unit:     job.Unit,
				JobFunc:  job.JobFunc,
				AtTime:   job.AtTime,
				LastRun:  job.LastRun,
				NextRun:  job.NextRun,
				Period:   job.Period,
			}
			response = append(response, status)
		}
	}
	c.IndentedJSON(http.StatusOK, response)
}

func handleStartJob(c *gin.Context) {
	var request StartJobRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	for _, job := range s.Jobs {
		if job != nil && job.JobID == request.JobID {
			_, err = job.Run()
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func handleHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, health)
}
