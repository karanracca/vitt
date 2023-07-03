package core

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func StringToDate(s string) (time.Time, error) {
	var date time.Time
	var err error
	day := strings.Split(s, "/")[1]
	month := strings.Split(s, "/")[0]
	year := strings.Split(s, "/")[2]

	var stdDay string
	if len(day) == 2 {
		stdDay = "02"
	} else if len(month) == 1 {
		stdDay = "2"
	} else {
		err = errors.New(fmt.Sprintf("Invalid Date Format for Date: %s", s))
	}

	var stdMonth string
	if len(month) == 2 {
		stdMonth = "01"
	} else if len(month) == 1 {
		stdMonth = "1"
	} else {
		err = errors.New(fmt.Sprintf("Invalid Date Format for Date: %s", s))
	}

	var stdYear string
	if len(year) == 4 {
		stdYear = "2006"
	} else if len(year) == 2 {
		stdYear = "06"
	} else {
		err = errors.New(fmt.Sprintf("Invalid Date Format for Date: %s", s))
	}
	layout := fmt.Sprintf("%s/%s/%s", stdMonth, stdDay, stdYear)
	date, err = time.Parse(layout, s)
	if err != nil {
		return date, err
	}
	return date, nil
}

func StringToAmount(s string) (float64, error) {
	var amount float64
	var err error
	amount, err = strconv.ParseFloat(s, 32)
	// Convert to double precision
	amount = float64(int(amount*100)) / 100
	if err != nil {
		return amount, err
	}
	return amount, nil
}
