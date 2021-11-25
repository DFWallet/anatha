package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/client"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	sdkerrors "github.com/DFWallet/anatha/types/errors"
	"github.com/DFWallet/anatha/x/ibc/02-client/exported"
	"github.com/DFWallet/anatha/x/ibc/02-client/types"
)

// QuerierClients defines the sdk.Querier to query all the light client states.
func QuerierClients(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	var params types.QueryAllClientsParams

	if err := k.cdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	clients := k.GetAllClients(ctx)

	start, end := client.Paginate(len(clients), params.Page, params.Limit, 100)
	if start < 0 || end < 0 {
		clients = []exported.ClientState{}
	} else {
		clients = clients[start:end]
	}

	res, err := codec.MarshalJSONIndent(k.cdc, clients)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
