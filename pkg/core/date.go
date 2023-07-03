package core

import (
	"time"
)

type Date time.Time

var dateFormat = map[string]string{
	"dd/mm/yy":   "02/01/06",
	"dd/mm/yyyy": "02/01/2006",
	"mm/dd/yyyy": "01/02/2006", // Default Date format
}

func ToDate(raw string, format string) (Date, error) {
	if format != "" {
		date, err := time.Parse(dateFormat[format], raw)
		if err != nil {
			return Date(date), err
		}
		return Date(date), nil
	} else {
		date, err := time.Parse(dateFormat["mm/dd/yyyy"], raw)
		if err != nil {
			return Date(date), err
		}
		return Date(date), nil
	}
}

func (d Date) String() string {
	return time.Time(d).Format("01/02/2006")
}
