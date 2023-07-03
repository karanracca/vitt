package cmd

import (
	"os"
	"vitt/pkg/core"

	"github.com/urfave/cli/v2"
)

func WriteLedgerFile(ledgerFilePath string, transactions []*core.Transaction) error {
	file, err := os.OpenFile(ledgerFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	return core.Write(file, transactions)
}

func WriteStdOut(transactions []*core.Transaction) error {
	return core.Write(os.Stdout, transactions)
}

func Print(ctx *cli.Context) error {
	transactions, err := core.ReadLedger(ctx.Path("file"))
	if err != nil {
		return err
	}

	return WriteStdOut(transactions)
}
