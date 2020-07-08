//main() executes two anonymous functions. Keyword 'go' makes them run concurrently, in goroutines.
// Go runtime scheduler schedules goroutines in a non-deterministic way so there are multiple possible interleavings -
// each time we run the application scheduler might change the order of their execution. Sometimes function which
// increments x will be executed first and sometimes the function which prints out its value. Because both functions
// in goroutines are communicating through variable x there is a chance of race condition - the order of accessing
// variable x is non-deterministic. If we run the application multiple times, it will sometimes print 0 and sometimes 1.
// My test were showing that majority of times it would print 1 but sometimes it would print 0. This randomness in
// output proves there is a race condition.

package 

import (
	"fmt"
	"time"
)

var x int 

func main() {

	go func() {
		x = x+1
	}()
	go func() {
		fmt.Print("x =",x)
	}()

	time.Sleep(1* time.Second)
}