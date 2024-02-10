package account

import (
	"fmt"
	"strconv"
)

type SavingAccount struct {
	AccountHolder AccountHolder
	AccountNumber int
	AgencyNumber  int
	Operation     int
	balance       float64
}

func (c *SavingAccount) Withdraw(amount float64) {
	if amount <= 0 {
		c.ShowMessage("Withdraw must be greather than 0")
		return
	}

	avaiablebalance := amount <= c.balance
	if avaiablebalance {
		c.balance -= amount

		c.ShowMessage("Info:\n Withdraw request amount: " + strconv.FormatFloat(amount, 'f', 2, 64) + "\n Withdraw successfully completed")
		return
	}

	c.ShowMessage("Error: Unavailable balance")
}

func (c *SavingAccount) Deposit(amount float64) {
	if amount <= 0 {
		c.ShowMessage("Error: Deposit must be greather than 0")
		return
	}

	c.balance += amount
	c.ShowMessage("Info:\n Deposit request amount: " + strconv.FormatFloat(amount, 'f', 2, 64) + "\n Deposit successfully completed!")
}

func (c *SavingAccount) ShowBalance() {
	fmt.Printf("Account Holder: "+c.AccountHolder.Name+" - balance %.2f\n", c.balance)
}

func (c *SavingAccount) ShowMessage(message string) {
	fmt.Println(message + "\n")
}
