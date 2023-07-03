package core

type Account struct {
	Name string
	Amt  Amount
}

type Transaction struct {
	Date        Date
	Description string
	Payee       Account
	Payer       Account
}
