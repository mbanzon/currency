package currency

import (
	"log"
	"math/rand"
	"testing"
)

var converter *CurrencyConverter

func init() {
	c, err := NewConverter()
	if err != nil {
		log.Fatal("Error initializing converter.", err)
	}
	converter = c
}

func TestConverterCreation(t *testing.T) {
	currencies := converter.GetCurrencies()

	// must add one for EUR
	if len(currencies)+1 != len(currencyData) {
		t.Fatalf("Wrong currency number (%d vs %d)", len(currencies), len(currencyData))
	}
}

func TestNoConvertion(t *testing.T) {
	for i := 0; i < 100; i++ {
		amount := rand.Float64() * float64(i)
		converted, err := converter.Convert(amount, "EUR", "EUR")
		if err != nil {
			t.Fatal("Convertion failed.", err)
		}
		if amount != converted {
			t.Fatal("Convertion to/from same currency gave different amount.")
		}
	}
}

func TestRenew(t *testing.T) {
	err := converter.Renew()
	if err != nil {
		t.Fatal("Renew failed.", err)
	}
}
