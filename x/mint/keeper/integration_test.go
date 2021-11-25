package keeper_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/simapp"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/mint/types"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.MintKeeper.SetParams(ctx, types.DefaultParams())
	app.MintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}
