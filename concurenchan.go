package main

import(
	"fmt"
	"math"
)


func sum(slice []int, ch1 chan int) {

	result := 0

	for _, n := range slice {
		result += n
	}

	ch1 <- result
}



func abs(a int, ch2 chan int) {

	ch2 <- int(math.Abs(float64(a)))
}



func pow(x int, ch3 chan int) {

	ch3 <- int(math.Pow(float64(x), 2))
}



func main() {

	sum_ch := make(chan int)
	abs_ch := make(chan int)
	pow_ch := make(chan int)
	slice := []int{1,2,3,4,5}

	go sum(slice,sum_ch)
	go abs(-5,abs_ch)
	go pow(4,pow_ch)

	fmt.Println(<- sum_ch)
	fmt.Println(<- abs_ch)
	fmt.Println(<- pow_ch)

}