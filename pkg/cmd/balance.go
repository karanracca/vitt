package cmd

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
	"vitt/pkg/core"

	"github.com/urfave/cli/v2"
)

func Balance(ctx *cli.Context) error {

	transactions, err := core.ReadLedger(ctx.Path("file"))
	if err != nil {
		return fmt.Errorf("rading ledger file failed with error %w", err)
	}

	balMap := make(map[string]core.Amount)

	for _, transaction := range transactions {
		for _, a := range []core.Account{transaction.Payee, transaction.Payer} {
			if amt, ok := balMap[a.Name]; ok {
				tot := core.Amount(a.Amt) + amt
				balMap[a.Name] = tot
			} else {
				balMap[a.Name] = core.Amount(a.Amt)
			}
		}
	}

	for acc, amt := range balMap {
		spaceCount := 80 - 4 - utf8.RuneCountInString(acc) - utf8.RuneCountInString(amt.String())
		if spaceCount < 1 {
			spaceCount = 1
		}
		fmt.Fprintf(os.Stdout, "\t%s%s%s\n", acc, strings.Repeat(" ", spaceCount), amt)
	}

	return nil
}
