package cmd

import (
	"strings"
)

type ColumnHeader []string

func (columnHeader ColumnHeader) Exist(str string) bool {
	for _, v := range columnHeader {
		str = strings.ToLower(str)
		if v == str {
			return true
		}
	}
	return false
}

var DATE_HEADERS ColumnHeader = []string{"date", "posted date", "transaction date"}
var DESCRIPTION_HEADERS ColumnHeader = []string{"description", "payee"}
var AMOUNT_HEADERS ColumnHeader = []string{"amount"}
var ACCOUNT_HEADERS ColumnHeader = []string{"account"}
