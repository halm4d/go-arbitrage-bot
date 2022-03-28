package arb

type SymbolResp struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	Status     string `json:"status"`
}

type ExchangeInfoResp struct {
	Symbols []SymbolResp `json:"symbols"`
}

type BookTickerSocketResp struct {
	U  int    `json:"u"`
	S  string `json:"s"`
	B  string `json:"b"`
	B1 string `json:"B"`
	A  string `json:"a"`
	A1 string `json:"A"`
}
