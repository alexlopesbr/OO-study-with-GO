package accounts

import "bank/customers"

type SavingAccount struct {
	Titular       customers.Titular
	AgencyNumber  int
	AccountNumber int
	Operation     int
	balance       float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "withdraw realized with success"
	} else {
		return "Insufficient funds"
	}
}

func (c *SavingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit successfully", c.balance
	}
	return "Deposit failed", c.balance
}

func (c *SavingAccount) GetSold() float64 {
	return c.balance
}
