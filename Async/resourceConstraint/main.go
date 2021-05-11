package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	itertionsNum  = 6
	gorutinesNums = 10
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{}
	defer wg.Done()

	for j := 0; j < itertionsNum; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
		if j%2 == 0 {
			<-quotaCh
			quotaCh <- struct{}{}
		}
	}
	<-quotaCh
}
func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", gorutinesNums-in),
		"th", in, "iter", j, strings.Repeat("*", j))
}
func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Rate limit***************")
	fmt.Println(string(colorReset), "")

	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < gorutinesNums; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()
}
