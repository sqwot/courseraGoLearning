package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	gorutinesNum  = 5
	iterationsNum = 5
)

func startWorker(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationsNum; j++ {
		fmt.Printf(formatedWork(in, j))
		runtime.Gosched()
	}
}
func formatedWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", gorutinesNum-in),
		"th", in, "iter", j, strings.Repeat("*", j))
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************sync.WaitGroups***************")
	fmt.Println(string(colorReset), "")
	wg := &sync.WaitGroup{}
	for i := 0; i < gorutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg)
	}
	time.Sleep(time.Millisecond)

	wg.Wait()
}
