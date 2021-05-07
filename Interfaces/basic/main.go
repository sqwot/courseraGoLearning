package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("No money")
	}
	w.Cash -= amount
	return nil
}

func main() {

}
