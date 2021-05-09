package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iterationsNum = 7
	gorutinesNum  = 5
)

func doSomeWork(in int) {
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "█", strings.Repeat(" ", gorutinesNum-in), "th", in, "iter", j, strings.Repeat("■", j))
}

func imports() {
	fmt.Println(time.Millisecond, runtime.NumCPU())
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************PLACEHOLDER***************")
	fmt.Println(string(colorReset), "")

	for i := 0; i < gorutinesNum; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}
