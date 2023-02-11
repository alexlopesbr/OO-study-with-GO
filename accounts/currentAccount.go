package accounts

import (
	"bank/customers"
)

type CurrentAccount struct {
	Titular       customers.Titular
	NumeroAgencia int
	NumeroConta   int
	balance       float64
}

func (c *CurrentAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "withdraw realized with success"
	} else {
		return "Insufficient funds"
	}
}

func (c *CurrentAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successfully", c.balance
	}
	return "Deposit failed", c.balance
}

func (c *CurrentAccount) Transfer(transferValue float64, destinationAccount *CurrentAccount) bool {
	if transferValue > 0 && transferValue < c.balance {
		c.balance -= transferValue

		destinationAccount.Deposit(transferValue)
		return true
	}
	return false
}

func (c *CurrentAccount) GetSold() float64 {
	return c.balance
}
