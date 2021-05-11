package main

import (
	"fmt"
	"time"
)

func getComments() chan string {
	result := make(chan string, 1)
	go func(out chan<- string) {
		time.Sleep(2 * time.Second)
		fmt.Println("async operation ready, return comments")
		out <- "32 комментария"
	}(result)
	return result
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Async get data***************")
	fmt.Println(string(colorReset), "")

	resultCh := getComments()
	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	commentsData := <-resultCh
	fmt.Println("main gorutine:", commentsData)

}
