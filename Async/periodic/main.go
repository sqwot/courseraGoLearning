package main

import (
	"fmt"
	"time"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Pereodic***************")
	fmt.Println(string(colorReset), "")

	ticker := time.NewTicker(time.Second)
	i := 0
	for tickTime := range ticker.C {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			ticker.Stop()
			break
		}
	}
	fmt.Println("total", i)

	c := time.Tick(time.Second)
	i = 0
	for tickTime := range c {
		i++
		fmt.Println("step", i, "time", tickTime)
		if i >= 5 {
			break
		}
	}

	timer := time.AfterFunc(5*time.Second, sayHello)
	fmt.Scanln()
	timer.Stop()
	fmt.Scanln()
}

func sayHello() {
	fmt.Println("HelloWorld!")
}
