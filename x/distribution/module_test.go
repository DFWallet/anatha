package distribution_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/simapp"
	"github.com/DFWallet/anatha/x/auth"
	"github.com/DFWallet/anatha/x/distribution"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, types.Header{})

	app.InitChain(
		types.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, auth.NewModuleAddress(distribution.ModuleName))
	require.NotNil(t, acc)
}
