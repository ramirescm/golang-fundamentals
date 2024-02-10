package account

import (
	"fmt"
	"strconv"
)

type CurrentAccount struct {
	AccountHolder AccountHolder
	AccountNumber int
	AgencyNumber  int
	balance       float64
}

func (c *CurrentAccount) Withdraw(amount float64) {
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

func (c *CurrentAccount) Deposit(amount float64) {
	if amount <= 0 {
		c.ShowMessage("Error: Deposit must be greather than 0")
		return
	}

	c.balance += amount
	c.ShowMessage("Info:\n Deposit request amount: " + strconv.FormatFloat(amount, 'f', 2, 64) + "\n Deposit successfully completed!")
}

func (c *CurrentAccount) Transfer(amount float64, receiver *CurrentAccount) bool {
	if amount > 0 && amount <= c.balance {
		c.balance -= amount
		receiver.Deposit(amount)
		return true
	}

	return false
}

func (c *CurrentAccount) ShowBalance() {
	fmt.Printf("Account Holder: "+c.AccountHolder.Name+" - balance %.2f\n", c.balance)
}

func (c *CurrentAccount) ShowMessage(message string) {
	fmt.Println(message + "\n")
}
