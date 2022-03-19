package util

func AllCryptoCurrency() []string {
	return []string{"BTC", "ETH", "GMT", "LUNA", "BNB", "AVAX", "XRP", "SOL", "WAVES", "JASMY", "SAND", "ADA", "FTM", "DOT", "GALA", "LINK", "ALPINE", "SHIB", "RUNE", "APE", "MANA", "ATOM", "LTC", "NEAR", "CAKE", "CELO", "DOGE", "MATIC", "XNO", "TRX", "ZEC", "SLP", "AAVE", "BCH", "SANTOS", "GRT", "XLM", "AXS", "SAND", "PEOPLE", "VET", "FIL", "ANC", "EGLD", "ALGO", "ENS", "NEO", "ETC", "UNI", "MINA", "ICP", "XMR", "ALICE", "LAZIO", "EOS", "ROSE", "CRV", "BAT", "AGLD", "SNX", "SLP", "DAI", "SUSHI", "XTZ", "ONE", "TLM", "LUNA", "IMX", "CHZ", "ENJ", "DASH", "OGN", "THETA", "GALA", "YFI", "MBL", "AR", "SXP", "RSR", "CGLD", "LUNA", "DYDX", "LRC", "IOTX", "USDP", "CHR", "UMA", "OMG", "ADX", "MASK", "GLMR", "MBOX", "1INCH", "KNC", "WBTC", "TROY", "HBAR", "JOE", "REN", "CVC", "GRT", "PORTO", "LOKA", "VOXEL", "POLS", "LINA", "BTT", "MIR", "WIN", "WOO", "FTT", "C98", "QTUM", "ADX", "TWT", "PAXG", "BETA", "BICO", "HOT", "FLOW", "JST", "KSM", "DAR", "DCR", "KLAY", "HNT", "MKR", "SCRT", "SYS", "KDA", "SPELL", "INJ", "SUN", "ZEN", "COMP", "IDEX", "BTTC", "OG", "BAND", "KP3R", "MC", "DENT", "IOST", "ABBC", "PYR", "ARPA", "BAKE", "API3", "DUSK", "COCOS", "OXT", "CELR", "XEC", "REEF", "T", "ZRX", "BTCDOWN", "RARE", "IOT", "OOKI", "CTSI", "RNDR", "HIGH", "ANT", "POND", "ICX", "SRM", "XVS", "CVP", "UNFI", "MTL", "QNT", "BTCUP", "JUV", "ZIL", "VITE", "BAR", "CVX", "COTI", "KAVA", "ANKR", "MOVR", "TFUEL", "RVN", "MDX", "DODO", "ANY", "QI", "OCEAN", "FLUX", "SUPER", "YGG", "WAXP", "STX", "FET", "ORN", "DIA", "EPS", "STRAX", "TORN", "MCONTENT", "LOOM", "AMA", "HIVE", "HT", "BSV", "ONT", "WRX", "SFP", "TVK", "POWR", "MXC", "ACA", "GXS", "REQ", "LIT", "CITY", "LUNA", "CLV", "RAY", "AUDIO", "NKN", "STORJ", "FIDA", "NAV", "CTK", "WOO", "JASMY", "ASTR", "COS", "MDT", "GTC", "ATA", "ALPHA", "UTK", "BTCST", "BEL", "BAL", "MFT", "ASR", "ATM", "BNX", "OKB", "TKO", "WTC", "NEXO", "CKB", "CRO", "PLA", "SG", "STEEM", "CHESS", "SKL", "HARD", "DGB", "VIDT", "XVG", "SC", "WAN", "DEGO", "FRONT", "PSG", "NANO", "TT", "XEM", "HEX", "ILV", "LPT", "BURGER", "RAD", "BNT", "BLZ", "ELF", "FORTH", "FARM", "KEY", "RAMP", "BMX", "NMR", "WNXM", "AVA", "PERP", "TCT", "FIS", "YFII", "FOR", "IRIS", "ERN", "POLY", "ALCX", "DFA", "SEELE", "TRU", "NULS", "FLM", "CTXC", "AKRO", "BEAM", "ACH", "LSK"}
}

func Contains(slice []string, t string) bool {
	for _, s := range slice {
		if s == t {
			return true
		}
	}
	return false
}
