package currency

import (
	"encoding/xml"
	"github.com/mbanzon/simplehttp"
	"time"
)

const (
	ecbResourceUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml" // The url to fetch from.
	dateFormat     = "2006-01-02"                                                   // The format of the dates to be parsed.
)

// Structure for getting the currencies out of the nested cube elements.
type currencyEnvelope struct {
	Sender string `xml:"Sender>name"`
	Cube   []cube `xml:"Cube>Cube>Cube"`
}

type timeEnvelope struct {
	Time timeCube `xml:"Cube>Cube"`
}

// structure used for parsing the time.
type timeCube struct {
	Time string `xml:"time,attr"`
}

// structure used for parsing the currencies.
type cube struct {
	Name string  `xml:"currency,attr"`
	Rate float64 `xml:"rate,attr"`
}

func parseEcbData() (time.Time, map[string]float64, error) {
	r := simplehttp.NewGetRequest(ecbResourceUrl)
	data, err := r.MakeRequest()
	if err != nil {
		return time.Time{}, nil, err
	}

	var c currencyEnvelope
	err = xml.Unmarshal(data, &c) // first unmarshal to get currencies
	if err != nil {
		return time.Time{}, nil, err
	}

	var t timeEnvelope
	err = xml.Unmarshal(data, &t) // second unmarshal to get time
	if err != nil {
		return time.Time{}, nil, err
	}

	currencyTime, err := time.Parse(dateFormat, t.Time.Time)
	if err != nil {
		return time.Time{}, nil, err
	}

	currencies := make(map[string]float64)

	currencies[eur] = 1

	for _, currency := range c.Cube {
		currencies[currency.Name] = currency.Rate
	}

	return currencyTime, currencies, nil
}
