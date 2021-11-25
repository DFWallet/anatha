package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/anatha/simapp"
	"github.com/DFWallet/anatha/x/slashing/keeper"
	"github.com/DFWallet/anatha/x/slashing/types"
)

func TestNewQuerier(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, abci.Header{})
	app.SlashingKeeper.SetParams(ctx, keeper.TestParams())

	querier := keeper.NewQuerier(app.SlashingKeeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)
}

func TestQueryParams(t *testing.T) {
	cdc := codec.New()
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, abci.Header{})
	app.SlashingKeeper.SetParams(ctx, keeper.TestParams())

	querier := keeper.NewQuerier(app.SlashingKeeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	var params types.Params

	res, err := querier(ctx, []string{types.QueryParameters}, query)
	require.NoError(t, err)

	err = cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)
	require.Equal(t, app.SlashingKeeper.GetParams(ctx), params)
}
