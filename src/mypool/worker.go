package mypool

import (
	"fmt"
	"math/rand"
	"time"
)

// Job represents the job to be run
type Job struct {
	Payload int
}

//JobQueue A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

//NewWorker ccc
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			fmt.Println("Worker rotine: idle. put job channel in workerpool.")
			w.WorkerPool <- w.JobChannel
			fmt.Println("Worker rotine: success: put job channel in workerpool.")
			fmt.Println("Worker rotine: to get a job from job channel.")
			select {

			case job := <-w.JobChannel:
				fmt.Println("Worker rotine: success: get a job from job channel.")
				// we have received a work request.
				heavyWork(job)

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func heavyWork(job Job) {
	//heavy work
	var jobID = job.Payload
	fmt.Println("job ", jobID, " is working")
	time.Sleep(time.Duration(randSleep()) * time.Second)
}

func randSleep() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10) + 1
}
