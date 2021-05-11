package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperations int32 = 0

func inc() {
	atomic.AddInt32(&totalOperations, 1)
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Atomic***************")
	fmt.Println(string(colorReset), "")

	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation = ", totalOperations)
}
