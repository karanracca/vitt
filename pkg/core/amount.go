package core

import (
	"fmt"
	"regexp"
	"strconv"
)

type Amount float64

func ToAmount(raw string) (Amount, error) {
	normalizeRegex := regexp.MustCompile(`,`)
	normalizedAmt := normalizeRegex.ReplaceAllString(raw, "")

	var amt float64
	var err error
	amt, err = strconv.ParseFloat(normalizedAmt, 32)
	// Convert to double precision
	amt = float64(int(amt*100)) / 100
	if err != nil {
		return 0, err
	}
	return Amount(amt), nil
}

func (amount Amount) String() string {
	return fmt.Sprintf("%.2f", float64(amount))
}

func (amount Amount) Negate() Amount {
	return Amount(float64(amount) * -1)
}