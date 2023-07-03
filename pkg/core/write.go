package core

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

func Write(writer io.Writer, transactions []*Transaction) error {

	for _, transaction := range transactions {
		fmt.Fprintf(writer, "%s  %s\n", transaction.Date, transaction.Description)

		spaceCount2 := 80 - 4 - utf8.RuneCountInString(transaction.Payee.Name) - utf8.RuneCountInString(transaction.Payee.Amt.String())
		if spaceCount2 < 1 {
			spaceCount2 = 1
		}
		fmt.Fprintf(writer, "\t%s%s%s\n", transaction.Payee.Name, strings.Repeat(" ", spaceCount2), transaction.Payee.Amt.String())

		spaceCount := 80 - 4 - utf8.RuneCountInString(transaction.Payer.Name) - utf8.RuneCountInString(transaction.Payer.Amt.String())
		if spaceCount < 1 {
			spaceCount = 1
		}
		fmt.Fprintf(writer, "\t%s%s%s\n", transaction.Payer.Name, strings.Repeat(" ", spaceCount), transaction.Payer.Amt.String())

		fmt.Fprintln(writer, "")
	}

	return nil

}
