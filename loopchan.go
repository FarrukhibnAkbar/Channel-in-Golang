package main

import(
	"fmt"
)

func channel(ch chan int) {

	for i := 0; i < 3; i++ {
		ch <- i
	}

	close(ch)
}


func main() {

	ch := make(chan int)

	go channel(ch)
	
	for n := range ch {

		fmt.Println(n)
	}
}