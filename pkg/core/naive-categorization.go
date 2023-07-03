// ------------ Depracated ------------
package core

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Category string

const (
	ONLINE_SERVICES  Category = "ONLINE_SERVICES"
	MISC             Category = "MISC"
	BANK_TRANSACTION Category = "BANK_TRANSACTION"
	FOOD_AND_DRINK   Category = "FOOD_AND_DRINK"
	PAYMENT          Category = "PAYMENT"
	RENT             Category = "RENT"
	RECREATION       Category = "RECREATION"
	SHOPPING         Category = "SHOPPING"
	SALARY           Category = "SALARY"
	GROCERY          Category = "GROCERY"
)

var SALARY_WORD_LIST = []string{
	"patientping",
	"payroll",
}

var ONLINE_SERVICES_WORD_LIST = []string{
	"amazon",
	"prime",
	"fee",
	"services",
	"netflix",
	"netflix.com",
	"spotify",
	"spotifyusai",
	"digitalocea",
	"google",
	"etsy.com",
}

var SHOPPING_WORD_LIST = []string{
	"amazon",
	"retail",
	"marketplace",
	"nike",
}

var FOOD_AND_DRINK_WORD_LIST = []string{
	"cafe",
	"dining",
	"bar",
	"cheesecake",
	"restaurants",
	"restaurant",
	"grill",
	"ihop",
	"pizza",
	"spice",
	"biryanis",
	"uber",
	"eats",
}

var GROCERY_WORD_LIST = []string{
	"convenience",
	"wholefds",
	"wholefoods",
	"trader",
	"joe's",
	"grocery",
	"7-eleven",
}

var TRAVEL_AND_FUEL_WORD_LIST = []string{
	"speedway",
	"uber",
}

var CategoryMap = map[Category][]string{
	"SALARY":          SALARY_WORD_LIST,
	"GROCERY":         GROCERY_WORD_LIST,
	"FOOD_AND_DRINK":  FOOD_AND_DRINK_WORD_LIST,
	"SHOPPING":        SHOPPING_WORD_LIST,
	"ONLINE_SERVICES": ONLINE_SERVICES_WORD_LIST,
}

type Categorization interface {
	GetCategory(desc string) Category
}

func contains(list *[]string, stringToFind string) bool {
	for _, str := range *list {
		if strings.ToLower(str) == strings.ToLower(stringToFind) {
			return true
		}
	}
	return false
}

func getCategoryByHighestCount(probabilityMap *map[Category]int16) Category {
	var highest int16
	var category Category = MISC
	for cat, count := range *probabilityMap {
		if count > highest {
			highest = count
			category = cat
		} else if count == highest {
			category = MISC
		}
	}
	return category
}

func cleanString(str string) string {
	specialCharRegex := regexp.MustCompile(`\*`)
	specialCharRegex.ReplaceAllString(str, "")
	return str
}

func GetCategory(desc string) string {
	cleanedDesc := cleanString(desc)
	wordRegexp := regexp.MustCompile(`(?i)([a-z]{4,})+`)
	wordList := wordRegexp.FindAllString(cleanedDesc, -1)
	probabilityMap := make(map[Category]int16)
	for _, word := range wordList {
		for category, categoryWordList := range CategoryMap {
			if contains(&categoryWordList, word) {
				count, ok := probabilityMap[category]
				if ok {
					probabilityMap[category] = count + 1
				} else {
					probabilityMap[category] = 1
				}
			}
		}
	}
	log.Printf("Probability map for description %s: %+v", desc, probabilityMap)
	return fmt.Sprintf("%s", getCategoryByHighestCount(&probabilityMap))
}
