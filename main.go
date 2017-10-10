package main

import (
	"fmt"

	"github.com/georgicodes/hello-go-concurrency/pinger"
)

func main() {
	fmt.Println("Starting program")
	pinger.Run()
	fmt.Println("\nTerminating program")
}
