package app

import (
	"errors"
	"github.com/halm4d/arbitragecli/client"
	"github.com/halm4d/arbitragecli/util"
	"sort"
	"strconv"
)

func AllCryptoCurrency() []string {
	//return []string{"BTC", "ETH", "GMT", "LUNA", "BNB", "AVAX", "XRP", "SOL", "WAVES", "JASMY", "SAND", "ADA", "FTM", "DOT", "GALA", "LINK", "ALPINE", "SHIB", "RUNE", "APE", "MANA", "ATOM", "LTC", "NEAR", "CAKE", "CELO", "DOGE", "MATIC", "XNO", "TRX", "ZEC", "SLP", "AAVE", "BCH", "SANTOS", "GRT", "XLM", "AXS", "SAND", "PEOPLE", "VET", "FIL", "ANC", "EGLD", "ALGO", "ENS", "NEO", "ETC", "UNI", "MINA", "ICP", "XMR", "ALICE", "LAZIO", "EOS", "ROSE", "CRV", "BAT", "AGLD", "SNX", "SLP", "DAI", "SUSHI", "XTZ", "ONE", "TLM", "LUNA", "IMX", "CHZ", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX", "MASK", "GLMR", "MBOX", "1INCH", "KNC", "WBTC", "TROY", "HBAR", "JOE", "REN", "CVC", "GRT", "PORTO", "LOKA", "VOXEL", "POLS", "LINA", "BTT", "MIR", "WIN", "WOO", "FTT", "C98", "QTUM", "ADX", "TWT", "PAXG", "BETA", "BICO", "HOT", "FLOW", "JST", "KSM", "DAR", "DCR", "KLAY", "HNT", "MKR", "SCRT", "SYS", "KDA", "SPELL", "INJ", "SUN", "ZEN", "COMP", "IDEX", "BTTC", "OG", "BAND", "KP3R", "MC", "DENT", "IOST", "ABBC", "PYR", "ARPA", "BAKE", "API3", "DUSK", "COCOS", "OXT", "CELR", "XEC", "REEF", "T", "ZRX", "BTCDOWN", "RARE", "IOT", "OOKI", "CTSI", "RNDR", "HIGH", "ANT", "POND", "ICX", "SRM", "XVS", "CVP", "UNFI", "MTL", "QNT", "BTCUP", "JUV", "ZIL", "VITE", "BAR", "CVX", "COTI", "KAVA", "ANKR", "MOVR", "TFUEL", "RVN", "MDX", "DODO", "ANY", "QI", "OCEAN", "FLUX", "SUPER", "YGG", "WAXP", "STX", "FET", "ORN", "DIA", "EPS", "STRAX", "TORN", "MCONTENT", "LOOM", "AMA", "HIVE", "HT", "BSV", "ONT", "WRX", "SFP", "TVK", "POWR", "MXC", "ACA", "GXS", "REQ", "LIT", "CITY", "LUNA", "CLV", "RAY", "AUDIO", "NKN", "STORJ", "FIDA", "NAV", "CTK", "WOO", "JASMY", "ASTR", "COS", "MDT", "GTC", "ATA", "ALPHA", "UTK", "BTCST", "BEL", "BAL", "MFT", "ASR", "ATM", "BNX", "OKB", "TKO", "WTC", "NEXO", "CKB", "CRO", "PLA", "SG", "STEEM", "CHESS", "SKL", "HARD", "DGB", "VIDT", "XVG", "SC", "WAN", "DEGO", "FRONT", "PSG", "NANO", "TT", "XEM", "HEX", "ILV", "LPT", "BURGER", "RAD", "BNT", "BLZ", "ELF", "FORTH", "FARM", "KEY", "RAMP", "BMX", "NMR", "WNXM", "AVA", "PERP", "TCT", "FIS", "YFII", "FOR", "IRIS", "ERN", "POLY", "ALCX", "DFA", "SEELE", "TRU", "NULS", "FLM", "CTXC", "AKRO", "BEAM", "ACH", "LSK"}
	//return []string{"BTC", "ETH","SOL", "LTC", "DOGE", "MATIC", "BCH", "SANTOS", "GRT", "XLM", "AXS", "ANC", "EGLD", "ALGO", "MINA", "ICP", "XMR", "ALICE", "BAT", "AGLD", "SNX",  "ONE", "TLM", "LUNA", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX"}
	return []string{"BNB", "BTC", "COCOS"}
}

type Symbols []Symbol

type Symbol struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	Price      float64
}

func NewSymbols() (Symbols, Symbols) {
	prices := make(chan *[]client.TickerPriceResp)
	exchange := make(chan *[]client.SymbolResp)

	go client.GetPrices(prices)
	go client.GetExchangeInfo(exchange)

	var busdSymbols Symbols
	var symbols Symbols
	priceResp := *<-prices
	exchangeResp := *<-exchange
	for _, price := range priceResp {
		for _, exchange := range exchangeResp {
			if price.Symbol != exchange.Symbol {
				continue
			}
			if exchange.QuoteAsset == "BUSD" || exchange.BaseAsset == "BUSD" {
				parsedPrice, _ := strconv.ParseFloat(price.Price, 64)
				busdSymbols = append(busdSymbols, Symbol{
					Symbol:     exchange.Symbol,
					BaseAsset:  exchange.BaseAsset,
					QuoteAsset: exchange.QuoteAsset,
					Price:      parsedPrice,
				})
			}
			if util.Contains(AllCryptoCurrency(), exchange.QuoteAsset) && util.Contains(AllCryptoCurrency(), exchange.BaseAsset) {
				parsedPrice, _ := strconv.ParseFloat(price.Price, 64)
				symbols = append(symbols, Symbol{
					Symbol:     exchange.Symbol,
					BaseAsset:  exchange.BaseAsset,
					QuoteAsset: exchange.QuoteAsset,
					Price:      parsedPrice,
				})
			}
		}
	}
	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i].Symbol < symbols[j].Symbol
	})
	sort.Slice(busdSymbols, func(i, j int) bool {
		return busdSymbols[i].Symbol < busdSymbols[j].Symbol
	})
	return symbols, busdSymbols
}

func (s Symbols) updatePrices() {
	pricesChan := make(chan *[]client.TickerPriceResp)
	go client.GetPrices(pricesChan)
	priceResp := *<-pricesChan
	for i, symbol := range s {
		for _, price := range priceResp {
			if symbol.Symbol == price.Symbol {
				parsedPrice, _ := strconv.ParseFloat(price.Price, 64)
				s[i].Price = parsedPrice
			}
		}
	}
}

func (s Symbols) findBySymbol(target string) (Symbol, error) {
	for _, symbol := range s {
		if symbol.Symbol == target {
			return symbol, nil
		}
	}
	return Symbol{Symbol: "", BaseAsset: "", QuoteAsset: "", Price: -1}, errors.New("target symbol not found")
}

func (s Symbols) findAllByAsset(asset string) Symbols {
	var symbols Symbols
	for _, symbol := range s {
		if symbol.BaseAsset == asset || symbol.QuoteAsset == asset {
			symbols = append(symbols, symbol)
		}
	}
	return symbols
}

func (s Symbols) findByAssets(asset1 string, asset2 string) Symbols {
	var symbols Symbols
	for _, symbol := range s {
		if symbol.BaseAsset == asset1 || symbol.QuoteAsset == asset1 || symbol.BaseAsset == asset2 || symbol.QuoteAsset == asset2 {
			symbols = append(symbols, symbol)
		}
	}
	return symbols
}

func (s Symbols) findByAssetPairs(asset1 string, asset2 string) (Symbol, error) {
	for _, symbol := range s {
		if (symbol.BaseAsset == asset1 || symbol.QuoteAsset == asset1) && (symbol.BaseAsset == asset2 || symbol.QuoteAsset == asset2) {
			return symbol, nil
		}
	}
	return Symbol{}, errors.New("symbol not found")
}

//func (s Symbols) removeAssets(asset ...string) Symbols {
//	var symbols Symbols
//	for i, symbol := range s {
//		if contains(asset, symbol.BaseAsset) || contains(asset, symbol.QuoteAsset)  {
//
//		}
//	}
//}

func contains(slice []string, t string) bool {
	for _, s := range slice {
		if s == t {
			return true
		}
	}
	return false
}
