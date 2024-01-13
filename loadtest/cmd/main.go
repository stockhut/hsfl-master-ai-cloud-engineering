package main

import (
	"errors"
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/loadtest/config"
	"log"
	"net"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

func httpStatus(re *regexp.Regexp, buff []byte) (int, error) {

	matches := re.FindStringSubmatch(string(buff))
	if len(matches) < 2 {
		return 0, errors.New("not enough matches")
	}

	return strconv.Atoi(matches[1])
}

func main() {

	configFilePath := "loadtest.yaml"
	cfg, err := config.FromFile("loadtest.yaml")
	if err != nil {
		log.Fatalf("Failed to load config file %s: %s\n", configFilePath, err)
	}

	gatherResponseStats := cfg.ResponseStats

	phases := fun.Map(cfg.Phases, func(p config.Phase) loadtest.Phase {
		return loadtest.Phase{
			Rps:      float64(p.Rps),
			Rampup:   p.Rampup,
			Duration: p.Duration,
		}
	})

	headersMap := cfg.Headers
	_, hasHostHeader := headersMap["Host"]
	if !hasHostHeader {
		headersMap["Host"] = cfg.Host
	}

	headers := loadtest.HeadersToString(headersMap)

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
				rps := int(loadtest.RpsAfterTime(phases, batchStartTime))

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

						path := loadtest.RandomItemFromSlice(cfg.Targets)
						req := []byte(fmt.Sprintf("GET %s HTTP/1.1\n%s\n\n", path, headers))

						var responseBuff []byte
						var requestStartTime time.Time
						if gatherResponseStats {
							responseBuff = make([]byte, 16)
							requestStartTime = time.Now()
						}

						conn, err := net.Dial("tcp", cfg.Host)
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

					numErrorCodes := fun.Count(responseStatusCodes, loadtest.HttpStatusIsError)

					avgResponseTime := float64(responseTimes.Load()) / float64(r)
					log.Printf("%s: %d RPS %d Errors %0.2fms avg", batchStartTime.Round(time.Second), rps, numErrorCodes, avgResponseTime)
				} else {
					log.Printf("%s: %d RPS", batchStartTime.Round(time.Second), rps)
				}
			}()

		}

	}

	fmt.Println(totalRequests.Load())
}
