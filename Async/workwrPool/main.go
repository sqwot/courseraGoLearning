package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goruntinesNum = 3

func startWorker(workerNum int, in <-chan string) {
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched()
	}
	printFinishWork(workerNum)
}

func printFinishWork(num int) {
	fmt.Println(num, "finished")
}

func formatWork(in int, j string) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "||",
		strings.Repeat(" ", goruntinesNum-in),
		"th", in,
		"iter", j)
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Workers pool***************")
	fmt.Println(string(colorReset), "")

	workerInput := make(chan string, 2)
	for i := 0; i < goruntinesNum; i++ {
		go startWorker(i, workerInput)
	}

	months := []string{"Январь", "Февраль", "Март", "Апрель", "Май",
		"Июнь", "Июль", "Август", "Сентябрь", "Октябрь",
		"Ноябрь", "Декабрь"}

	for _, monthName := range months {
		workerInput <- monthName
	}
	close(workerInput)
	time.Sleep(time.Millisecond)

}
