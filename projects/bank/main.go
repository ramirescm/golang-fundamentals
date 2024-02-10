package main

import "github.com/ramirescm/golang-fundamentals/projects/bank/account"

func main() {
	account1 := account.CurrentAccount{
		AccountHolder: account.AccountHolder{Name: "Mad Max"},
		AccountNumber: 1234,
		AgencyNumber:  2222,
	}
	account2 := account.CurrentAccount{}

	account1.Deposit(1000)
	//account1.ShowBalance()
	account1.Withdraw(50)
	//account1.ShowBalance()
	account1.Deposit(150.00)
	account1.ShowBalance()
	account1.Transfer(100, &account2)
	account2.ShowBalance()
	account1.ShowBalance()
	account.Payment(&account1, 150)
	account1.ShowBalance()

	savingAccount := account.SavingAccount{}
	savingAccount.Deposit(100)
	account.Payment(&savingAccount, 25)
	savingAccount.ShowBalance()
}
