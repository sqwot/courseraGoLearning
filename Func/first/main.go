package main

import "fmt"

func singleIn(in int) int {
	return in
}

func multiIn(a, b int, c int) int {
	return a + b + c
}

func namedReturn() (out int) {
	out = 2
	return
}

func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

func multiplrNamedReturn(ok bool) (rez int, err error) {
	if ok {
		err = fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {
	a, err := multiplrNamedReturn(true)
	fmt.Println(a, err)
	fmt.Println(sum(3, 4, 5))
}
