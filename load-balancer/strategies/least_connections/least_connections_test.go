package least_connections

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestLeastConnections_GetTarget(t *testing.T) {

	t.Run("should pick first host with 0 connections", func(t *testing.T) {
		strategy := &LeastConnections{
			connectionCount: map[string]uint32{
				"a": 0,
				"b": 0,
				"c": 0,
			},
			m: &sync.RWMutex{},
		}

		replicas := []string{
			"a",
			"b",
			"c",
		}

		var picked string
		strategy.GetTarget(nil, replicas, func(host string) {
			picked = host
		})

		assert.Equal(t, "a", picked)
	})

	t.Run("should pick host with fewest connections", func(t *testing.T) {

		type testCase struct {
			connectionCount map[string]uint32
			expectedHost    string
		}
		testCases := []testCase{
			{
				connectionCount: map[string]uint32{
					"a": 0,
					"b": 10,
					"c": 10,
				},
				expectedHost: "a",
			},
			{
				connectionCount: map[string]uint32{
					"a": 10,
					"b": 10,
				},
				expectedHost: "c",
			},
			{
				connectionCount: map[string]uint32{
					"a": 10,
					"b": 10,
					"c": 2,
				},
				expectedHost: "c",
			},
		}

		for i, tc := range testCases {
			t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
				strategy := &LeastConnections{
					connectionCount: tc.connectionCount,
					m:               &sync.RWMutex{},
				}

				replicas := []string{
					"a",
					"b",
					"c",
				}

				var picked string
				strategy.GetTarget(nil, replicas, func(host string) {
					picked = host
				})

				assert.Equal(t, tc.expectedHost, picked)
			})

		}

	})

}
