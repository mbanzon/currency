currency
========

[![Build Status](https://travis-ci.org/mbanzon/currency.png?branch=master)](https://travis-ci.org/mbanzon/currency)

[![GoDoc](https://godoc.org/gopkg.in/mbanzon/currency.v0?status.svg)](https://godoc.org/gopkg.in/mbanzon/currency.v0)

This is a simple currency conversion library for Go.

The currencies exchange rates are fetched from the European Central Bank (ECB).
The raw XML feed with currencies can be found here: http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml

Import the library to your Go workspace:

```
go get gopkg.in/mbanzon/currency.v0
```

Import the package:

```Go
import "gopkg.in/mbanzon/currency.v0"
```

The library enables easy conversion between currencies like this:

```Go
ecb, _ := currency.NewConverter()
from, to := "DKK", "GBP"
amount := 100.0
converted, _ := ecb.Convert(amount, from, to)
fmt.Printf("%f in %s is %f in %s\n", amount, from, converted, to)
```

The library lets you create a fixed converter that converts between two fixed currencies:

```Go
from, to := "DKK", "GBP"
amount := 100.0
ecb, _ := currency.NewConverter()
single, _ := ecb.GetSingleCurrencyConverter(from, to)
fmt.Printf("%f in %s is %f in %s\n", amount, from, single.Convert(amount), to)
```

The code is released under a 3-clause BSD license. See the LICENSE file for more information.