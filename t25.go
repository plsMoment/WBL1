package main

import (
	"fmt"
	"time"
)

// ownSleep represents sleep from package time using package time:)
func ownSleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	start := time.Now()
	ownSleep(5 * time.Second)
	fmt.Printf("Approximatly sleep duration: %f\n", time.Since(start).Seconds())
}
