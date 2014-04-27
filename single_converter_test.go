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
	_, err := converter.GetSingleCurrencyConverter(KNOWN_CURRENCY_1, KNOWN_CURRENCY_2)
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

	_, err1 := converter.GetSingleCurrencyConverter(KNOWN_CURRENCY_1, NOT_PRESENT_CURRENCY)
	if err1 == nil {
		t.Fatal("Currency shouldn't be present.")
	}

	_, err2 := converter.GetSingleCurrencyConverter(NOT_PRESENT_CURRENCY, KNOWN_CURRENCY_1)
	if err2 == nil {
		t.Fatal("Currency shouldn't be present.")
	}
}

func TestSingleMultiConvert(t *testing.T) {
	var amounts []float64
	for i := 100; i < 200; i++ {
		amounts = append(amounts, float64(i)*rand.Float64())
	}
	s, err := converter.GetSingleCurrencyConverter(KNOWN_CURRENCY_1, KNOWN_CURRENCY_2)
	if err != nil {
		t.Fatal("Couldn't create single converter.", err)
	}

	converted := s.MultiConvert(amounts)

	if len(amounts) != len(converted) {
		t.Fatal("Incorrect number of conversions.")
	}
}
