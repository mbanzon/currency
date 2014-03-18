package currency

type CurrencyInformation struct {
	LongName string
	Country  string
}

var currencyData = map[string]CurrencyInformation{
	"EUR": CurrencyInformation{
		LongName: "Euro",
		Country:  "Economic and Monetary Union of the European Union",
	},
	"USD": CurrencyInformation{
		LongName: "US Dollars",
		Country:  "Unites States of America",
	},
	"JPY": CurrencyInformation{
		LongName: "Yen",
		Country:  "Japan",
	},
	"BGN": CurrencyInformation{
		LongName: "Bulgarian lev",
		Country:  "Bulgaria",
	},
	"CZK": CurrencyInformation{
		LongName: "Czech koruna",
		Country:  "Czech Republic",
	},
	"DKK": CurrencyInformation{
		LongName: "Danish kroner",
		Country:  "Denmark",
	},
	"GBP": CurrencyInformation{
		LongName: "Pound sterling",
		Country:  "United Kingdom",
	},
	"HUF": CurrencyInformation{
		LongName: "Hungarian forint",
		Country:  "Hungary",
	},
	"LTL": CurrencyInformation{
		LongName: "Lithuanian litas",
		Country:  "Lithuania",
	},
	"LVL": CurrencyInformation{
		LongName: "Latvian lats",
		Country:  "Latvia",
	},
	"PLN": CurrencyInformation{
		LongName: "Polish złoty",
		Country:  "Poland",
	},
	"RON": CurrencyInformation{
		LongName: "Romanian leu",
		Country:  "Romania",
	},
	"SEK": CurrencyInformation{
		LongName: "Swedish krona",
		Country:  "Sweden",
	},
	"CHF": CurrencyInformation{
		LongName: "Swiss franc",
		Country:  "Switzerland",
	},
	"NOK": CurrencyInformation{
		LongName: "Norwegian krone",
		Country:  "Norway",
	},
	"HRK": CurrencyInformation{
		LongName: "Croatian kuna",
		Country:  "Croatia",
	},
	"RUB": CurrencyInformation{
		LongName: "Russian ruble",
		Country:  "Croatia",
	},
	"TRY": CurrencyInformation{
		LongName: "Turkish lira",
		Country:  "Turkey",
	},
	"AUD": CurrencyInformation{
		LongName: "Australian dollar",
		Country:  "Australia",
	},
	"BRL": CurrencyInformation{
		LongName: "Brazilian real",
		Country:  "Brazil",
	},
	"CAD": CurrencyInformation{
		LongName: "Canadian dollar",
		Country:  "Canada",
	},
	"CNY": CurrencyInformation{
		LongName: "Renminbi",
		Country:  "China",
	},
	"HKD": CurrencyInformation{
		LongName: "Hong Kong dollar",
		Country:  "Hong Kong",
	},
	"IDR": CurrencyInformation{
		LongName: "Indonesian rupiah",
		Country:  "Indonesia",
	},
	"ILS": CurrencyInformation{
		LongName: "Israeli new shekel",
		Country:  "Israel",
	},
	"INR": CurrencyInformation{
		LongName: "Indian rupee",
		Country:  "India",
	},
	"KRW": CurrencyInformation{
		LongName: "South Korean won",
		Country:  "South Korea",
	},
	"MXN": CurrencyInformation{
		LongName: "Mexican peso",
		Country:  "Mexico",
	},
	"MYR": CurrencyInformation{
		LongName: "Malaysian ringgit",
		Country:  "Malaysia",
	},
	"NZD": CurrencyInformation{
		LongName: "New Zealand dollar",
		Country:  "New Zealand",
	},
	"PHP": CurrencyInformation{
		LongName: "Philippine peso",
		Country:  "Philippines",
	},
	"SGD": CurrencyInformation{
		LongName: "Singapore dollar",
		Country:  "Singapore",
	},
	"THB": CurrencyInformation{
		LongName: "Thai baht",
		Country:  "Thailand",
	},
	"ZAR": CurrencyInformation{
		LongName: "South African rand",
		Country:  "South Africa",
	},
}
