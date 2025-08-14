package util

type Currency int

const (
	USD Currency = iota
	EUR
	VND
)

var currencyNames = map[Currency]string{
	USD: "USD",
	EUR: "EUR",
	VND: "VND",
}

func IsSupportedCurrency(currency string) bool {
	for _, v := range currencyNames {
		if v == currency {
			return true
		}
	}
	return false
}
