package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************PLACEHOLDER***************")
	fmt.Println(string(colorReset), "")
	//alreadySean := make(map[string]bool)
	var prev string
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		txt := in.Text()
		//if _, found := alreadySean[txt]; found {
		//	continue
		//}
		//alreadySean[txt] = true
		if txt == prev {
			continue
		}
		if txt < prev {
			panic("file notsorted")
		}
		prev = txt
		fmt.Println(txt)
	}
}
