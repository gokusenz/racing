package main

import "fmt"

func main() {
	fmt.Println(getNumber(5))
}

func getNumber(n int) int {
	var i int
	// Create a channel to push an empty struct to once we're done
	done := make(chan struct{})
	go func() {
		i = n
		// Push an empty struct once we're done
		done <- struct{}{}
	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done
	return i
}
