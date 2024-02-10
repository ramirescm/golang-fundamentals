package account

type Account interface {
	Withdraw(amount float64) // current account and saving account must have this method.. respect this contract
}

func Payment(account Account, amount float64) { // generic method that allow shared logic between differents accounts
	account.Withdraw(amount)
}
