package concurrency

import (
	"time"
)

func doSomething() {}

func ExampleLimiter() {
	if Limiter("image processing", 4, 1000).Do(time.Second) != nil {
		return
	}
	doSomething()
}
