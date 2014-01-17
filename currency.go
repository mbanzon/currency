// The package provides methods for converting amounts between currencies. The
// exchange rates are provided by the ECB (http://www.ecb.europa.eu/).
//
// Author: Michael Banzon
package currency

import (
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

func NewConverter() (*CurrencyConverter, error) {
	currencyTime, currencies, err := parseEcbData()
	if err != nil {
		return nil, err
	}
	converter := CurrencyConverter{date: currencyTime, currencies: currencies}
	return &converter, nil
}

func (c *CurrencyConverter) Convert(amount float64, to string, from string) (float64, error) {
	return 0, nil
}
