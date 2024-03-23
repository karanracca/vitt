package core

import (
	"fmt"
	"io"
)

func Write(writer io.Writer, transactions []*Transaction) error {

	for _, transaction := range transactions {
		fmt.Fprintf(writer, "%s | %s | %s | %s | %s | %s\n", transaction.Id, transaction.Date, transaction.Description, transaction.Amt, transaction.Source, transaction.Dest)
		fmt.Fprintln(writer, "")
	}

	return nil

}
