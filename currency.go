// The package provides methods for converting amounts between currencies. The
// exchange rates are provided by the ECB (http://www.ecb.europa.eu/).
//
// Author: Michael Banzon
package currency

import (
	"fmt"
	"github.com/mbanzon/simplehttp"
	"time"
)

const ecbResourceUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type Envelope struct {
	subject string
	Sender  string `xml:"Sender>name"`
	Cube    []Cube `xml:"Cube>Cube>Cube"`
	//	Time    TimeCube `xml:"Cube>Cube"`
}

type TimeCube struct {
	Time string `xml:"time,attr"`
}

type Cube struct {
	//time     string `xml:"time,attr"`
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

type CurrencyConverter struct {
	date       time.Time
	currencies map[string]float64
}

func NewConverter() (*CurrencyConverter, error) {
	var e Envelope
	r := simplehttp.NewGetRequest(ecbResourceUrl)
	r.MakeXMLRequest(&e)

	fmt.Printf("%#v\n", e)

	var foo map[string]string
	re := simplehttp.NewGetRequest(ecbResourceUrl)
	re.MakeXMLRequest(&foo)

	fmt.Printf("%#v\n", foo)

	return &CurrencyConverter{}, nil
}
