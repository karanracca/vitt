package core

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func SortByDate(ledgerFilePath string) error {
	transactions, err := ReadLedger(ledgerFilePath)
	if err != nil {
		return fmt.Errorf("reading ledger file failed with %w", err)
	}

	// sort transactions
	sort.Slice(transactions, func(i, j int) bool {
		return time.Time(transactions[i].Date).Before(time.Time(transactions[j].Date))
	})

	// Rewrite ledger file with sorted transactions
	// Write to a temp file

	file, err := os.OpenFile("temp.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("writing to temp file failed with %w", err)
	}
	defer file.Close()

	if err := Write(file, transactions); err != nil {
		return err
	}

	// Delete current file
	if err := os.Remove(ledgerFilePath); err != nil {
		return fmt.Errorf("unable to remove the ledger file %w", err)
	}
	// Rename temp to dest name
	if err := os.Rename("temp.dat", ledgerFilePath); err != nil {
		return fmt.Errorf("unable to rename the ledger file %w", err)
	}

	return nil

}
