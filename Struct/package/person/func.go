package person

import "fmt"

func NewPerson(id int, name string, sec string) *Person {
	return &Person{
		Id:     id,
		Name:   name,
		secret: sec,
	}
}
func GetSecret(p *Person) string {
	return p.secret
}
func printSecret(p *Person) {
	fmt.Println(p.secret)
}
