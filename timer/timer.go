package timer

import (
	"fmt"
	"time"
)

// TimeIt returns a function that when deferred, prints the elapsed time in nanoseconds
func TimeIt(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("Function %s took %d ns\n", name, elapsed.Nanoseconds())
	}
}
