package main

import "fmt"

func main() {
	fmt.Println(getNumber(5))
}

func getNumber(n int) int {
	var i int
	go func() {
		i = n
	}()

	return i
}
