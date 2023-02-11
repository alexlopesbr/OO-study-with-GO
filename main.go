package main

import (
	"bank/accounts"
	"bank/customers"
	"fmt"
)

func main() {
	accountAle := accounts.CurrentAccount{
		Titular: customers.Titular{
			Name:       "Alexandre",
			Cpf:        "123.123.123-12",
			Occupation: "Backend Developer",
		},
		NumeroAgencia: 123,
		NumeroConta:   321,
	}
	savingAccountAle := accounts.SavingAccount{
		Titular: customers.Titular{
			Name:       "Alexandre",
			Cpf:        "123.123.123-12",
			Occupation: "Backend Developer",
		},
		AgencyNumber:  101,
		AccountNumber: 34534,
		Operation:     12,
	}

	// Another way
	customerMaria := customers.Titular{
		Name:       "Maria",
		Cpf:        "00000",
		Occupation: "Front-end Developer",
	}
	accountMaria := accounts.CurrentAccount{
		Titular:       customerMaria,
		NumeroAgencia: 123,
		NumeroConta:   000,
	}

	fmt.Println(accountAle)
	fmt.Println(savingAccountAle)
	fmt.Println(accountMaria)

	accountAle.Deposit(1000)
	fmt.Println(savingAccountAle.GetSold())
	PayTicket(&savingAccountAle, 50)
	fmt.Println(savingAccountAle.GetSold())
}

func PayTicket(account verifyAccount, ticketValue float64) {
	account.Withdraw(ticketValue)
}

type verifyAccount interface {
	Withdraw(value float64) string
}
