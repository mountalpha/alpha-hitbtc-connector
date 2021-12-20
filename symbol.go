package hitbtc

// Symbol represents data of a Currency Pair on a market.
type Symbol struct {
	Type               string  `json:"type"`
	BaseCurrency       string  `json:"base_currency"`
	QuoteCurrency      string  `json:"quote_currency"`
	QuantityIncrement  float64 `json:"quantity_increment,string"`
	TickSize           float64 `json:"tick_size,string"`
	TakeRate           float64 `json:"take_rate,string"`
	MakeRate           float64 `json:"make_rate,string"`
	FeeCurrency        string  `json:"fee_currency"`
	MarginTrading      bool    `json:"margin_trading"`
	MaxInitialLeverage float64 `json:"max_initial_leverage,string"`
}
