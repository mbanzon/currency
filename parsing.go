package currency

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var alternativeResourceUrl string
var client *http.Client

func init() {
	client = http.DefaultClient
}

func SetClient(c *http.Client) {
	client = c
}

const (
	ecbResourceUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml" // The url to fetch from.
	dateFormat     = "2006-01-02"                                                   // The format of the dates to be parsed.
)

// Structure for getting the currencies out of the nested cube elements.
type currencyEnvelope struct {
	Sender string `xml:"Sender>name"`
	Cube   []cube `xml:"Cube>Cube>Cube"`
}

// structure used when parsing the time from the XML.
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

// Fetches the data from the ECB and returns a map of parsed currencies.
func parseEcbData() (time.Time, map[string]float64, error) {
	var res *http.Response
	var err error

	if alternativeResourceUrl == "" {
		res, err = client.Get(ecbResourceUrl)
	} else {
		res, err = client.Get(alternativeResourceUrl)
	}

	if err != nil {
		return time.Time{}, nil, err
	}

	if res.StatusCode != http.StatusOK {
		return time.Time{}, nil, fmt.Errorf("Error getting currency data from server, code: %d", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return time.Time{}, nil, err
	}

	err = res.Body.Close()
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
