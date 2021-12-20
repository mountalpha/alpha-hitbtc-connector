package hitbtc

// Currency represents currency data.
type Currency struct {
	FullName          string `json:"full_name"`
	PayinEnabled      bool   `json:"payin_enabled"`
	PayoutEnabled     bool   `json:"payout_enabled"`
	TransferEnabled   bool   `json:"transfer_enabled"`
	PrecisionTransfer string `json:"precision_transfer"`
}
