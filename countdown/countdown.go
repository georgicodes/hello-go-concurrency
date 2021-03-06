package countdown

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// When set to 1, a single logical processor for scheduler to use,
	// meaning only 1 goroutine is running at any given time. The scheduler will
	// timeslice goroutines on a single thread.
	//
	// When set to 2, the goroutines will run in parallel as there will be
	// 2 processors with a thread each.
	runtime.GOMAXPROCS(2)
}

// Run counts up and down
func Run() {
	// a WaitGroup is a synchronous counting semaphore, used to manage concurrency
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutines")

	// create an annoymous goroutine
	// go here will schedule this goroutine to run in the LRQ on P
	go func() {
		countdown()
		wg.Done()
	}()

	// create another annoymous goroutine
	go func() {
		countup()
		wg.Done()
	}()

	// at this point we have 3 goroutines running concurrently (main, lowercase, uppercase)

	fmt.Println("waiting to finish")
	// when main exits, the whole program exits, wg.Wait(), will wait for the goroutines
	// to finish. It prevents main from terminating until other two goroutines are finished.
	// When that happens it will wake up the main goroutine and terminate.
	wg.Wait()

	fmt.Println("\nTerminating program")
}

func countdown() {
	fmt.Println("countdown")
	i := 100
	for i >= 0 {
		fmt.Printf("%d:%s ", i, "UP")
		i = i - 1
	}
}

func countup() {
	fmt.Println("countup")
	i := 0
	for i <= 100 {
		fmt.Printf("%d:%s ", i, "dn")
		i = i + 1
	}
}
