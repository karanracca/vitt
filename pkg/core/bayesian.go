package core

import (
	"strings"

	"github.com/navossoc/bayesian"
)

type Categorizer struct {
	classifier *bayesian.Classifier
}

func (cat *Categorizer) Categorize(desc string) string {
	_, likely, _ := cat.classifier.LogScores(strings.Fields(desc))
	return string(cat.classifier.Classes[likely])
}

func InitClassifier(filePath string) (Categorizer, error) {
	// Get all exisiting transactions
	transactions, err := ReadLedger(filePath)
	if err != nil {
		return Categorizer{}, err
	}

	// Create classes
	uniqueAccounts := make(map[string]interface{})
	for _, tran := range transactions {
		if _, ok := uniqueAccounts[tran.Payee.Name]; !ok {
			uniqueAccounts[tran.Payee.Name] = nil
		}
	}

	classes := make([]bayesian.Class, 0)
	for acc := range uniqueAccounts {
		classes = append(classes, bayesian.Class(acc))
	}

	classifier := bayesian.NewClassifier(classes...)

	// Learn
	for _, tran := range transactions {
		classifier.Learn(strings.Fields(tran.Description), bayesian.Class(tran.Payee.Name))
	}

	return Categorizer{classifier}, nil
}
