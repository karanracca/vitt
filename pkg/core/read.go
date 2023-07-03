package core

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var spaceSplitter = regexp.MustCompile(`\s+`)

func ReadLedger(ledgerFilePath string) ([]*Transaction, error) {

	transactions := make([]*Transaction, 0)

	file, err := os.Open(ledgerFilePath)
	if err != nil {
		return transactions, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// remove heading and tailing space from the line
		trimmedLine := strings.TrimSpace(line)
		// Skip empty lines
		if len(trimmedLine) == 0 {
			continue
		}

		splited := spaceSplitter.Split(trimmedLine, 2)

		directive := strings.TrimSpace(splited[0])
		switch directive {
		// parse Transaction
		default:
			transaction, err := parseTransaction(scanner, trimmedLine, transactions)
			if err != nil {
				return transactions, err
			}

			transactions = append(transactions, &transaction)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return transactions, nil

}

func parseTransaction(scanner *bufio.Scanner, curLine string, transactions []*Transaction) (Transaction, error) {
	splited := spaceSplitter.Split(curLine, 2)
	// Parse Date
	date, err := ToDate(strings.TrimSpace(splited[0]), "")
	if err != nil {
		log.Printf("reading ledger file failed with error %v", err)
		return Transaction{}, err
	}

	// Parse Description
	discription := strings.TrimSpace(splited[1])

	// Prase Payee Account line
	scanner.Scan()
	payeeLine := scanner.Text()
	trimmedPayeeLine := strings.TrimSpace(payeeLine)
	if len(trimmedPayeeLine) == 0 {
		return Transaction{}, errors.New("reading ledger file failed, empty payee line")
	}
	splitedPayee := spaceSplitter.Split(trimmedPayeeLine, 2)
	payeeAccount := strings.TrimSpace(splitedPayee[0])
	payeeAmount, err := ToAmount(strings.TrimSpace(splitedPayee[1]))
	if err != nil {
		log.Printf("unable to parse amount for: %v", splitedPayee[1])
		return Transaction{}, err
	}

	// Prase Payer Account line
	scanner.Scan()
	payerLine := scanner.Text()
	trimmedPayerLine := strings.TrimSpace(payerLine)
	if len(trimmedPayerLine) == 0 {
		return Transaction{}, errors.New("reading ledger file failed, empty payer line")
	}
	splitedPayer := spaceSplitter.Split(trimmedPayerLine, 2)
	payerAccount := strings.TrimSpace(splitedPayer[0])
	payerAmount, err := ToAmount(strings.TrimSpace(splitedPayer[1]))
	if err != nil {
		log.Printf("unable to parse amount for: %v", splitedPayer[1])
		return Transaction{}, err
	}

	t := Transaction{
		Date:        date,
		Description: discription,
		Payer:       Account{Name: payerAccount, Amt: payerAmount},
		Payee:       Account{Name: payeeAccount, Amt: payeeAmount},
	}

	return t, nil
}
