package currency

import (
	"log"
	"math/rand"
	"testing"
)

func init() {
	c, err := NewConverter()
	if err != nil {
		log.Fatal("Error initializing converter.", err)
	}
	converter = c
}

func TestSingleRenew(t *testing.T) {
	_, err := converter.GetSingleCurrencyConverter("EUR", "DKK")
	if err != nil {
		t.Fatal("Couldn't create single converter.", err)
	}
	converter.Renew()
}

func TestSingleConverterCreation(t *testing.T) {
	currencies := converter.GetCurrencies()
	for _, from := range currencies {
		for _, to := range currencies {
			_, err := converter.GetSingleCurrencyConverter(from, to)
			if err != nil {
				t.Fatal("Couldn't create single converter.", err)
			}
		}
	}

	_, err1 := converter.GetSingleCurrencyConverter("EUR", "not present")
	if err1 == nil {
		t.Fatal("Currency shouldn't be present.")
	}

	_, err2 := converter.GetSingleCurrencyConverter("not present", "EUR")
	if err2 == nil {
		t.Fatal("Currency shouldn't be present.")
	}
}

func TestSingleMultiConvert(t *testing.T) {
	var amounts []float64
	for i := 100; i < 200; i++ {
		amounts = append(amounts, float64(i)*rand.Float64())
	}
	s, err := converter.GetSingleCurrencyConverter("EUR", "DKK")
	if err != nil {
		t.Fatal("Couldn't create single converter.", err)
	}

	converted := s.MultiConvert(amounts)

	if len(amounts) != len(converted) {
		t.Fatal("Incorrect number of conversions.")
	}
}
