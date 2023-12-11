package bankcore

import (
	"errors"
	"fmt"
)

func Hello() string {
	return "Hey! I'm working!"
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be lower than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (origin *Account) Transfer(destination *Account, amount float64) error {
	err := origin.Withdraw(amount)
                if err != nil {
                    //fmt.Fprintf(w, "%v", err)
					return err
                } else {
                    //fmt.Fprintf(w, origin.Statement())
                    err := destination.Deposit(amount)
                    if err != nil {
                        //fmt.Fprintf(w, "%v", err)
                        origin.Deposit(amount)
						return err
                    } else {
                        //fmt.Fprintf(w, destination.Statement())
						return nil
                    }
                }
}
