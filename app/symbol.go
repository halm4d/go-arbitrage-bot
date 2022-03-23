package app

import (
	"errors"
	"github.com/halm4d/arbitragecli/client"
	"github.com/halm4d/arbitragecli/constants"
	"github.com/halm4d/arbitragecli/util"
	"strconv"
	"strings"
)

func AllCryptoCurrency() []string {
	//return []string{"BTC", "ETH", "GMT", "LUNA", "BNB", "AVAX", "XRP", "SOL", "WAVES", "JASMY", "SAND", "ADA", "FTM", "DOT", "GALA", "LINK", "ALPINE", "SHIB", "RUNE", "APE", "MANA", "ATOM", "LTC", "NEAR", "CAKE", "CELO", "DOGE", "MATIC", "XNO", "TRX", "ZEC", "SLP", "AAVE", "BCH", "SANTOS", "GRT", "XLM", "AXS", "SAND", "PEOPLE", "VET", "FIL", "ANC", "EGLD", "ALGO", "ENS", "NEO", "ETC", "UNI", "MINA", "ICP", "XMR", "ALICE", "LAZIO", "EOS", "ROSE", "CRV", "BAT", "AGLD", "SNX", "SLP", "DAI", "SUSHI", "XTZ", "ONE", "TLM", "LUNA", "IMX", "CHZ", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX", "MASK", "GLMR", "MBOX", "1INCH", "KNC", "WBTC", "TROY", "HBAR", "JOE", "REN", "CVC", "GRT", "PORTO", "LOKA", "VOXEL", "POLS", "LINA", "BTT", "MIR", "WIN", "WOO", "FTT", "C98", "QTUM", "ADX", "TWT", "PAXG", "BETA", "BICO", "HOT", "FLOW", "JST", "KSM", "DAR", "DCR", "KLAY", "HNT", "MKR", "SCRT", "SYS", "KDA", "SPELL", "INJ", "SUN", "ZEN", "COMP", "IDEX", "BTTC", "OG", "BAND", "KP3R", "MC", "DENT", "IOST", "ABBC", "PYR", "ARPA", "BAKE", "API3", "DUSK", "COCOS", "OXT", "CELR", "XEC", "REEF", "T", "ZRX", "BTCDOWN", "RARE", "IOT", "OOKI", "CTSI", "RNDR", "HIGH", "ANT", "POND", "ICX", "SRM", "XVS", "CVP", "UNFI", "MTL", "QNT", "BTCUP", "JUV", "ZIL", "VITE", "BAR", "CVX", "COTI", "KAVA", "ANKR", "MOVR", "TFUEL", "RVN", "MDX", "DODO", "ANY", "QI", "OCEAN", "FLUX", "SUPER", "YGG", "WAXP", "STX", "FET", "ORN", "DIA", "EPS", "STRAX", "TORN", "MCONTENT", "LOOM", "AMA", "HIVE", "HT", "BSV", "ONT", "WRX", "SFP", "TVK", "POWR", "MXC", "ACA", "GXS", "REQ", "LIT", "CITY", "LUNA", "CLV", "RAY", "AUDIO", "NKN", "STORJ", "FIDA", "NAV", "CTK", "WOO", "JASMY", "ASTR", "COS", "MDT", "GTC", "ATA", "ALPHA", "UTK", "BTCST", "BEL", "BAL", "MFT", "ASR", "ATM", "BNX", "OKB", "TKO", "WTC", "NEXO", "CKB", "CRO", "PLA", "SG", "STEEM", "CHESS", "SKL", "HARD", "DGB", "VIDT", "XVG", "SC", "WAN", "DEGO", "FRONT", "PSG", "NANO", "TT", "XEM", "HEX", "ILV", "LPT", "BURGER", "RAD", "BNT", "BLZ", "ELF", "FORTH", "FARM", "KEY", "RAMP", "BMX", "NMR", "WNXM", "AVA", "PERP", "TCT", "FIS", "YFII", "FOR", "IRIS", "ERN", "POLY", "ALCX", "DFA", "SEELE", "TRU", "NULS", "FLM", "CTXC", "AKRO", "BEAM", "ACH", "LSK"}
	return []string{"BUSD", "USDT", "ETH", "BTC", "LTC", "BNB", "NEO", "QTUM", "EOS", "SNT", "BNT", "GAS", "USDT", "WTC", "LRC", "OMG", "ZRX", "KNC", "FUN", "SNM", "IOTA", "LINK", "XVG", "MDA", "MTL", "ETC", "DNT", "ZEC", "AST", "DASH", "OAX", "BTG", "REQ", "VIB", "TRX", "POWR", "ARK", "XRP", "ENJ", "STORJ", "KMD", "NULS", "XMR", "AMB", "BAT", "GXS", "QSP", "BTS", "LSK", "MANA", "ADX", "ADA", "XLM", "WABI", "WAVES", "GTO", "ICX", "ELF", "AION", "NEBL", "BRD", "NAV", "RLC", "PIVX", "IOST", "STEEM", "BLZ", "ZIL", "ONT", "XEM", "WAN", "QLC", "SYS", "GRS", "LOOM", "REP", "TUSD", "ZEN", "CVC", "THETA", "IOTX", "QKC", "NXS", "DATA", "SC", "KEY", "NAS", "MFT", "DENT", "ARDR", "HOT", "VET", "DOCK", "POLY", "GO", "RVN", "DCR", "MITH", "REN", "USDC", "ONG", "FET", "CELR", "MATIC", "ATOM", "PHB", "TFUEL", "ONE", "FTM", "ALGO", "DOGE", "DUSK", "ANKR", "WIN", "COS", "COCOS", "TOMO", "PERL", "CHZ", "BAND", "BUSD", "BEAM", "XTZ", "HBAR", "NKN", "STX", "KAVA", "NGN", "ARPA", "CTXC", "BCH", "RUB", "TROY", "VITE", "FTT", "TRY", "EUR", "OGN", "DREP", "TCT", "WRX", "LTO", "MBL", "COTI", "STPT", "SOL", "IDRT", "CTSI", "HIVE", "CHR", "MDT", "STMX", "IQ", "PNT", "GBP", "DGB", "UAH", "COMP", "BIDR", "SXP", "SNX", "VTHO", "IRIS", "MKR", "RUNE", "AUD", "FIO", "AVA", "BAL", "YFI", "DAI", "JST", "SRM", "ANT", "CRV", "SAND", "OCEAN", "NMR", "DOT", "LUNA", "IDEX", "RSR", "PAXG", "WNXM", "TRB", "WBTC", "SUSHI", "YFII", "KSM", "EGLD", "DIA", "UMA", "BEL", "WING", "CREAM", "UNI", "NBS", "OXT", "SUN", "AVAX", "HNT", "BAKE", "BURGER", "FLM", "SCRT", "CAKE", "SPARTA", "ORN", "UTK", "XVS", "ALPHA", "VIDT", "AAVE", "BRL", "NEAR", "FIL", "INJ", "AERGO", "AUDIO", "CTK", "AKRO", "KP3R", "AXS", "HARD", "RENBTC", "SLP", "CVP", "STRAX", "FOR", "UNFI", "FRONT", "ROSE", "HEGIC", "PROM", "SKL", "SUSD", "GLM", "GHST", "DF", "GRT", "JUV", "PSG", "1INCH", "REEF", "OG", "ATM", "ASR", "CELO", "RIF", "BTCST", "TRU", "DEXE", "CKB", "TWT", "FIRO", "BETH", "PROS", "LIT", "VAI", "SFP", "FXS", "DODO", "UFT", "ACM", "AUCTION", "PHA", "TVK", "BADGER", "FIS", "OM", "POND", "DEGO", "ALICE", "BIFI", "LINA", "PERP", "RAMP", "SUPER", "CFX", "EPS", "AUTO", "TKO", "PUNDIX", "TLM", "MIR", "BAR", "FORTH", "EZ", "SHIB", "ICP", "AR", "POLS", "MDX", "MASK", "LPT", "AGIX", "ATA", "GTC", "TORN", "ERN", "KLAY", "BOND", "MLN", "QUICK", "C98", "CLV", "QNT", "FLOW", "XEC", "MINA", "RAY", "FARM", "ALPACA", "MBOX", "VGX", "WAXP", "TRIBE", "GNO", "DYDX", "USDP", "GALA", "ILV", "YGG", "FIDA", "AGLD", "RAD", "BETA", "RARE", "SSV", "LAZIO", "CHESS", "DAR", "BNX", "MOVR", "CITY", "ENS", "QI", "PORTO", "JASMY", "AMP", "PLA", "PYR", "RNDR", "ALCX", "SANTOS", "MC", "ANY", "BICO", "FLUX", "VOXEL", "HIGH", "CVX", "PEOPLE", "OOKI", "SPELL", "UST", "JOE", "ACH", "IMX", "GLMR", "LOKA", "API3", "BTTC", "ACA", "ANC", "BDOT", "XNO", "WOO", "ALPINE", "T", "ASTR", "NBT", "GMT", "KDA", "APE"}
	//return []string{"BUSD", "USDT", "NBT", "BIDR", "BTC", "ETH", "SOL", "LTC", "DOGE", "MATIC", "BCH", "SANTOS", "GRT", "XLM", "AXS", "ANC", "EGLD", "ALGO", "MINA", "ICP", "XMR", "ALICE", "BAT", "AGLD", "SNX", "ONE", "TLM", "LUNA", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX"}
	//return []string{"BNB", "BTC", "COCOS"}
}

type Symbols []Symbol

type Symbol struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	//Price      float64
	BidPrice float64
	AskPrice float64
}

func NewSymbols() (map[string]Symbol, map[string]Symbol) {
	prices := make(chan *[]client.TickerPriceResp)
	exchange := make(chan *[]client.SymbolResp)

	go client.GetPrices(prices)
	go client.GetExchangeInfo(exchange)

	var usdtSymbols = make(map[string]Symbol)
	var symbols = make(map[string]Symbol)
	priceResp := *<-prices
	exchangeResp := *<-exchange
	for _, price := range priceResp {
		for _, exchange := range exchangeResp {
			if exchange.Status != "TRADING" {
				continue
			}
			if price.Symbol != exchange.Symbol {
				continue
			}
			bidPrice, _ := strconv.ParseFloat(price.BidPrice, 64)
			askPrice, _ := strconv.ParseFloat(price.AskPrice, 64)
			if exchange.QuoteAsset == constants.USDT || exchange.BaseAsset == constants.USDT {
				usdtSymbols[exchange.Symbol] = Symbol{
					Symbol:     exchange.Symbol,
					BaseAsset:  exchange.BaseAsset,
					QuoteAsset: exchange.QuoteAsset,
					AskPrice:   askPrice,
					BidPrice:   bidPrice,
				}
			}
			if util.Contains(AllCryptoCurrency(), exchange.QuoteAsset) && util.Contains(AllCryptoCurrency(), exchange.BaseAsset) {
				symbols[exchange.Symbol] = Symbol{
					Symbol:     exchange.Symbol,
					BaseAsset:  exchange.BaseAsset,
					QuoteAsset: exchange.QuoteAsset,
					AskPrice:   askPrice,
					BidPrice:   bidPrice,
				}
			}
		}
	}
	return symbols, usdtSymbols
}

func (s Symbols) updatePrices() {
	pricesChan := make(chan *[]client.TickerPriceResp)
	go client.GetPrices(pricesChan)
	priceResp := *<-pricesChan
	for i, symbol := range s {
		for _, price := range priceResp {
			if symbol.Symbol == price.Symbol {
				bidPrice, _ := strconv.ParseFloat(price.BidPrice, 64)
				askPrice, _ := strconv.ParseFloat(price.AskPrice, 64)
				s[i].BidPrice = bidPrice
				s[i].AskPrice = askPrice
			}
		}
	}
}

func (s *Symbols) findBySymbol(target string) (*Symbol, error) {
	for _, symbol := range *s {
		if symbol.Symbol == target {
			return &symbol, nil
		}
	}
	return &Symbol{Symbol: "", BaseAsset: "", QuoteAsset: "", BidPrice: -1, AskPrice: -1}, errors.New("target symbol not found")
}

func FindAllSymbolByAsset(s map[string]Symbol, asset string) *Symbols {
	var symbols Symbols
	for _, symbol := range s {
		if symbol.BaseAsset == asset || symbol.QuoteAsset == asset {
			symbols = append(symbols, symbol)
		}
	}
	return &symbols
}

func (s *Symbols) findAllByAssets(asset1 string, asset2 string) *Symbols {
	var symbols Symbols
	for _, symbol := range *s {
		if symbol.BaseAsset == asset1 || symbol.QuoteAsset == asset1 || symbol.BaseAsset == asset2 || symbol.QuoteAsset == asset2 {
			symbols = append(symbols, symbol)
		}
	}
	return &symbols
}

func FindByAssetPair(s *map[string]Symbol, asset1 string, asset2 string) (*Symbol, error) {
	for _, symbol := range *s {
		if (symbol.BaseAsset == asset1 || symbol.QuoteAsset == asset1) && (symbol.BaseAsset == asset2 || symbol.QuoteAsset == asset2) {
			return &symbol, nil
		}
	}
	return &Symbol{}, errors.New("symbol not found")
}

func FindBookTickerByAssetPair(bookTickerMap map[string]BookTicker, asset1 string, asset2 string) (*BookTicker, error) {
	for _, bookTicker := range bookTickerMap {
		if (bookTicker.BaseAsset == asset1 || bookTicker.QuoteAsset == asset1) && (bookTicker.BaseAsset == asset2 || bookTicker.QuoteAsset == asset2) {
			return &bookTicker, nil
		}
	}
	return &BookTicker{}, errors.New("symbol not found")
}

func (s *Symbol) getTargetAsset(ignore string) string {
	var targetAsset2 string
	if s.QuoteAsset != ignore {
		targetAsset2 = s.QuoteAsset
	} else {
		targetAsset2 = s.BaseAsset
	}
	return targetAsset2
}

func ConvertPrice(bookTicker *BookTicker, basePrice float64, from string, to string, addFee bool) float64 {
	fee := .0
	if addFee {
		fee = constants.Fee
	}
	if strings.EqualFold(bookTicker.BaseAsset, from) && strings.EqualFold(bookTicker.QuoteAsset, to) { // SELL
		return bookTicker.BidPrice * basePrice * (1 - (fee / 100))
	} else {
		return (1 / bookTicker.AskPrice) * basePrice * (1 - (fee / 100))
	}
}
