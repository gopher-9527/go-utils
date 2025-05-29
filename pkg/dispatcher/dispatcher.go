package dispatcher

import "sync"

type Job func() any

type Dispatcher interface{}

type dispatcher struct {
	jobQueue chan Job
	workers  int
	result   chan any
	done     chan struct{}
	closed   bool
	lock     sync.Mutex
}

func (d *dispatcher) NewDispatch(queueLength, workers int) Dispatcher {
	return &dispatcher{
		jobQueue: make(chan Job, queueLength),
		workers:  workers,
		done:     make(chan struct{}),
		result:   make(chan any),
	}
}

func (d *dispatcher) Run() (chan any, chan struct{}) {
	d.closed = false
	for i := 0; i < d.workers; i++ {
		go func() {
			for {
				select {
				case job := <-d.jobQueue:
					r := job()
					d.result <- r
				case <-d.done:
					return
				}
			}
		}()
	}
	return d.result, d.done
}

func (d *dispatcher) AddJob(job Job) {
	if d.closed {
		return
	}
	d.jobQueue <- job
}

func (d *dispatcher) Close() {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.closed {
		return
	}
	d.closed = true
	close(d.done)
	close(d.jobQueue)
	close(d.result)
}
