package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Starting program")

	// sig channel receives Unix signals
	sig := make(chan os.Signal, 1)
	// done channel is used to notify when the program can exit
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals.
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// program will block waiting for a signal
		signal := <-sig
		fmt.Println()
		fmt.Println(signal)
		// send a value to notify that we are done
		done <- true
	}()

	fmt.Println("awaiting signal")
	// block from exiting until we receive a notification.
	// If this wasn't here, the program would exit before the above goroutine would even run
	<-done
	fmt.Println("exiting")
}
