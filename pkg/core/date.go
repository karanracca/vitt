package core

import (
	"fmt"
	"time"
)

type Date time.Time

func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := t.Format("2006-01-02")
	return []byte(fmt.Sprintf(`"%s"`, formatted)), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	trimmed := string(data)[1 : len(data)-1] // Remove the surrounding quotes
	parsed, err := time.Parse("2006-01-02", trimmed)
	if err != nil {
		return err
	}
	*d = Date(parsed)
	return nil
}

var dateFormat = map[string]string{
	"dd/mm/yy":   "02/01/06",
	"dd/mm/yyyy": "02/01/2006",
	"mm/dd/yyyy": "01/02/2006", // Default Date format
	"utc": "2006-01-02 15:04:05 +0000 UTC",
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
