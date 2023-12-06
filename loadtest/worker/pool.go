package worker

import (
	"fmt"
	"time"
)

type Pool[R any] interface {
	Start()
	Stop()
}

// rampedWorkerPool is a pool that increases it's number of workers until a limit is reached, which it then will hold for a specified duration
type rampedWorkerPool[R any] struct {
	jobs       chan Job[R]
	results    chan R
	rampUpTime int
	size       int
	job        Job[R]
	cancel     chan any
}

func NewRampedPool[R any](numWorkers int, job Job[R]) Pool[R] {
	return &rampedWorkerPool[R]{
		jobs:    make(chan Job[R], 1),
		results: make(chan R, numWorkers),
		size:    numWorkers,
		job:     job,
		cancel:  make(chan any),
	}
}

func (p *rampedWorkerPool[R]) Stop() {
	fmt.Println("stopping pool")
	p.cancel <- true
}

func (p *rampedWorkerPool[R]) Start() {

	p.rampUp()

	for {
		r := <-p.results
		fmt.Println("r", r)
		p.jobs <- p.job
	}
}

func (p *rampedWorkerPool[R]) rampUp() {

	workersPerStep := 100

	increaseEvery := p.rampUpTime / (p.size / workersPerStep)
	steps := p.size / workersPerStep

	workerCount := 0
	for w := 1; w <= steps; w++ {
		fmt.Printf("spawning %d new workers\n", workersPerStep)
		for i := 0; i < workersPerStep; i++ {
			p.jobs <- p.job
			go DefaultWorker[R](workerCount, p.jobs, p.results)

			workerCount++
		}

		time.Sleep(time.Duration(increaseEvery) * time.Millisecond)
	}
	fmt.Printf("got %d workers\n", workerCount)

}
