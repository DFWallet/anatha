package types_test

import (
	"github.com/DFWallet/anatha/simapp"
)

var (
	app         = simapp.Setup(false)
	appCodec, _ = simapp.MakeCodecs()
)
