package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Init() {

	app := &cli.App{
		Name:  "vitt",
		Usage: "A cli tool to manage your finances",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "file",
				Usage:    "Path to `LEDGER_FILE`",
				Required: true,
				Aliases:  []string{"f"},
				EnvVars:  []string{"LEDGER_FILE"},
			},
		},
		Action: Print,
		Commands: []*cli.Command{
			{
				Name:   "bal",
				Usage:  "Show accounts and their balances",
				Action: Balance,
			},
			{
				Name:   "import",
				Usage:  "Import transactions from CSV files",
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
				Name:   "sort",
				Usage:  "Sort transactions by date",
				Action: Sort,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
