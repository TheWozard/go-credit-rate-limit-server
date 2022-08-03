package rate_test

import (
	"go-credit-rate-limit-server/pkg/rate"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLimiter(t *testing.T) {
	test_rate := rate.NewRate(1*time.Millisecond, 5)
	defer test_rate.Close()
	limiter := rate.Limiter{
		Accounts: map[string]rate.Rate{
			"test": test_rate,
		},
	}
	time.Sleep(2 * time.Millisecond) // Again super ugly. Would need to abstract ticker to fix
	result, err := limiter.Acquire("bad", 20)
	require.Equal(t, 0, result)
	require.NotNil(t, err)
	require.Equal(t, "no account named bad", err.Error())

	result, err = limiter.Acquire("test", 10)
	require.Equal(t, 5, result)
	require.Nil(t, err)

	result, err = limiter.Acquire("test", 10)
	require.Equal(t, 0, result)
	require.Nil(t, err)
}
