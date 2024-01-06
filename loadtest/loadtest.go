package loadtest

import "time"

type Phase struct {
	Rps      float64
	Duration time.Duration
	Rampup   time.Duration
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
