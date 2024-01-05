package main

import (
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

type loadPhase struct {
	Rps      float64
	Duration time.Duration
	Rampup   time.Duration
}

func rpsAfterTime(phases []loadPhase, t time.Duration) float64 {

	fmt.Printf("Finding RPS after %s\n", t)
	totalT := 0 * time.Second
	lastRps := 0.0

	for i, phase := range phases {

		fmt.Printf("looking at phase %d. Duration so far: %s\n", i, totalT)

		if t <= totalT+phase.Rampup {
			fmt.Printf("i am in rampup for phase %d\n", i)
			return lerp(lastRps, phase.Rps, t-totalT, phase.Rampup)
		}
		totalT += phase.Rampup

		totalT += phase.Duration
		if t < totalT {
			fmt.Printf("i am in phase %d\n", i)
			return phase.Rps
		}

		lastRps = phase.Rps
	}

	return -1
}

func lerp(start float64, end float64, elapsed time.Duration, totalDuration time.Duration) float64 {
	fmt.Println("start", start, "end", end, "elapsed", elapsed, "totalDuration", totalDuration)
	percent := float64(elapsed) / float64(totalDuration)

	fmt.Printf("lerp between %f, %f at %f\n", start, end, percent)

	return start + (end-start)*percent
}

func main() {

	phases := []loadPhase{
		{
			Rps:      500,
			Duration: 40 * time.Second,
			Rampup:   60 * time.Second,
		},
		{
			Rps:      2000,
			Duration: 60 * time.Second,
			Rampup:   20 * time.Second,
		},
	}

	var totalTestDuration time.Duration
	for _, p := range phases {
		totalTestDuration += p.Rampup + p.Duration
	}
	fmt.Println(totalTestDuration)

	testStartTime := time.Now()

	quitChan := make(chan any)

	go func() {
		time.Sleep(totalTestDuration)
		quitChan <- nil
	}()

	totalRequests := atomic.Uint32{}

	running := true
	for running {

		select {
		case <-quitChan:
			fmt.Println("Test finished")
			running = false
		case <-time.After(time.Second):
			rps := int(rpsAfterTime(phases, time.Since(testStartTime)))
			fmt.Printf("Current RPS: %d\n", rps)
			for n := 0; n < rps; n++ {
				go func() {
					totalRequests.Add(1)

					jwt := "XXX"

					req, err := http.NewRequest("GET", "http://127.0.0.1:80/api/v1/recipe/by/test", nil)
					if err != nil {
						fmt.Println(err)
					}

					req.Header.Add("Cookie", fmt.Sprintf("jwt=%s", jwt))

					conn, err := net.Dial("tcp", "127.0.0.1:80")
					if err != nil {
						fmt.Println(err)
					}

					fmt.Fprintf(conn, "GET /api/v1/recipe/by/test HTTP/1.1\nHost: 127.0.0.1:80\nCookie: jwt=%s\n\n", jwt)
					conn.Close()
				}()

			}

		}

	}

	fmt.Println(totalRequests.Load())
	//configPath := "load.config.json"
	//conf, err := config.FromFS(configPath)
	//if err != nil {
	//	log.Fatalf("Failed to load config file %s: %s", configPath, err)
	//}
	//
	//numWorkers := conf.Users
	//
	//jobFactory := worker.JobFactoryFunc[any](func() worker.Job[any] {
	//	return worker.JobFunc[any](func() any {
	//		target := randomItemFromSlice(conf.Targets)
	//
	//		_, err := http.Get(target)
	//		if err != nil {
	//			log.Printf("Request to target %s failed: %s", target, err)
	//		}
	//		//time.Sleep(1 * time.Second)
	//		fmt.Println(target)
	//
	//		return true
	//	})
	//})
	//p := worker.NewRampedPool[any](numWorkers, jobFactory)
	//
	//go p.Start()
	//
	//timeout := time.After(time.Duration(conf.Duration) * time.Millisecond)
	//
	//<-timeout
	//p.Stop()

}

//func randomItemFromSlice[T any](ts []T) T {
//	i := rand.Intn(len(ts))
//	return ts[i]
//}
