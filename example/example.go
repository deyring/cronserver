package main

import (
	"fmt"

	"github.com/deyring/cronserver"
)

func main() {
	server, err := cronserver.CreateNewCronServer("0.1", "09.08.2016")
	if err != nil {
		panic(err)
	}
	// Do jobs without params
	server.Scheduler.Every(1, 0).Second().Do(task)
	server.Scheduler.Every(2, 1).Seconds().Do(task)
	server.Scheduler.Every(1, 2).Minute().Do(task)
	server.Scheduler.Every(2, 3).Minutes().Do(task)
	server.Scheduler.Every(1, 4).Hour().Do(task)
	server.Scheduler.Every(2, 5).Hours().Do(task)
	server.Scheduler.Every(1, 6).Day().Do(task)
	server.Scheduler.Every(2, 7).Days().Do(task)

	// Do jobs on specific weekday
	server.Scheduler.Every(1, 8).Monday().Do(task)
	server.Scheduler.Every(1, 9).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	server.Scheduler.Every(1, 10).Day().At("10:30").Do(task)
	server.Scheduler.Every(1, 11).Monday().At("18:30").Do(longtask)

	<-server.Scheduler.Start()
}

func task() {
	fmt.Println("Task done")
}

func longtask() {
	fmt.Println("ManualTask started")
}
