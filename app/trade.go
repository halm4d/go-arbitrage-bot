package app

type Trades []Trade

type Trade struct {
	From   string
	To     string
	Symbol string
	Type   Type
}
