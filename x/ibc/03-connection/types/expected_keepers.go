package types

import (
	sdk "github.com/DFWallet/anatha/types"
	clientexported "github.com/DFWallet/anatha/x/ibc/02-client/exported"
)

// ClientKeeper expected account IBC client keeper
type ClientKeeper interface {
	GetClientState(ctx sdk.Context, clientID string) (clientexported.ClientState, bool)
	GetClientConsensusState(ctx sdk.Context, clientID string, height uint64) (clientexported.ConsensusState, bool)
	GetSelfConsensusState(ctx sdk.Context, height uint64) (clientexported.ConsensusState, bool)
	IterateClients(ctx sdk.Context, cb func(clientexported.ClientState) bool)
	ClientStore(ctx sdk.Context, clientID string) sdk.KVStore
}
