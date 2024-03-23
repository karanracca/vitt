package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"vitt/pkg/core"

	"github.com/lithammer/shortuuid"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

type Store struct {
	DB *sql.DB
}

func Init(path string) (*Store, error) {
	db, err := sql.Open("sqlite", "./vitt.db")
	if err != nil {
		return &Store{}, fmt.Errorf("opening a connection to DB failed %w", err)
	}

	err = runMigrations(db)
	if err != nil {
		return &Store{}, fmt.Errorf("running migrations failed %w", err)
	}

	log.Println("migrations completed successfully")
	return &Store{DB: db}, nil
}

func runMigrations(db *sql.DB) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return fmt.Errorf("error creating SQLite driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite",
		driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	// Migrate to the latest version
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error applying migrations: %w", err)
	}

	return nil
}

func (store *Store) toTransaction(tranRow core.TransactionRow) (*core.Transaction, error) {
	date, err := core.ToDate(tranRow.Date, "utc")
	if err != nil {
		return nil, err
	}

	tran := core.Transaction{
		Id:          tranRow.Id,
		Hash:        tranRow.Hash,
		Date:        date,
		Description: tranRow.Description,
		Source:      tranRow.Source,
		Dest:        tranRow.Dest,
		Direction:   core.Direction(tranRow.Direction),
		AccType:     core.AccountType(tranRow.AccType),
		Amt:         core.Amount(tranRow.Amt),
	}

	return &tran, nil
}

func (store *Store) FindTransactionByHash(hash string) (*core.Transaction, error) {
	var tranRow core.TransactionRow
	err := store.DB.QueryRow("SELECT * FROM transactions WHERE hash=?", hash).Scan(&tranRow.Id, &tranRow.Hash, &tranRow.Date, &tranRow.Description, &tranRow.Source, &tranRow.Dest, &tranRow.Direction, &tranRow.AccType, &tranRow.Amt, &tranRow.Comments)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
	}

	tran, err := store.toTransaction(tranRow)
	if err != nil {
		return nil, err
	}

	return tran, nil
}

func (store *Store) CreateTransactions(transactions []*core.Transaction) error {
	for _, tran := range transactions {
		id := fmt.Sprintf("trx_%s", shortuuid.New())

		if _, err := store.DB.Exec(`
			INSERT INTO transactions (id, hash, date, description, acc_type, source, dest, comments, amount, direction) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, id, tran.Hash, time.Time(tran.Date), tran.Description, tran.AccType, tran.Source, tran.Dest, tran.Comments, tran.Amt, tran.Direction); err != nil {
			return err
		}
	}

	return nil
}

func (store *Store) GetTransactions() ([]*core.Transaction, error) {
	var trans []*core.Transaction
	rows, err := store.DB.Query("SELECT * FROM transactions")
	if err != nil {
		return trans, err
	}
	defer rows.Close()

	for rows.Next() {
		var tranRow core.TransactionRow
		if err := rows.Scan(&tranRow.Id, &tranRow.Hash, &tranRow.Date, &tranRow.Description, &tranRow.Source, &tranRow.Dest, &tranRow.Direction, &tranRow.AccType, &tranRow.Amt, &tranRow.Comments); err != nil {
			return trans, fmt.Errorf("failed to fetch row %w", err)
		}
		tran, err := store.toTransaction(tranRow)
		if err != nil {
			return trans, fmt.Errorf("failed to fetch row %w", err)
		}
		trans = append(trans, tran)
	}

	return trans, nil
}

// func (store *Store) GetAccountByName(name string) (core.Account, error) {
// 	rows, err := store.DB.Query("SELECT * FROM accounts WHERE name = ?", name)
// 	if err != nil {
// 		return core.Account{}, err
// 	}
// 	defer rows.Close()

// 	var accounts []core.Account
// 	for rows.Next() {
// 		var account core.Account
// 		if err := rows.Scan(&account.Id, &account.Name, &account.Type, &account.Parent); err != nil {
// 			return core.Account{}, err
// 		}
// 		accounts = append(accounts, account)
// 	}
// 	// Check for errors from iterating over rows.
// 	if err := rows.Err(); err != nil {
// 		return core.Account{}, err
// 	}

// 	if len(accounts) == 0 {
// 		return core.Account{}, nil
// 	}

// 	return accounts[0], nil
// }

// func (store *Store) GetAccountById(id string) (core.Account, error) {
// 	rows, err := store.DB.Query("SELECT * FROM accounts WHERE id=?", id)
// 	if err != nil {
// 		return core.Account{}, err
// 	}
// 	defer rows.Close()

// 	var accounts []core.Account
// 	for rows.Next() {
// 		var account core.Account
// 		if err := rows.Scan(&account.Id, &account.Name, &account.Type, &account.Parent); err != nil {
// 			return core.Account{}, err
// 		}
// 		accounts = append(accounts, account)
// 	}
// 	// Check for errors from iterating over rows.
// 	if err := rows.Err(); err != nil {
// 		return core.Account{}, err
// 	}

// 	if len(accounts) == 0 {
// 		return core.Account{}, fmt.Errorf("account not found")
// 	}

// 	return accounts[0], nil
// }

// func (store *Store) GetOrCreateAccountByName(name string, accType string, parentId string) (core.Account, error) {
// 	var acc core.Account
// 	var err error

// 	acc, err = store.GetAccountByName(name)
// 	if err != nil {
// 		return acc, fmt.Errorf("unable to fetch account %s %w", name, err)
// 	}

// 	// Check for no account found
// 	if acc == (core.Account{}) {
// 		// Create new account
// 		acc, err = store.CreateAccount(name, string(accType), "NULL")
// 		if err != nil {
// 			return acc, fmt.Errorf("unable to create account %s %w", name, err)
// 		}
// 		log.Printf("new account %s created", name)
// 	} else {
// 		log.Printf("account %s found with id %s", name, acc.Id)
// 	}

// 	return acc, nil
// }

// func (store *Store) CreateAccount(name string, accType string, parentId string) (core.Account, error) {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		return core.Account{}, err
// 	}

// 	if _, err := store.DB.Exec("INSERT INTO accounts (id, name, acc_type, parent, total) VALUES (?, ?, ?, ?, ?)", id, name, accType, parentId, 0.0); err != nil {
// 		return core.Account{}, err
// 	}

// 	return store.GetAccountById(id.String())
// }

// func (store *Store) UpdateAccountTotal(accId string, total core.Amount) error {
// 	if _, err := store.DB.Exec("UPDATE accounts SET total = ? WHERE id = ?", total, accId); err != nil {
// 		return err
// 	}

// 	return nil

// }
