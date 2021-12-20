package main

import (
	"fmt"

	hitbtc "github.com/childlycorp/alpha-hitbtc-connector"
)

const (
	API_KEY    = "ayH64Q63cDWfCu2vAHXNy1cYbFxZMR2Q"
	API_SECRET = "-YVnorkMy09k77mWKf3Vi5nbZsNxBr6B"
)

func main() {
	// hitbtc client
	hitbtc := hitbtc.New(API_KEY, API_SECRET)

	// GetCurrencies
	// currencies, _ := hitbtc.GetCurrencies()

	// for id, cur := range currencies {
	// 	fmt.Println(id, cur)
	// }

	// GetSymbols
	symbols, err := hitbtc.GetSymbols()
	fmt.Println(len(symbols), err)

	for _, symbol := range symbols {
		fmt.Println(symbol.BaseCurrency, symbol.QuoteCurrency)
	}

	// GetBalances
	// balances, err := hitbtc.GetBalances()
	// fmt.Println(len(balances), err)

	// for _, bal := range balances {
	// 	if bal.Currency == "BTC" {
	// 		fmt.Println(bal.Available)
	// 	}
	// }

}
