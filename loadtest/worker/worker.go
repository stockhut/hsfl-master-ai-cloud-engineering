package worker

import (
	"fmt"
)

// DefaultWorker executes jobs and writes their output ro results until the channel is closed
func DefaultWorker[R any](id int, jobs <-chan Job[R], results chan<- R) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		results <- j.Do()
		fmt.Println("worker", id, "finished job", j)

	}
	fmt.Println("worker", id, "died")
}
