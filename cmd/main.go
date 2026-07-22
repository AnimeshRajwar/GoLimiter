package main

import (
	"fmt"
	"time"

	"github.com/AnimeshRajwar/GoLimiter.git/limiter"
)

func main() {

	rl := limiter.NewFixedWindow(5, 24*time.Second, 40*time.Second)
	rl.StartCleanupWorker()
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
