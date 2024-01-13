package loadtest

import (
	"fmt"
	"github.com/stockhut/hsfl-master-ai-cloud-engineering/common/fun"
	"io"
	"math/rand"
	"strings"
	"time"
)

type Phase struct {
	Rps      float64
	Duration time.Duration
	Rampup   time.Duration
}

type Target struct {
	Method string
	Path   string
	Body   []byte
}

// MakeHttpRequest builds an HTTP Request from a Target and pre-formatted common headers and writes it to the
// given io.Writer, e.g. a net.Conn
func MakeHttpRequest(writer io.Writer, headers string, target Target) error {
	method := strings.ToUpper(target.Method)

	contentHeader := fmt.Sprintf("Content-Length: %d", len(target.Body))
	head := fmt.Sprintf("%s %s HTTP/1.1\n%s\n%s\n\n", method, target.Path, headers, contentHeader)

	_, err := writer.Write([]byte(head))
	if err != nil {
		return err
	}
	_, err = writer.Write(target.Body)

	return err
}

func RpsAfterTime(phases []Phase, t time.Duration) float64 {

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

func HeadersToString(headers map[string]string) string {
	headersStrings := fun.MapToSlice(headers, func(name string, value string) string {
		return name + ": " + value
	})

	return strings.Join(headersStrings, "\n")
}

func RandomItemFromSlice[T any](ts []T) T {
	i := rand.Intn(len(ts))
	return ts[i]
}

func HttpStatusIsError(code int) bool {
	return code/100 != 2 && code/100 != 3
}
