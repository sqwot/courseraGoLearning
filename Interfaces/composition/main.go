package main

import "fmt"

type Player interface {
	Pay(int) error
}
type Ringer interface {
	Ring(string) error
}
type NFCPhone interface {
	Player
	Ringer
}
type Phone struct {
	Money   int
	AppleID string
}

func (a *Phone) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("No money on account")
	}
	a.Money -= amount
	return nil
}
func (p *Phone) Ring(string) error {
	err := p.Ring("sqwot")
	if err != nil {
		return fmt.Errorf("Ошибка вызова")
	}
	fmt.Println("Вызов принят")
	return nil
}
func PayForMetroWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("Ошибка при оплате %v\n\n")
		return
	}
	fmt.Printf("Турникет открыт через %T\n\n", phone)
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"

	fmt.Println(string(colorGreen), "****Composition of innterfaces****")
	fmt.Println(string(colorReset), "")
	myPhone := &Phone{
		Money:   9,
		AppleID: "",
	}
	PayForMetroWithPhone(myPhone)
}
