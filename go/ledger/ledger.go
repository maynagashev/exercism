// Package ledger Description: This package provides a function to format a ledger for a bank account.
// Before:
// goos: darwin
// goarch: amd64
// pkg: ledger
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkFormatLedger-12           19683             57479 ns/op            8400 B/op        248 allocs/op
// PASS
package ledger

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Entry represents a single entry in a ledger
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func (e Entry) Less(f Entry) bool {
	if e.Date < f.Date {
		return true
	} else {
		if e.Description < f.Description {
			return true
		} else {
			return e.Change < f.Change
		}
	}
}

func EntryLess(e []Entry) func(int, int) bool {
	return func(i, j int) bool {
		return e[i].Less(e[j])
	}
}

// FormatMoney Adds the decimal cleanly.
func FormatMoney(value int, separator, decimal string) string {
	value = int(math.Abs(float64(value)))
	cents := value % 100
	value /= 100
	var dollarSplit string
	for value > 0 {
		dollarSplit = strconv.Itoa(value%1000) + dollarSplit
		value /= 1000
		if value > 0 {
			dollarSplit = separator + dollarSplit
		}
	}
	if len(dollarSplit) == 0 {
		dollarSplit = "0"
	}
	return fmt.Sprintf("%s%s%02d", dollarSplit, decimal, cents)
}

// FormatLedger formats a ledger for a bank account
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {

	// If there are no entries, we still need to print the header with default date
	if len(entries) == 0 {
		_, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01"}})
		if err != nil {
			return "", err
		}
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
	sort.SliceStable(entriesCopy, EntryLess(entriesCopy))

	var header string
	headerFormatString := "%-10s | %-25s | %s"

	// A few locale-specific formatting values
	var dateSeparator string
	var thousandSeparator string
	var decimal string
	var currencyFormat, negCurrencyFormat string
	switch locale {
	case "nl-NL":
		header = fmt.Sprintf(headerFormatString, "Datum", "Omschrijving", "Verandering")
		dateSeparator = "-"
		thousandSeparator = "."
		decimal = ","
		// Valid NL currency $ 12.345,67- = -$12345.67
		//                   € 12.345,67  =   12345.67€
		currencyFormat = "%s %s "
		negCurrencyFormat = "%s %s-"
	case "en-US":
		header = fmt.Sprintf(headerFormatString, "Date", "Description", "Change")
		dateSeparator = "/"
		thousandSeparator = ","
		decimal = "."
		// Valid US currency ($12,345.67) = -$12345.67
		//                    €12,345.67  =   12345.67€
		currencyFormat = "%s%s "
		negCurrencyFormat = "(%s%s)"
	default:
		return "", errors.New("Unsupported locale")
	}
	var currencySymbol string
	switch currency {
	case "EUR":
		currencySymbol = "€"
	case "USD":
		currencySymbol = "$"
	default:
		return "", errors.New("Unsupported currency: " + currency)
	}

	ledger := []string{header}
	dateFormat := regexp.MustCompile(`^(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})$`)
	for _, entry := range entriesCopy {
		dateSegments := dateFormat.FindStringSubmatch(entry.Date)
		if len(dateSegments) != 4 {
			return "", errors.New("Invalid date format")
		}
		var description string
		if len(entry.Description) > 25 {
			description = entry.Description[:22] + "..."
		} else {
			description = fmt.Sprintf("%-25s", entry.Description)
		}
		switch locale {
		case "nl-NL":
			dateSegments = []string{dateSegments[3], dateSegments[2], dateSegments[1]}
		case "en-US":
			dateSegments = []string{dateSegments[2], dateSegments[3], dateSegments[1]}
		}

		date := strings.Join(dateSegments, dateSeparator)
		cost := FormatMoney(entry.Change, thousandSeparator, decimal)
		var currencyDisplay string
		if entry.Change < 0 {
			currencyDisplay = fmt.Sprintf(negCurrencyFormat, currencySymbol, cost)
		} else {
			currencyDisplay = fmt.Sprintf(currencyFormat, currencySymbol, cost)
		}
		ledger = append(ledger, fmt.Sprintf("%-10s | %s | %13s", date, description, currencyDisplay))
	}
	completeLedger := strings.Join(ledger, "\n") + "\n"
	return completeLedger, nil
}
