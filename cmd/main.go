package main

import (
	"fmt"
	"time"

	"github.com/AnimeshRajwar/GoLimiter.git/limiter"
)

func main() {
	var rl limiter.RateLimiter

	rl = limiter.NewFixedWindow(5, 24*time.Second)

	user := "user-123"

	for i := 0; i < 20; i++ {

		allowed := rl.AllowUser(user)

		fmt.Println(
			i,
			allowed,
		)

		time.Sleep(
			5 * time.Second,
		)
		fmt.Println("")

	}

}
