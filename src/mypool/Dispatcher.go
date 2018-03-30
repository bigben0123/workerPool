package mypool

import "fmt"

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	MaxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	JobQueue = make(chan Job, 100)

	pool := make(chan chan Job, maxWorkers)

	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	// starting n number of workers
	fmt.Printf("starting %d number of workers\n", d.MaxWorkers)
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	fmt.Printf("start finished.")

	go d.dispatch()
}

/*base on worker to select a job.*/
func (d *Dispatcher) dispatch() {
	fmt.Println("start dispatch.")
	for {
		// try to obtain a worker job channel that is available.
		// this will block until a worker is idle
		fmt.Println("dispatch: select to get a job channel from workpool")
		select {
		case jobChannel := <-d.WorkerPool:
			fmt.Println("dispatch: success select to get a job channel from workpool")

			go func(jobChannel chan Job) {
				fmt.Println("dispatch rotine: to get one job.")
				job := <-JobQueue
				// a job request has been received
				fmt.Println("dispatch rotine: success to get one job.")
				// dispatch the job to the worker job channel
				fmt.Println("dispatch rotine: to put job to job channel.")
				jobChannel <- job
				fmt.Println("dispatch rotine: success: put job to job channel.")
			}(jobChannel)
		}
	}
}

/*base on job to select a worker.*/
func (d *Dispatcher) dispatch0() {
	fmt.Println("start dispatch.")
	for {
		fmt.Println("dispatch: select.")
		select {
		case job := <-JobQueue:
			fmt.Println("dispatch: got one job.")
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				fmt.Println("dispatch rotine: to get a job channel from workpool.")
				jobChannel := <-d.WorkerPool
				fmt.Println("dispatch rotine: success: get a job channel from workpool.")
				// dispatch the job to the worker job channel
				fmt.Println("dispatch rotine: to put job to job channel.")
				jobChannel <- job
				fmt.Println("dispatch rotine: success: put job to job channel.")
			}(job)
		}
	}
}

func (d *Dispatcher) MakeJob(jobID int) Job {
	return Job{Payload: jobID}
}
