package main

import (
	"fmt"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"

	fmt.Println(string(colorGreen), "************First Project***************")
	fmt.Println(string(colorReset), "")
}
