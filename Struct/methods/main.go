package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

type MySlice []int

func (m *MySlice) Add(val int) {
	*m = append(*m, val)
	return
}
func (m *MySlice) Size() int {
	return len(*m)
}

func main() {
	fmt.Println("*******Methods********")
	person := Person{
		Id:   1,
		Name: "firstName",
	}
	fmt.Printf("%#v\n", person)
	person.UpdateName("SecondName")
	fmt.Printf("%#v\n", person)
	person.SetName("ThirdName")
	fmt.Printf("%#v\n", person)

	acc := Account{
		Id:   2,
		Name: "accFirstName",
	}
	acc.SetName("assSecondName")
	fmt.Printf("%#v\n", acc)
	sl := MySlice{1, 2, 3, 4}
	sl.Add(5)
	fmt.Printf("%#v: %d\n", sl, sl.Size())
}
