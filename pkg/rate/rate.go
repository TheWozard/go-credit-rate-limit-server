package rate

import (
	"io"
	"sync"
	"time"
)

type Rate interface {
	io.Closer
	// Attempts to acquire n credits. Returns the number of credits acquired
	Acquire(credits int) int
}

// NewRate creates a new memory based rate and opens a new goroutine to maintain it
func NewRate(interval time.Duration, amount int) Rate {
	rate := &memoryRate{
		acquired: amount,
		max:      amount,
	}
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			rate.Clear()
		}
	}()
	rate.close = ticker.Stop

	return rate
}

type memoryRate struct {
	sync.Mutex
	acquired int
	max      int
	close    func()
}

func (r *memoryRate) Acquire(credits int) int {
	r.Lock()
	defer r.Unlock()

	r.acquired += credits
	if r.acquired > r.max {
		actual := credits - (r.acquired - r.max)
		if actual < 0 {
			return 0
		}
		return actual
	} else {
		return credits
	}

}

func (r *memoryRate) Clear() {
	r.Lock()
	defer r.Unlock()
	r.acquired = 0
}

func (r *memoryRate) Close() error {
	if r.close != nil {
		r.close()
	}
	return nil
}
