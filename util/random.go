package util

import (
	"math/rand"
	"time"
)

var currencies = []string{"USD", "EUR", "VND", "JPY", "GBP"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomID() int64 {
	// Giả sử ID từ 1 đến 1e9
	return rand.Int63n(1_000_000_000) + 1
}

func RandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Emma"}
	return names[rand.Intn(len(names))]
}

func RandomMoney() int64 {
	// Số tiền từ 1 đến 10_000
	return rand.Int63n(10_000) + 1
}

func RandomCurrency() string {
	return currencies[rand.Intn(len(currencies))]
}
