package types

import (
	"github.com/DFWallet/anatha/codec"
)

var (
	amino = codec.New()
)

func init() {
	codec.RegisterCrypto(amino)
	amino.Seal()
}
