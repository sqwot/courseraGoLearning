package main

import "fmt"

func main() {
	boolVal := true
	if boolVal {
		fmt.Println("BoolVal is true")
	}

	mapVal := map[string]string{"name": "rvasily"}
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name =", keyValue)
	}

	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else {
		fmt.Println("cond is not 1")
	}
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough
	case "test", "lastName":
		//sd
	default:
		fmt.Println("ты пидор")

	}

	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("firstblock")
	case val2 > 10:
		fmt.Println("secodblock")
	}
Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			println("Dont pront this")
		case key == "firstName" && val == "vasiliy":
			println("switch - break loop here")
			break Loop
		}
	}
	println("****LOOPS****")

	for {
		fmt.Println("loop iteration")
		break
	}

	isRun := true
	for isRun {
		println("loop iteration with condition")
		isRun = false
	}
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration #", i)
		if i == 1 {
			continue
		}
	}

	s1 := []int{1, 2, 3}
	idx := 0
	for idx < len(s1) {
		fmt.Println("while-stile loop, idx:", idx, "value:", s1[idx])
		idx++
	}
	for idx, val := range s1 {
		fmt.Println("range slixe by idx value", idx, val)
	}

	str := "Hello World!"

	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
