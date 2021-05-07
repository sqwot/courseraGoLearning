package main

import (
	"fmt"
	"strconv"
)

type Wallet struct {
	Cash int
}
type Payer interface {
	Pay(int) error
}

func (w *Wallet) String() string {
	return "Кошелек в котором " + strconv.Itoa(w.Cash) + " денег"
}
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("No money, no honey")
	}
	w.Cash -= amount
	return nil
}

func Buy(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платежным средством\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	fmt.Printf("Raw payment: %#v\n", myWallet)
	fmt.Printf("Способ оплаты: %s\n", myWallet)
	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)

}
