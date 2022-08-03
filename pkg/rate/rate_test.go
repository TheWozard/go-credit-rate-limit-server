package rate_test

import (
	"fmt"
	"go-credit-rate-limit-server/pkg/rate"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRateCreation(t *testing.T) {
	rate := rate.NewRate(1*time.Millisecond, 15)
	defer rate.Close()
	require.NotNil(t, rate)
}

func TestRateAcquisition(t *testing.T) {
	type acquisition struct {
		input  int
		output int
	}

	testCases := []struct {
		desc     string
		interval time.Duration
		max      int
		acquire  []acquisition
		repeats  int
	}{
		{
			desc:     "basic usage",
			interval: 1 * time.Millisecond,
			max:      15,
			acquire: []acquisition{
				{input: 10, output: 10},
			},
			repeats: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rate := rate.NewRate(tC.interval, tC.max)
			defer rate.Close()
			// All rates are created as at capacity.
			require.Equal(t, rate.Acquire(10), 0)
			for repeat := 0; repeat < tC.repeats; repeat++ {
				// Ugly but kinda works.
				time.Sleep(2 * tC.interval)
				for i, request := range tC.acquire {
					require.Equal(t, request.output, rate.Acquire(request.input), fmt.Sprintf("failed request %d/%d", repeat, i))
				}
			}
		})
	}
}
