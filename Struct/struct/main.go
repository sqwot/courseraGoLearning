package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string)
	Owner   Person
}

func main() {
	var acc Account = Account{
		Id:   1,
		Name: "rvasiliy",
		Cleaner: func(string) {
		},
		Owner: Person{},
	}

	fmt.Printf("%#v\n", acc)
	acc.Owner = Person{2, "Romanov Vasily", "Moscow"}
	fmt.Printf("%#v\n", acc)
}
