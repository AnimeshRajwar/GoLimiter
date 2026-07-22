package limiter

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindow struct {
	limit  int
	window time.Duration

	requests map[string]int
	start    map[string]time.Time

	mutex sync.Mutex

	ttl time.Duration
}

func NewFixedWindow(limit int, window time.Duration, ttl time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:    limit,
		window:   window,
		requests: make(map[string]int),
		start:    make(map[string]time.Time),
		ttl:      ttl,
	}
}

func (f *FixedWindow) AllowUser(key string) bool {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	now := time.Now()

	f.Cleanup()

	fmt.Printf("Request time: %v\n", now)

	if _, ok := f.start[key]; !ok {
		fmt.Println("New User")
		f.start[key] = now
		f.requests[key] = 0
	} else if now.After(f.start[key].Add(f.window)) {
		fmt.Println("New Window")
		f.start[key] = now
		f.requests[key] = 0
	}

	if f.requests[key] == f.limit {
		fmt.Println("Limit Reached")
		return false
	}

	f.requests[key]++
	fmt.Printf("Allowing request: %d\n", f.requests[key])
	return true
}

func (f *FixedWindow) RemainingRequests(key string) int {
	return f.limit - f.requests[key]
}

func (f *FixedWindow) Cleanup() {
	cutoff := time.Now().Add(-f.ttl)

	for key, startTime := range f.start {
		if startTime.Before(cutoff) {
			delete(f.start, key)
			delete(f.requests, key)
		}
	}
}
func (f *FixedWindow) StartCleanupWorker() {
	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			f.Cleanup()
		}
	}()
}
