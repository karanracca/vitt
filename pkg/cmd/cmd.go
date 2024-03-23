package cmd

import (
	"context"
	"log"
	"os"
	"vitt/pkg/store"

	"github.com/urfave/cli/v2"
)

func Init(store *store.Store) {

	app := &cli.App{
		Name:  "vitt",
		Usage: "A tool to manage your finances",
		Commands: []*cli.Command{
			{
				Name:  "import",
				Usage: "Import transactions from CSV files",
				Before: func(cCtx *cli.Context) error {
					cCtx.Context = context.WithValue(cCtx.Context, "db", store)
					return nil
				},
				Action: Import,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "negate",
						Aliases: []string{"n"},
						Value:   false,
						Usage:   "Negate the amount value",
					},
					&cli.StringFlag{
						Name:  "datefmt",
						Usage: "Date format used in the CSV",
					},
					&cli.BoolFlag{
						Name:  "dry-run",
						Value: false,
						Usage: "Prints to the stdout instead of writing to the ledger file",
					},
				},
			},
			{
				Name:  "print",
				Usage: "Print all transactions",
				Before: func(cCtx *cli.Context) error {
					cCtx.Context = context.WithValue(cCtx.Context, "db", store)
					return nil
				},
				Action: PrintTransactions,
			},
			// {
			// 	Name:  "web",
			// 	Usage: "Start the web server",
			// 	Before: func(cCtx *cli.Context) error {
			// 		cCtx.Context = context.WithValue(cCtx.Context, "db", store)
			// 		return nil
			// 	},
			// 	Action: api.Init,
			// },
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
