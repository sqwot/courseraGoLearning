package main

import "fmt"

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("No money, no honey")
	}
	w.Cash -= amount
	return nil
}

type Card struct {
	Ballance    int
	VallidUntil string
	CardHolder  string
	CVV         string
	Number      string
}

func (c *Card) Pay(amount int) error {
	if c.Ballance < amount {
		return fmt.Errorf("No money on card")
	}
	c.Ballance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("No money on account")
	}
	a.Money -= amount
	return nil
}

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(100)
	if err != nil {
		fmt.Printf("Error on paing %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Tx for buing through %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Ballance: 100, CardHolder: "sqwot"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
