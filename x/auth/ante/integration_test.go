package ante_test

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/simapp"
	sdk "github.com/DFWallet/anatha/types"
	authtypes "github.com/DFWallet/anatha/x/auth/types"
)

// returns context and app with params set on account keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, abci.Header{})
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	return app, ctx
}
