package utils

import (
	"errors"
	"fmt"
)

type Account struct {
	Balance float64
}

func (a Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Println("存款:", amount)
}

func (a Account) Withdtaw(amount float64) error {
	if amount > a.Balance {
		return errors.New("余额不足")
	}
	a.Balance -= amount
	fmt.Println("取款：", amount)
	return nil
}