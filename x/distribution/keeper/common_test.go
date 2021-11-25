package keeper_test

import (
	"github.com/DFWallet/anatha/simapp"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/distribution/types"
)

var (
	PKS = simapp.CreateTestPubKeys(5)

	valConsPk1 = PKS[0]
	valConsPk2 = PKS[1]
	valConsPk3 = PKS[2]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())

	distrAcc = auth.NewEmptyModuleAccount(types.ModuleName)
)
