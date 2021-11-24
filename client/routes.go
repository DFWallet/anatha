package client

import (
	"github.com/gorilla/mux"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
