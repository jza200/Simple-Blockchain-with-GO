package work_queue

import(
	//"fmt"
)

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs    chan Worker
	Results chan interface{}
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	q.Jobs = make(chan Worker, maxJobs) //buffered channel so items in Jobs does not exceed maxJobs
	q.Results = make(chan interface{})
	for i := 0; i < int(nWorkers); i++ {
		go q.worker()
		//fmt.Println("This is Create: ", i)
	}
	return q
}

// A worker goroutine that processes tasks from .Jobs (until shut down).
func (queue WorkQueue) worker() {
	//the channel checking method is adopted from stack overflow and go by example
	for i := range queue.Jobs {
		myWorker := i
		queue.Results <- myWorker.Run()
	}

	return
	// TODO: Listen on the .Jobs channel for incoming tasks. For each task...
	// TODO: run tasks by calling .Run(),
	// TODO: send the return value back on Results channel.
	// TODO: Exit (return) when .Jobs is closed.
}

func (queue WorkQueue) Enqueue(work Worker) {
	queue.Jobs <- work
	// TODO: put the work into the Jobs channel so a worker can find it and start the task.
}

func (queue WorkQueue) Shutdown() {
	close(queue.Jobs)
	// TODO: close .Jobs and remove all remaining jobs from the channel.
}
