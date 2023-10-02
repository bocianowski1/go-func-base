package function

import (
	"fmt"
	"net/http"
	"sync"
)

type WorkerPool struct {
	workers   []*Worker
	workQueue chan *http.Request
	wg        sync.WaitGroup
	shutdown  chan struct{}
}

type Worker struct {
	id        int
	workQueue chan *http.Request
	shutdown  chan struct{}
}

func NewWorkerPool(poolSize int) *WorkerPool {
	pool := &WorkerPool{
		workQueue: make(chan *http.Request),
		shutdown:  make(chan struct{}),
	}

	for i := 1; i <= poolSize; i++ {
		worker := &Worker{
			id:        i,
			workQueue: pool.workQueue,
			shutdown:  pool.shutdown,
		}
		pool.workers = append(pool.workers, worker)

		go worker.Start()
	}

	return pool
}

func (wp *WorkerPool) Enqueue(request *http.Request) {
	wp.workQueue <- request
}

func (wp *WorkerPool) Shutdown() {
	close(wp.shutdown)
	wp.wg.Wait()
}

func (w *Worker) Start() {
	for {
		select {
		case request := <-w.workQueue:
			fmt.Printf("Worker %d processing request for URL: %s\n", w.id, request.URL.Path)
		case <-w.shutdown:
			fmt.Printf("Worker %d shutting down.\n", w.id)
			return
		}
	}
}
