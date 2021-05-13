package main

import (
	"fmt"
)

//cgen: binpack
type User struct {
	ID       int
	RealName string `cgen:"-"`
	Login    string
	Flags    int
}

type Avatar struct {
	ID  int
	Url string
}

var test = 42

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Code generation***************")
	fmt.Println(string(colorReset), "")

	data := []byte{
		128, 36, 17, 0,

		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}

	u := new(User)
	u.Unpack(data)
	fmt.Printf("Unpacked user %#v", u)
}
