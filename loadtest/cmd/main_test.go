package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_rpsAtTime(t *testing.T) {

	phases := []loadPhase{
		{
			Rps:      100,
			Duration: 10 * time.Second,
			Rampup:   3 * time.Second,
		},
		{
			Rps:      200,
			Duration: 10 * time.Second,
			Rampup:   3 * time.Second,
		},
	}

	testCases := []struct {
		name string
		t    time.Duration
		rps  float64
	}{
		{
			name: "rampup phase 1",
			t:    1500 * time.Millisecond,
			rps:  50,
		},
		{
			name: "phase 1",
			t:    4 * time.Second,
			rps:  100,
		},
		{
			name: "rampup phase 2",
			t:    14500 * time.Millisecond,
			rps:  150,
		},
		{
			name: "phase 2",
			t:    18 * time.Second,
			rps:  200,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			rps := rpsAfterTime(phases, tc.t)

			assert.Equal(t, tc.rps, rps)
		})
	}

}

func Test_lerp(t *testing.T) {

	start := 0.0
	end := 100.0

	totalTime := 10 * time.Second
	elapsed := 5 * time.Second

	assert.Equal(t, 50.0, lerp(start, end, elapsed, totalTime))
}
