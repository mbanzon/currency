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

type CurrencyConverter struct {
	date       time.Time
	currencies map[string]float64
}

type SingleCurrencyConverter struct {
	date             time.Time
	from, to         string
	fromRate, toRate float64
}

func NewConverter() (*CurrencyConverter, error) {
	currencyTime, currencies, err := parseEcbData()
	if err != nil {
		return nil, err
	}
	converter := CurrencyConverter{date: currencyTime, currencies: currencies}
	return &converter, nil
}

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

func (c *SingleCurrencyConverter) Convert(amount float64) float64 {
	return amount / c.fromRate * c.toRate
}
