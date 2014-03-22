package currency

type currencyInformation struct {
	LongName string
	Country  string
}

// Hard coded values with extended information about currencies.
var currencyData = map[string]currencyInformation{
	"EUR": currencyInformation{
		LongName: "Euro",
		Country:  "Economic and Monetary Union of the European Union",
	},
	"USD": currencyInformation{
		LongName: "US Dollars",
		Country:  "Unites States of America",
	},
	"JPY": currencyInformation{
		LongName: "Yen",
		Country:  "Japan",
	},
	"BGN": currencyInformation{
		LongName: "Bulgarian lev",
		Country:  "Bulgaria",
	},
	"CZK": currencyInformation{
		LongName: "Czech koruna",
		Country:  "Czech Republic",
	},
	"DKK": currencyInformation{
		LongName: "Danish kroner",
		Country:  "Denmark",
	},
	"GBP": currencyInformation{
		LongName: "Pound sterling",
		Country:  "United Kingdom",
	},
	"HUF": currencyInformation{
		LongName: "Hungarian forint",
		Country:  "Hungary",
	},
	"LTL": currencyInformation{
		LongName: "Lithuanian litas",
		Country:  "Lithuania",
	},
	"LVL": currencyInformation{
		LongName: "Latvian lats",
		Country:  "Latvia",
	},
	"PLN": currencyInformation{
		LongName: "Polish złoty",
		Country:  "Poland",
	},
	"RON": currencyInformation{
		LongName: "Romanian leu",
		Country:  "Romania",
	},
	"SEK": currencyInformation{
		LongName: "Swedish krona",
		Country:  "Sweden",
	},
	"CHF": currencyInformation{
		LongName: "Swiss franc",
		Country:  "Switzerland",
	},
	"NOK": currencyInformation{
		LongName: "Norwegian krone",
		Country:  "Norway",
	},
	"HRK": currencyInformation{
		LongName: "Croatian kuna",
		Country:  "Croatia",
	},
	"RUB": currencyInformation{
		LongName: "Russian ruble",
		Country:  "Croatia",
	},
	"TRY": currencyInformation{
		LongName: "Turkish lira",
		Country:  "Turkey",
	},
	"AUD": currencyInformation{
		LongName: "Australian dollar",
		Country:  "Australia",
	},
	"BRL": currencyInformation{
		LongName: "Brazilian real",
		Country:  "Brazil",
	},
	"CAD": currencyInformation{
		LongName: "Canadian dollar",
		Country:  "Canada",
	},
	"CNY": currencyInformation{
		LongName: "Renminbi",
		Country:  "China",
	},
	"HKD": currencyInformation{
		LongName: "Hong Kong dollar",
		Country:  "Hong Kong",
	},
	"IDR": currencyInformation{
		LongName: "Indonesian rupiah",
		Country:  "Indonesia",
	},
	"ILS": currencyInformation{
		LongName: "Israeli new shekel",
		Country:  "Israel",
	},
	"INR": currencyInformation{
		LongName: "Indian rupee",
		Country:  "India",
	},
	"KRW": currencyInformation{
		LongName: "South Korean won",
		Country:  "South Korea",
	},
	"MXN": currencyInformation{
		LongName: "Mexican peso",
		Country:  "Mexico",
	},
	"MYR": currencyInformation{
		LongName: "Malaysian ringgit",
		Country:  "Malaysia",
	},
	"NZD": currencyInformation{
		LongName: "New Zealand dollar",
		Country:  "New Zealand",
	},
	"PHP": currencyInformation{
		LongName: "Philippine peso",
		Country:  "Philippines",
	},
	"SGD": currencyInformation{
		LongName: "Singapore dollar",
		Country:  "Singapore",
	},
	"THB": currencyInformation{
		LongName: "Thai baht",
		Country:  "Thailand",
	},
	"ZAR": currencyInformation{
		LongName: "South African rand",
		Country:  "South Africa",
	},
}
