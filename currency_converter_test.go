package currency

import (
	"log"
	"math"
	"math/rand"
	"testing"
)

var converter *CurrencyConverter

const (
	KNOWN_CURRENCY_1     = "EUR"
	KNOWN_CURRENCY_2     = "DKK"
	NOT_PRESENT_CURRENCY = "not present"
	NOT_WORKING_URL      = "http://what.ever.dude"
)

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

func TestConvertion(t *testing.T) {
	for i := 0; i < 100; i++ {
		amount := rand.Float64() * float64(i)
		_, err := converter.Convert(amount, KNOWN_CURRENCY_1, KNOWN_CURRENCY_2)
		if err != nil {
			t.Fatal("Convertion failed.", err)
		}
	}

	_, err := converter.Convert(100, KNOWN_CURRENCY_1, NOT_PRESENT_CURRENCY)
	if err == nil {
		t.Fatal("Currency shouldn't be present.")
	}
}

func TestNoConvertion(t *testing.T) {
	for i := 0; i < 100; i++ {
		amount := rand.Float64() * float64(i)
		converted, err := converter.Convert(amount, KNOWN_CURRENCY_1, KNOWN_CURRENCY_1)
		if err != nil {
			t.Fatal("Convertion failed.", err)
		}
		if amount != converted {
			t.Fatal("Convertion to/from same currency gave different amount.")
		}
	}
}

func TestHasCurrencies(t *testing.T) {
	currencies := converter.GetCurrencies()
	for _, currency := range currencies {
		if !converter.HasCurrency(currency) {
			t.Fatalf("Converter doesn't have currency: %s", currency)
		}
	}
	if converter.HasCurrency("not present") {
		t.Fatal("Converter has currency: %s")
	}
}

func TestCurrencyAge(t *testing.T) {
	age := converter.Age() - converter.CurrencyAge()
	if age > 0 {
		t.Fatal("Currency age is newer than fetch.", age)
	}
}

func TestRenew(t *testing.T) {
	err := converter.Renew()
	if err != nil {
		t.Fatal("Renew failed.", err)
	}

	alternativeResourceUrl = NOT_WORKING_URL
	err = converter.Renew()
	if err == nil {
		t.Fatal("Alternative URL should not be working.")
	}
	alternativeResourceUrl = ""
}

func TestShouldRenew(t *testing.T) {
	if !converter.ShouldRenew(0) {
		t.Fatal("Converter should be older than 0.")
	}

	if converter.ShouldRenew(math.MaxFloat64) {
		t.Fatalf("Converter shouldn't be older than %f (%f).", math.MaxFloat64, converter.Age())
	}
}

func TestMultiConvert(t *testing.T) {
	var amounts []float64
	for i := 100; i < 200; i++ {
		amounts = append(amounts, float64(i)*rand.Float64())
	}
	converted, err := converter.MultiConvert(amounts, KNOWN_CURRENCY_1, KNOWN_CURRENCY_2)

	if err != nil {
		t.Fatal("Couldn't multiconvert.", err)
	}

	if len(amounts) != len(converted) {
		t.Fatal("Incorrect number of conversions.")
	}

	_, err2 := converter.MultiConvert(amounts, NOT_PRESENT_CURRENCY, KNOWN_CURRENCY_1)
	if err2 == nil {
		t.Fatal("Currency shouldn't be present.")
	}
}

func TestCurrencyInformation(t *testing.T) {
	for _, currency := range converter.GetCurrencies() {
		_, _, err := converter.GetCurrencyInformation(currency)
		if err != nil {
			t.Fatal("Couldn't get currency information.", err)
		}
	}

	_, _, err := converter.GetCurrencyInformation(NOT_PRESENT_CURRENCY)
	if err == nil {
		t.Fatal("Currency shouldn't be present.")
	}
}

func TestFailingCreation(t *testing.T) {
	alternativeResourceUrl = NOT_WORKING_URL
	_, err := NewConverter()
	if err == nil {
		t.Fatal("Alternative URL should not be working.")
	}
	alternativeResourceUrl = ""
}
