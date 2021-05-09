package main

import (
	"fmt"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Channels***************")
	fmt.Println(string(colorReset), "")

	// ch1 := make(chan int, 1)

	// go func(in chan int) {
	// 	val := <-in
	// 	fmt.Println("GO: get from chan", val)
	// 	fmt.Println("GO: after read drom chan")

	// }(ch1)

	// ch1 <- 42
	// var input string
	// fmt.Scanln(&input)
	// intInput, err := strconv.Atoi(input)
	// if err != nil {
	// 	fmt.Println("This is not number")
	// }
	// ch1 <- intInput
	// fmt.Println("MAIN: after put to chan")
	in := make(chan int)
	go func(out chan<- int) {
		for i := 0; i <= 4; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finished")
	}(in)
	for i := range in {
		fmt.Println("\tget", i)
	}
	//fmt.Scanln()
}
