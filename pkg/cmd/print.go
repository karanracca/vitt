package cmd

import (
	"os"
	"vitt/pkg/core"
	"vitt/pkg/store"

	"github.com/urfave/cli/v2"
)

func PrintTransactions(ctx *cli.Context) error {
	db := ctx.Context.Value("db").(*store.Store)

	trans, err := db.GetTransactions()
	if err != nil {
		return err
	}

	return WriteStdOut(trans)

}

func WriteStdOut(transactions []*core.Transaction) error {
	return core.Write(os.Stdout, transactions)
}
