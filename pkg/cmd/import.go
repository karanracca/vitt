package cmd

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"vitt/pkg/core"
	"vitt/pkg/store"

	"github.com/urfave/cli/v2"
)

func Import(ctx *cli.Context) error {

	db := ctx.Context.Value("db").(*store.Store)

	acctype := ctx.Args().First()
	if acctype == "" {
		return errors.New("invalid account. Account should follow format <type>:<acc_name>")
	}
	accSlice := strings.Split(acctype, ":")
	if len(accSlice) < 1 {
		return errors.New("invalid account. Account should follow format <type>:<acc_name>")
	}

	// Validate account to be one of [Asset, Liability, Expense, Income, Equity]
	accType, err := core.ToAccountCategory(accSlice[0])
	if err != nil {
		return err
	}

	// Gather payer accounts
	payerAccName := accSlice[1]
	//Check if account is present else create it
	// payerAccount, err := db.GetOrCreateAccountByName(accountName, string(accType), "NUll")
	// if err != nil {
	// 	return err
	// }

	negateFlag := ctx.Bool("negate")

	inputFiles := ctx.Args().Tail()

	// Import transactions
	transToImport := make([]*core.Transaction, 0)

	for _, inputFile := range inputFiles {
		log.Printf("Reading from file %s", inputFile)

		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		csvReader := csv.NewReader(file)
		lines, err := csvReader.ReadAll()
		if err != nil {
			log.Printf("reading CSV file failed with error %v", err)
		}

		// Read header row
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

			var direction core.Direction
			if amount < 0 {
				direction = "Out"
			} else {
				direction = "In"
			}

			date, err = core.ToDate(line[headerMap["date"]], ctx.String("datefmt"))
			if err != nil {
				log.Print(fmt.Errorf("skipping line %d. %w", index, err))
				continue
			}

			tran := &core.Transaction{
				Id:          "",
				Date:        date,
				Description: line[headerMap["description"]],
				Hash:        "",
				AccType:     accType,
				Source:      payerAccName,
				Direction:   direction,
				Comments:    "",
				Amt:         amount,
			}

			tran.GenHash()

			//Check if the same transaction already exisits
			exisits, err := db.FindTransactionByHash(tran.Hash)
			if err != nil {
				return err
			} else if exisits != nil {
				log.Printf("skipping transaction. Found duplicate transaction with id %s", exisits.Id)
			} else {
				transToImport = append(transToImport, tran)
			}
		}
	}

	if ctx.Bool("dry-run") {
		WriteStdOut(transToImport)
	} else {
		// Write to the database
		if err := db.CreateTransactions(transToImport); err != nil {
			return err
		}
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
