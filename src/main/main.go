package main

import (
	"fmt"
	"mypool"
	"time"
)

var (
	//MaxWorker os.Getenv("MAX_WORKERS")
	MaxWorker = 3
	//MaxQueue os.Getenv("MAX_QUEUE")
	MaxQueue = 10
)

func main() {
	dispatcher := mypool.NewDispatcher(MaxWorker)
	dispatcher.Run()

	go func() {
		for i := 0; i < MaxQueue; i++ {
			job := dispatcher.MakeJob(i)
			time.Sleep(1 * time.Second)
			fmt.Println("\twill send a job:", i)
			mypool.JobQueue <- job
			fmt.Println("\tend send a job.")
		}
	}()

	time.Sleep(120 * time.Second)
}
