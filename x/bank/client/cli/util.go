package cli

import (
	"fmt"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/bank/types"
)

func queryTotalSupply(cliCtx context.CLIContext, cdc *codec.Codec) error {
	params := types.NewQueryTotalSupplyParams(1, 0) // no pagination
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return err
	}

	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalSupply), bz)
	if err != nil {
		return err
	}

	var totalSupply sdk.Coins
	err = cdc.UnmarshalJSON(res, &totalSupply)
	if err != nil {
		return err
	}

	return cliCtx.PrintOutput(totalSupply)
}

func querySupplyOf(cliCtx context.CLIContext, cdc *codec.Codec, denom string) error {
	params := types.NewQuerySupplyOfParams(denom)
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return err
	}

	res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QuerySupplyOf), bz)
	if err != nil {
		return err
	}

	var supply sdk.Int
	err = cdc.UnmarshalJSON(res, &supply)
	if err != nil {
		return err
	}

	return cliCtx.PrintOutput(supply)
}
