// The package provides methods for converting amounts between currencies. The
// exchange rates are provided by the ECB (http://www.ecb.europa.eu/).
//
// Author: Michael Banzon
package currency

import (
	"encoding/xml"
	"fmt"
	"github.com/mbanzon/simplehttp"
	"time"
)

const (
	ecbResourceUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
	dateFormat     = "2006-01-02"
)

type currencyEnvelope struct {
	Sender string `xml:"Sender>name"`
	Cube   []cube `xml:"Cube>Cube>Cube"`
}

type timeEnvelope struct {
	Time timeCube `xml:"Cube>Cube"`
}

type timeCube struct {
	Time string `xml:"time,attr"`
}

type cube struct {
	Name string  `xml:"currency,attr"`
	Rate float64 `xml:"rate,attr"`
}

type CurrencyConverter struct {
	date       time.Time
	currencies map[string]float64
}

func NewConverter() (*CurrencyConverter, error) {
	r := simplehttp.NewGetRequest(ecbResourceUrl)
	data, err := r.MakeRequest()
	if err != nil {
		return nil, err
	}

	var c currencyEnvelope
	err = xml.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}

	var t timeEnvelope
	err = xml.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}

	currencyTime, err := time.Parse(dateFormat, t.Time.Time)
	if err != nil {
		return nil, err
	}

	currencies := make(map[string]float64)

	for _, currency := range c.Cube {
		currencies[currency.Name] = currency.Rate
	}

	converter := CurrencyConverter{date: currencyTime, currencies: currencies}
	return &converter, nil
}
