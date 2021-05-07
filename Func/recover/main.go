package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic happend: ", err)
		}
	}()
	fmt.Println("some userful work")
	panic("something bad happend")
	return
}

func main() {
	deferTest()
	return
}
