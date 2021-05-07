package main

import "fmt"

func main() {
	fmt.Println("******Arrays********")
	var a1 [3]int
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 full %#v\n", a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println("a2", a2)

	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)

	fmt.Println("******Slice********")
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)

	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)
	someInt := buf2[0]
	fmt.Println(someInt)

	var buf []int
	buf = append(buf, 9, 10)
	buf = append(buf, 12)
	fmt.Println(cap(buf))
	fmt.Println(buf)

	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)

	fmt.Println(buf, otherBuf)

	buf = []int{1, 2, 3, 4, 5}
	fmt.Println(buf[1:4], buf[:2], buf[2:])
	newBuf := buf[:]
	newBuf[0] = 9
	newBuf = append(newBuf, 6)
	fmt.Println(buf)
	fmt.Println(newBuf)

	//wrong
	var emptyBuf []int
	copied := copy(emptyBuf, buf)
	fmt.Println(copied, emptyBuf)

	//right
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf)

	ints := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ints[1:3])
	copy(ints[1:5], []int{5, 6})
	fmt.Println(ints)

	numbers := []int{0, 1, 2, 3, 4}
	fmt.Println(numbers[1:4])

	fmt.Println("******Map********")
	var user map[string]string = map[string]string{
		"name":     "Vasiliy",
		"lastName": "Romanov",
	}
	profile := make(map[string]string, 10)
	mapLen := len(user)
	fmt.Printf("%d, %+v\n", mapLen, profile)

	mName := user["middleName"]
	fmt.Println("mName: ", mName)
	mName, mNameExist := user["middleName"]
	fmt.Println("mName: ", mName, "mNameExist:", mNameExist)

	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
