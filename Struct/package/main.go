package main

import (
	"fmt"
	"D:/golang/Coursera/Struct/package/person"
)

func main() {
	fmt.Println("***********Package**********")
	p := person.NewPerson(1, "rvasily", "secret")
	fmt.Printf("main.PrintPerson: %+v\n", p.secret)
	secret := person.GetSecret(p)
	fmt.Println("GetSecret", secret)
}
