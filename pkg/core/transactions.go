package core

import (
	"strconv"

	"github.com/mitchellh/hashstructure/v2"
)

type Direction string

const (
	In  Direction = "In"
	Out Direction = "Out"
)

// type Account struct {
// 	Id     string      `json:"id" hash:"ignore"`
// 	Name   string      `json:"name"`
// 	Type   AccountType `json:"category"`
// 	Parent string      `json:"parent"`
// 	//Total  Amount      `json:"amount" hash:"ignore"` // Amount is ignored in hash calculation as it changes with every persist
// }

// type Participant struct {
// 	Account     Account    `json:"acc"`
// 	SubAccounts []*Account `json:"sub_acc"`
// 	Amt         Amount     `json:"amount"`
// }

// func (part *Participant) ToParticipantRow() ParticipantRow {
// 	subAccountIds := make([]string, 0)
// 	for _, subAcc := range part.SubAccounts {
// 		subAccountIds = append(subAccountIds, subAcc.Id)
// 	}

// 	return ParticipantRow{
// 		AccountId:     part.Account.Id,
// 		SubAccountIds: subAccountIds,
// 		Amt:           part.Amt,
// 	}
// }

// type ParticipantRow struct {
// 	AccountId     string   `json:"acc_id"`
// 	SubAccountIds []string `json:"sub_acc_ids"`
// 	Amt           Amount   `json:"amount"`
// }

type Transaction struct {
	Id          string `json:"id" hash:"ignore"`
	Hash        string `json:"hash" hash:"ignore"`
	Date        Date   `json:"date"`
	Description string `json:"description"`
	Source      string
	Dest        string
	Direction   Direction
	AccType     AccountType `json:"accType"`
	Comments    string      `json:"comments" hash:"ignore"`
	Amt         Amount      `json:"amount"`
}

func (tran *Transaction) GenHash() error {
	hash, err := hashstructure.Hash(tran, hashstructure.FormatV2, nil)
	if err != nil {
		return err
	}

	// Using strconv.FormatUint to convert uint64 to string
	tran.Hash = strconv.FormatUint(hash, 10)
	return nil
}

type TransactionRow struct {
	Id          string
	Hash        string
	Date        string
	Description string
	Source      string
	Dest        string
	Direction   string
	AccType     string
	Comments    string
	Amt         float64
}
