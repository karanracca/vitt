package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"vitt/pkg/core"

	"github.com/urfave/cli/v2"
)

func Import(ctx *cli.Context) error {
	// TODO: Validate account to be one of [Asset, Liability, Expense, Income, Equity]
	accountString := ctx.Args().First()
	if accountString == "" {
		return errors.New("invalid of missing account name")
	}
	negateFlag := ctx.Bool("negate")

	inputFiles := ctx.Args().Tail()

	// Import transactions
	transactionsToImport := make([]*core.Transaction, 0)

	for _, inputFile := range inputFiles {
		log.Printf("Reading from file %s", inputFile)

		// Open file
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		// Close the file at the end of the program
		defer file.Close()

		// Run the bayesian classifier
		categorizer, err := core.InitClassifier(ctx.String("file"))
		if err != nil {
			return fmt.Errorf("classifier failed with error %w", err)
		}

		// Read csv values using csv.Reader
		csvReader := csv.NewReader(file)

		// Read header row
		lines, err := csvReader.ReadAll()
		if err != nil {
			log.Printf("reading CSV file failed with error %v", err)
		}

		headerMap := extractColumnsFromHeader(lines[0])
		if headerMap["date"] < 0 || headerMap["description"] < 0 || headerMap["amount"] < 0 {
			return errors.New("unable to find columns required from header field names")
		}

		for index, line := range lines[1:] {

			var amount core.Amount
			var date core.Date
			var err error

			amount, err = core.ToAmount(line[headerMap["amount"]])
			if err != nil {
				log.Printf("skipping line %d. Parsing Amount failed with error %v", index, err)
				continue
			}
			if negateFlag {
				amount = amount.Negate()
			}

			date, err = core.ToDate(line[headerMap["date"]], ctx.String("datefmt"))
			if err != nil {
				log.Print(fmt.Errorf("skipping line %d. %w", index, err))
				continue
			}

			transactionsToImport = append(transactionsToImport, &core.Transaction{
				Date:        date,
				Description: line[headerMap["description"]],
				Payee: core.Account{
					Name: categorizer.Categorize(line[headerMap["description"]]),
					Amt:  amount.Negate(),
				},
				Payer: core.Account{
					Name: accountString,
					Amt:  amount,
				},
			})
		}
	}

	if ctx.Bool("dry-run") {
		WriteStdOut(transactionsToImport)
	} else {
		WriteLedgerFile(ctx.Path("file"), transactionsToImport)
	}

	return nil
}

func extractColumnsFromHeader(header []string) map[string]int {
	headerMap := map[string]int{
		"date":        -1,
		"description": -1,
		"amount":      -1,
	}

	for fieldIndex, fieldName := range header {
		fieldName = strings.ToLower(fieldName)
		if DATE_HEADERS.Exist(fieldName) {
			headerMap["date"] = fieldIndex
		} else if DESCRIPTION_HEADERS.Exist(fieldName) {
			headerMap["description"] = fieldIndex
		} else if AMOUNT_HEADERS.Exist(fieldName) {
			headerMap["amount"] = fieldIndex
		}
	}

	return headerMap
}

// TODO move to its own file
func Sort(ctx *cli.Context) error {
	return core.SortByDate(ctx.Path("file"))
}
