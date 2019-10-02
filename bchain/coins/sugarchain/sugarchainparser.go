package sugarchain

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0x9d4beb9f
	TestnetMagic wire.BitcoinNet = 0x709011b0
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{63}
	MainNetParams.ScriptHashAddrID = []byte{125}
	MainNetParams.Bech32HRPSegwit = "sugar"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{66}
	TestNetParams.ScriptHashAddrID = []byte{128}
	TestNetParams.Bech32HRPSegwit = "tugar"
}

// SugarchainParser handle
type SugarchainParser struct {
	*btc.BitcoinParser
}

// NewSugarchainParser returns new SugarchainParser instance
func NewSugarchainParser(params *chaincfg.Params, c *btc.Configuration) *SugarchainParser {
	return &SugarchainParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Sugarchain network,
// and the test Sugarchain network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
