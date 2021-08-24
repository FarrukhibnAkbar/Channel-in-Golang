package main

import(
	"fmt"
)


func main() {

	// i := 0

	// fmt.Println("A",i)

	// i = 1

	// fmt.Println("B", i)

	// i = 2

	// fmt.Println("Finish", i)



	j := 0
	
	defer	fmt.Println("A",j)

	j = 1

	defer fmt.Println("B", j)

	j = 2

	defer fmt.Println("Finish", j)


	ch := make(chan int, 2)


	fmt.Printf("%v %T",<-ch,<-ch)
}