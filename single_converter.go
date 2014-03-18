package currency

import (
	"errors"
	"fmt"
	"time"
)

// The SingleCurrencyConverter struct holds data about how to convert amounts
// between two pre-defined currencies.
type SingleCurrencyConverter struct {
	date             time.Time
	from, to         string
	fromRate, toRate float64
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

func (c *SingleCurrencyConverter) renew(r *CurrencyConverter) error {
	fromRate, fromOk := r.currencies[c.from]
	if !fromOk {
		return errors.New(fmt.Sprintf(unknown, c.from))
	}
	toRate, toOk := r.currencies[c.to]
	if !toOk {
		return errors.New(fmt.Sprintf(unknown, c.to))
	}

	c.fromRate = fromRate
	c.toRate = toRate
	c.date = r.date
	return nil
}
