package main

import(
	"fmt"
)

func F(ch chan <- int) {
	ch <- 100
}


func main() {
	// Unidirectional Channel

	ch := make(chan int)

	go F(ch)

	fmt.Println(<-ch)
}