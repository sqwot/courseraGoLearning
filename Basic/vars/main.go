package main

import "fmt"

type UserID int

func main() {
	//var num0 int

	//var num1 int = 1

	//var num2 = 20

	num := 30

	num += 1
	fmt.Println("+=", num)

	num++
	fmt.Println("++", num)

	const pi = 3.14
	const (
		zero = iota
		_
		_
		_
		five
	)

	const (
		_         = iota
		KB uint64 = 1 << (10 * iota)
		MB
		GB
	)

	idx := 1
	//var iud UserID = 42
	myID := UserID(idx)

	fmt.Println(myID)

	a := 2
	b := &a
	*b = 3
	c := &a

	d := new(int)
	*d = 12
	*c = *d
	*d = 13
	c = d

	fmt.Println(*c)
	fmt.Println(a)
}
