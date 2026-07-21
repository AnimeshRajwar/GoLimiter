package limiter

type RateLimiter interface {
	AllowUser(key string) bool
	RemainingRequests(key string) int
}
