package tendermint

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/codec"
	"github.com/DFWallet/anatha/x/ibc/07-tendermint/client/cli"
	"github.com/DFWallet/anatha/x/ibc/07-tendermint/client/rest"
	"github.com/DFWallet/anatha/x/ibc/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}

// RegisterRESTRoutes registers the REST routes for the IBC client
func RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router, queryRoute string) {
	rest.RegisterRoutes(ctx, rtr, fmt.Sprintf("%s/%s", queryRoute, types.SubModuleName))
}

// GetTxCmd returns the root tx command for the IBC client
func GetTxCmd(cdc *codec.Codec, storeKey string) *cobra.Command {
	return cli.GetTxCmd(cdc, fmt.Sprintf("%s/%s", storeKey, types.SubModuleName))
}
