package main

import (
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"log"
	"net"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type loadPhase struct {
	Rps      float64
	Duration time.Duration
	Rampup   time.Duration
}

func rpsAfterTime(phases []loadPhase, t time.Duration) float64 {

	//fmt.Printf("Finding RPS after %s\n", t)
	totalT := 0 * time.Second
	lastRps := 0.0

	for _, phase := range phases {

		//fmt.Printf("looking at phase %d. Duration so far: %s\n", i, totalT)

		if t <= totalT+phase.Rampup {
			//fmt.Printf("i am in rampup for phase %d\n", i)
			return lerp(lastRps, phase.Rps, t-totalT, phase.Rampup)
		}
		totalT += phase.Rampup

		totalT += phase.Duration
		if t < totalT {
			//fmt.Printf("i am in phase %d\n", i)
			return phase.Rps
		}

		lastRps = phase.Rps
	}

	return -1
}

func lerp(start float64, end float64, elapsed time.Duration, totalDuration time.Duration) float64 {
	//fmt.Println("start", start, "end", end, "elapsed", elapsed, "totalDuration", totalDuration)
	percent := float64(elapsed) / float64(totalDuration)

	//fmt.Printf("lerp between %f, %f at %f\n", start, end, percent)

	return start + (end-start)*percent
}

func httpStatus(re *regexp.Regexp, buff []byte) (int, error) {

	matches := re.FindStringSubmatch(string(buff))
	if len(matches) < 2 {
		return 0, errors.New("not enough matches")
	}

	return strconv.Atoi(matches[1])
}

func httpStatusIsError(code int) bool {
	return code/100 != 2 && code/100 != 3
}

func main() {

	gatherResponseStats := false

	phases := []loadPhase{
		{
			Rps:      2500,
			Duration: 40 * time.Second,
			Rampup:   10 * time.Second,
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

	httpStatusRegex, err := regexp.Compile(`HTTP/\S+ ([[:digit:]]+)`)
	if err != nil {
		log.Printf("Failed to compile status code regex: %s\n", err)
	}

	running := true
	for running {

		select {
		case <-quitChan:
			log.Println("Test finished")
			running = false
		case <-time.After(time.Second):
			go func() {
				batchStartTime := time.Since(testStartTime)
				rps := int(rpsAfterTime(phases, batchStartTime))

				wg := sync.WaitGroup{}

				var responseStatusCodes []int
				if gatherResponseStats {
					responseStatusCodes = make([]int, rps)
				}

				requests := atomic.Uint32{}
				responseTimes := atomic.Uint64{}

				for n := 0; n < rps; n++ {

					go func(idx int, stats []int) {
						requests.Add(1)
						wg.Add(1)

						jwt := "XX"

						req := []byte(fmt.Sprintf("GET /api/v1/recipe/by/test HTTP/1.1\nHost: 127.0.0.1:80\nCookie: jwt=%s\n\n", jwt))

						var responseBuff []byte
						var requestStartTime time.Time
						if gatherResponseStats {
							responseBuff = make([]byte, 16)
							requestStartTime = time.Now()
						}

						conn, err := net.Dial("tcp", "127.0.0.1:80")
						if err != nil {
							log.Println(err)
							return
						}
						_, err = conn.Write(req)
						if err != nil {
							log.Println(err)
						}

						if gatherResponseStats {
							_, err = conn.Read(responseBuff)
							if err != nil {
								log.Println(err)
							}
							err = conn.Close()
							if err != nil {
								log.Println(err)
							}
							rt := uint64(time.Since(requestStartTime).Milliseconds())
							responseTimes.Add(rt)

							status, err := httpStatus(httpStatusRegex, responseBuff)
							if err != nil {
								log.Println(err)
							}

							stats[idx] = status
						}

						wg.Done()
					}(n, responseStatusCodes)

				}

				if gatherResponseStats {
					wg.Wait()
					r := requests.Load()
					totalRequests.Add(r)

					numErrorCodes := fun.Count(responseStatusCodes, httpStatusIsError)

					avgResponseTime := float64(responseTimes.Load()) / float64(r)
					log.Printf("%s: %d RPS %d Errors %0.2fms avg", batchStartTime.Round(time.Second), rps, numErrorCodes, avgResponseTime)
				} else {
					log.Printf("%s: %d RPS", batchStartTime.Round(time.Second), rps)
				}
			}()

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
