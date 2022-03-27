package arb

import (
	"errors"
	"github.com/halm4d/arbitragecli/constants"
	"github.com/halm4d/arbitragecli/util"
	"math"
)

func AllCryptoCurrency() []string {
	//return []string{"BTC", "ETH", "GMT", "LUNA", "BNB", "AVAX", "XRP", "SOL", "WAVES", "JASMY", "SAND", "ADA", "FTM", "DOT", "GALA", "LINK", "ALPINE", "SHIB", "RUNE", "APE", "MANA", "ATOM", "LTC", "NEAR", "CAKE", "CELO", "DOGE", "MATIC", "XNO", "TRX", "ZEC", "SLP", "AAVE", "BCH", "SANTOS", "GRT", "XLM", "AXS", "SAND", "PEOPLE", "VET", "FIL", "ANC", "EGLD", "ALGO", "ENS", "NEO", "ETC", "UNI", "MINA", "ICP", "XMR", "ALICE", "LAZIO", "EOS", "ROSE", "CRV", "BAT", "AGLD", "SNX", "SLP", "DAI", "SUSHI", "XTZ", "ONE", "TLM", "LUNA", "IMX", "CHZ", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX", "MASK", "GLMR", "MBOX", "1INCH", "KNC", "WBTC", "TROY", "HBAR", "JOE", "REN", "CVC", "GRT", "PORTO", "LOKA", "VOXEL", "POLS", "LINA", "BTT", "MIR", "WIN", "WOO", "FTT", "C98", "QTUM", "ADX", "TWT", "PAXG", "BETA", "BICO", "HOT", "FLOW", "JST", "KSM", "DAR", "DCR", "KLAY", "HNT", "MKR", "SCRT", "SYS", "KDA", "SPELL", "INJ", "SUN", "ZEN", "COMP", "IDEX", "BTTC", "OG", "BAND", "KP3R", "MC", "DENT", "IOST", "ABBC", "PYR", "ARPA", "BAKE", "API3", "DUSK", "COCOS", "OXT", "CELR", "XEC", "REEF", "T", "ZRX", "BTCDOWN", "RARE", "IOT", "OOKI", "CTSI", "RNDR", "HIGH", "ANT", "POND", "ICX", "SRM", "XVS", "CVP", "UNFI", "MTL", "QNT", "BTCUP", "JUV", "ZIL", "VITE", "BAR", "CVX", "COTI", "KAVA", "ANKR", "MOVR", "TFUEL", "RVN", "MDX", "DODO", "ANY", "QI", "OCEAN", "FLUX", "SUPER", "YGG", "WAXP", "STX", "FET", "ORN", "DIA", "EPS", "STRAX", "TORN", "MCONTENT", "LOOM", "AMA", "HIVE", "HT", "BSV", "ONT", "WRX", "SFP", "TVK", "POWR", "MXC", "ACA", "GXS", "REQ", "LIT", "CITY", "LUNA", "CLV", "RAY", "AUDIO", "NKN", "STORJ", "FIDA", "NAV", "CTK", "WOO", "JASMY", "ASTR", "COS", "MDT", "GTC", "ATA", "ALPHA", "UTK", "BTCST", "BEL", "BAL", "MFT", "ASR", "ATM", "BNX", "OKB", "TKO", "WTC", "NEXO", "CKB", "CRO", "PLA", "SG", "STEEM", "CHESS", "SKL", "HARD", "DGB", "VIDT", "XVG", "SC", "WAN", "DEGO", "FRONT", "PSG", "NANO", "TT", "XEM", "HEX", "ILV", "LPT", "BURGER", "RAD", "BNT", "BLZ", "ELF", "FORTH", "FARM", "KEY", "RAMP", "BMX", "NMR", "WNXM", "AVA", "PERP", "TCT", "FIS", "YFII", "FOR", "IRIS", "ERN", "POLY", "ALCX", "DFA", "SEELE", "TRU", "NULS", "FLM", "CTXC", "AKRO", "BEAM", "ACH", "LSK"}
	return []string{"ETH", "BTC", "LTC", "BNB", "NEO", "QTUM", "EOS", "SNT", "BNT", "GAS", "USDT", "WTC", "LRC", "OMG", "ZRX", "KNC", "FUN", "SNM", "IOTA", "LINK", "XVG", "MDA", "MTL", "ETC", "DNT", "ZEC", "AST", "DASH", "OAX", "BTG", "REQ", "VIB", "TRX", "POWR", "ARK", "XRP", "ENJ", "STORJ", "KMD", "NULS", "XMR", "AMB", "BAT", "GXS", "QSP", "BTS", "LSK", "MANA", "ADX", "ADA", "XLM", "WABI", "WAVES", "GTO", "ICX", "ELF", "AION", "NEBL", "BRD", "NAV", "RLC", "PIVX", "IOST", "STEEM", "BLZ", "ZIL", "ONT", "XEM", "WAN", "QLC", "SYS", "GRS", "LOOM", "REP", "TUSD", "ZEN", "CVC", "THETA", "IOTX", "QKC", "NXS", "DATA", "SC", "KEY", "NAS", "MFT", "DENT", "ARDR", "HOT", "VET", "DOCK", "POLY", "GO", "RVN", "DCR", "MITH", "REN", "USDC", "ONG", "FET", "CELR", "MATIC", "ATOM", "PHB", "TFUEL", "ONE", "FTM", "ALGO", "DOGE", "DUSK", "ANKR", "WIN", "COS", "COCOS", "TOMO", "PERL", "CHZ", "BAND", "BUSD", "BEAM", "XTZ", "HBAR", "NKN", "STX", "KAVA", "NGN", "ARPA", "CTXC", "BCH", "RUB", "TROY", "VITE", "FTT", "TRY", "EUR", "OGN", "DREP", "TCT", "WRX", "LTO", "MBL", "COTI", "STPT", "SOL", "IDRT", "CTSI", "HIVE", "CHR", "MDT", "STMX", "IQ", "PNT", "GBP", "DGB", "UAH", "COMP", "BIDR", "SXP", "SNX", "VTHO", "IRIS", "MKR", "RUNE", "AUD", "FIO", "AVA", "BAL", "YFI", "DAI", "JST", "SRM", "ANT", "CRV", "SAND", "OCEAN", "NMR", "DOT", "LUNA", "IDEX", "RSR", "PAXG", "WNXM", "TRB", "WBTC", "SUSHI", "YFII", "KSM", "EGLD", "DIA", "UMA", "BEL", "WING", "CREAM", "UNI", "NBS", "OXT", "SUN", "AVAX", "HNT", "BAKE", "BURGER", "FLM", "SCRT", "CAKE", "SPARTA", "ORN", "UTK", "XVS", "ALPHA", "VIDT", "AAVE", "BRL", "NEAR", "FIL", "INJ", "AERGO", "AUDIO", "CTK", "AKRO", "KP3R", "AXS", "HARD", "RENBTC", "SLP", "CVP", "STRAX", "FOR", "UNFI", "FRONT", "ROSE", "HEGIC", "PROM", "SKL", "SUSD", "GLM", "GHST", "DF", "GRT", "JUV", "PSG", "1INCH", "REEF", "OG", "ATM", "ASR", "CELO", "RIF", "BTCST", "TRU", "DEXE", "CKB", "TWT", "FIRO", "BETH", "PROS", "LIT", "VAI", "SFP", "FXS", "DODO", "UFT", "ACM", "AUCTION", "PHA", "TVK", "BADGER", "FIS", "OM", "POND", "DEGO", "ALICE", "BIFI", "LINA", "PERP", "RAMP", "SUPER", "CFX", "EPS", "AUTO", "TKO", "PUNDIX", "TLM", "MIR", "BAR", "FORTH", "EZ", "SHIB", "ICP", "AR", "POLS", "MDX", "MASK", "LPT", "AGIX", "ATA", "GTC", "TORN", "ERN", "KLAY", "BOND", "MLN", "QUICK", "C98", "CLV", "QNT", "FLOW", "XEC", "MINA", "RAY", "FARM", "ALPACA", "MBOX", "VGX", "WAXP", "TRIBE", "GNO", "DYDX", "USDP", "GALA", "ILV", "YGG", "FIDA", "AGLD", "RAD", "BETA", "RARE", "SSV", "LAZIO", "CHESS", "DAR", "BNX", "MOVR", "CITY", "ENS", "QI", "PORTO", "JASMY", "AMP", "PLA", "PYR", "RNDR", "ALCX", "SANTOS", "MC", "ANY", "BICO", "FLUX", "VOXEL", "HIGH", "CVX", "PEOPLE", "OOKI", "SPELL", "UST", "JOE", "ACH", "IMX", "GLMR", "LOKA", "API3", "BTTC", "ACA", "ANC", "BDOT", "XNO", "WOO", "ALPINE", "T", "ASTR", "NBT", "GMT", "KDA", "APE"}
	//return []string{"BUSD", "USDT", "NBT", "BIDR", "BTC", "ETH", "SOL", "LTC", "DOGE", "MATIC", "BCH", "SANTOS", "GRT", "XLM", "AXS", "ANC", "EGLD", "ALGO", "MINA", "ICP", "XMR", "ALICE", "BAT", "AGLD", "SNX", "ONE", "TLM", "LUNA", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX"}
	//return []string{"BNB", "BTC", "COCOS"}
}

type Symbols struct {
	CS            SymbolsMap
	US            SymbolsMap
	AssetMap      AssetMap
	BaseAssetMap  AssetMap
	QuoteAssetMap AssetMap
}

type AssetMap map[string][]Symbol
type SymbolsMap map[string]Symbol

type Symbol struct {
	Symbol     string
	BaseAsset  string
	QuoteAsset string
	BidPrice   float64
	AskPrice   float64
}

func NewSymbols() *Symbols {
	return &Symbols{
		US:            make(map[string]Symbol),
		CS:            make(map[string]Symbol),
		AssetMap:      make(map[string][]Symbol),
		BaseAssetMap:  make(map[string][]Symbol),
		QuoteAssetMap: make(map[string][]Symbol),
	}
}

func (s *Symbols) Init(exchangeResp *ExchangeInfoResp) {
	for _, exchange := range exchangeResp.Symbols {
		if exchange.Status != "TRADING" {
			continue
		}
		symbol := Symbol{
			Symbol:     exchange.Symbol,
			BaseAsset:  exchange.BaseAsset,
			QuoteAsset: exchange.QuoteAsset,
		}
		if exchange.QuoteAsset == constants.USDT || exchange.BaseAsset == constants.USDT {
			s.US[exchange.Symbol] = symbol
		}
		if util.Contains(AllCryptoCurrency(), exchange.QuoteAsset) && util.Contains(AllCryptoCurrency(), exchange.BaseAsset) {
			s.CS[exchange.Symbol] = symbol
		}
		s.AssetMap[exchange.QuoteAsset] = append(s.AssetMap[exchange.QuoteAsset], symbol)
		s.AssetMap[exchange.BaseAsset] = append(s.AssetMap[exchange.BaseAsset], symbol)
		s.QuoteAssetMap[exchange.QuoteAsset] = append(s.QuoteAssetMap[exchange.QuoteAsset], symbol)
		s.BaseAssetMap[exchange.BaseAsset] = append(s.BaseAssetMap[exchange.BaseAsset], symbol)
	}
}

func (s *SymbolsMap) findAllSymbolByAsset(asset string) *SymbolsMap {
	var symbols = make(SymbolsMap)
	for k, v := range *s {
		if v.BaseAsset == asset || v.QuoteAsset == asset {
			symbols[k] = v
		}
	}
	return &symbols
}

func (s *SymbolsMap) findByAssetPair(asset1 string, asset2 string) (*Symbol, error) {
	for _, symbol := range *s {
		if (symbol.BaseAsset == asset1 || symbol.QuoteAsset == asset1) && (symbol.BaseAsset == asset2 || symbol.QuoteAsset == asset2) {
			return &symbol, nil
		}
	}
	return &Symbol{}, errors.New("symbol not found")
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

func (s *SymbolsMap) calculateArbsForSymbol(startEndAsset string) *Arbitrages {
	var arbs Arbitrages
	for _, asset1 := range *s.findAllSymbolByAsset(startEndAsset) {
		var targetAsset1 = asset1.getTargetAsset(startEndAsset)
		if asset1.BaseAsset == constants.USDT || asset1.QuoteAsset == constants.USDT {
			continue
		}
		for _, asset2 := range *s.findAllSymbolByAsset(targetAsset1) {
			if asset2.BaseAsset == constants.USDT || asset2.QuoteAsset == constants.USDT {
				continue
			}
			targetAsset2 := asset2.getTargetAsset(targetAsset1)
			if targetAsset2 == startEndAsset {
				continue
			}
			pair, err := s.findByAssetPair(targetAsset2, startEndAsset)
			if err != nil {
				continue
			}
			trades := Arbitrage{
				Trades: []Trade{
					{
						From:   startEndAsset,
						To:     targetAsset1,
						Symbol: asset1.Symbol,
					},
					{
						From:   targetAsset1,
						To:     targetAsset2,
						Symbol: asset2.Symbol,
					},
					{
						From:   targetAsset2,
						To:     startEndAsset,
						Symbol: pair.Symbol,
					},
				},
				Profit: -math.MaxFloat64,
			}
			arbs = append(arbs, trades)
		}
	}
	return &arbs
}
