// The package provides methods for converting amounts between currencies. The
// exchange rates are provided by the ECB (http://www.ecb.europa.eu/).
//
// Author: Michael Banzon
package currency

import (
	"errors"
	"fmt"
	"time"
)

const (
	eur     = "EUR"
	unknown = "Unknown currency: %s"
)

// The CurrencyConverter struct holds the data that enables the conversion.
// Upon creation the data is fetched from the ECB and parsed into the struct.
type CurrencyConverter struct {
	date       time.Time
	currencies map[string]float64
}

// The SingleCurrencyConverter struct holds data about how to convert amounts
// between two pre-defined currencies.
type SingleCurrencyConverter struct {
	date             time.Time
	from, to         string
	fromRate, toRate float64
}

// Creates a new converter by fetching the required data from the ECB.
func NewConverter() (*CurrencyConverter, error) {
	currencyTime, currencies, err := parseEcbData()
	if err != nil {
		return nil, err
	}
	converter := CurrencyConverter{date: currencyTime, currencies: currencies}
	return &converter, nil
}

// Calculates the age in days of the CurrencyConverter. The age is calculated
// using the date supplied in the currency feed from the ECB.
func (c *CurrencyConverter) Age() float64 {
	delta := c.date.Sub(time.Now())
	return delta.Hours() / 24
}

// Returns true if the currencies stores are so old they should be renewed from
// the ECB server.
func (c *CurrencyConverter) ShouldRenew() bool {
	if c.Age() >= 1 {
		today := time.Now()
		if today.Weekday() > time.Sunday && today.Weekday() < time.Saturday {
			return true
		}
	}
	return false
}

// Renew the currency data by fetching the  from the ECB server.
func (c *CurrencyConverter) Renew() error {
	date, currencies, err := parseEcbData()
	if err != nil {
		return err
	} else {
		c.date = date
		c.currencies = currencies
		return nil
	}
}

// Converts an amount between two currencies.
func (c *CurrencyConverter) Convert(amount float64, from string, to string) (float64, error) {
	fromRate, fromOk := c.currencies[from]
	if !fromOk {
		return 0, errors.New(fmt.Sprintf(unknown, from))
	}

	toRate, toOk := c.currencies[to]
	if !toOk {
		return 0, errors.New(fmt.Sprintf(unknown, to))
	}

	return amount / fromRate * toRate, nil
}

// Converts a slice of amounts from one currency to another.
func (c *CurrencyConverter) MultiConvert(amounts []float64, from, to string) ([]float64, error) {
	convertedAmounts := make([]float64, len(amounts))
	var e error
	for i, amount := range amounts {
		converted, err := c.Convert(amount, from, to)
		if err != nil {
			e = err
		}
		convertedAmounts[i] = converted
	}
	return convertedAmounts, e
}

// Creates a SingleCurrencyConverter that easilly translates amounts between
// two fixed currencies.
func (c *CurrencyConverter) GetSingleCurrencyConverter(from, to string) (*SingleCurrencyConverter, error) {
	fromRate, fromOk := c.currencies[from]
	if !fromOk {
		return nil, errors.New(fmt.Sprintf(unknown, from))
	}
	toRate, toOk := c.currencies[to]
	if !toOk {
		return nil, errors.New(fmt.Sprintf(unknown, to))
	}

	converter := SingleCurrencyConverter{date: c.date, from: from, to: to, fromRate: fromRate, toRate: toRate}
	return &converter, nil
}

// Converts a single amount.
func (c *SingleCurrencyConverter) Convert(amount float64) float64 {
	return amount / c.fromRate * c.toRate
}

// Converts multiple amounts.
func (c *SingleCurrencyConverter) MultiConvert(amounts []float64) []float64 {
	converted := make([]float64, len(amounts))
	for i, amount := range amounts {
		converted[i] = c.Convert(amount)
	}
	return converted
}
