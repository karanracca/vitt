package core

import "fmt"

type AccountType string

const (
	Asset     AccountType = "Asset"
	Liability AccountType = "Liability"
	Equity    AccountType = "Equity"
	Revenue   AccountType = "Revenue"
	Expense   AccountType = "Expense"
	Unknown   AccountType = "Unknown"
)

func ToAccountCategory(category string) (AccountType, error) {
	// Iterate through each AccountCategory constant and check for a match
	for _, c := range []AccountType{Asset, Liability, Equity, Revenue, Expense} {
		if string(c) == category {
			return c, nil
		}
	}
	// Return an empty AccountCategory if the input is not valid
	return AccountType(""), fmt.Errorf("invalid account category: %s. Should be one of Asset, Liability, Equity, Revenue or Expense", category)
}
