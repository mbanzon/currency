// The package provides methods for converting amounts between currencies. The
// exchange rates are provided by the ECB (http://www.ecb.europa.eu/).
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
	lastUpdate       time.Time
	date             time.Time
	currencies       map[string]float64
	singleConverters []*SingleCurrencyConverter
}

// Creates a new converter by fetching the required data from the ECB.
func NewConverter() (*CurrencyConverter, error) {
	currencyTime, currencies, err := parseEcbData()
	if err != nil {
		return nil, err
	}
	converter := CurrencyConverter{date: currencyTime, currencies: currencies, lastUpdate: time.Now()}
	return &converter, nil
}

func (c *CurrencyConverter) GetCurrencies() []string {
	currencies := make([]string, len(c.currencies))
	index := 0
	for currency, _ := range c.currencies {
		currencies[index] = currency
		index++
	}
	return currencies
}

// Returns true if the currency is known by the converter.
func (c *CurrencyConverter) HasCurrency(currency string) bool {
	_, ok := c.currencies[currency]
	return ok
}

// Calculates the age in days of the CurrencyConverter. The age is calculated
// using the date supplied in the currency feed from the ECB.
func (c *CurrencyConverter) Age() float64 {
	delta := time.Now().Sub(c.lastUpdate)
	return delta.Minutes()
}

// Returns true if the currencies stores are so old they should be renewed from
// the ECB server.
func (c *CurrencyConverter) ShouldRenew(minutes float64) bool {
	if c.Age() >= minutes {
		return true
	}
	return false
}

// Renew the currency data by fetching the  from the ECB server. This will
// also update all the SingleCurrencyConverter created from this CurrencyConverter.
func (c *CurrencyConverter) Renew() error {
	date, currencies, err := parseEcbData()
	if err != nil {
		return err
	} else {
		c.lastUpdate = time.Now()
		c.date = date
		c.currencies = currencies
		for _, s := range c.singleConverters {
			s.renew(c)
		}
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
	c.singleConverters = append(c.singleConverters, &converter)
	return &converter, nil
}

func (c *CurrencyConverter) GetCurrencyInformation(currency string) (longName string, country string, err error) {
	if info, ok := currencyData[currency]; ok {
		longName = info.LongName
		country = info.Country
	} else {
		err = errors.New(fmt.Sprintf(unknown, currency))
	}

	return
}
