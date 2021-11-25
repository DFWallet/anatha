package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/DFWallet/anatha/codec"

	"github.com/DFWallet/anatha/types/module"
	clientsims "github.com/DFWallet/anatha/x/ibc/02-client/simulation"
	clienttypes "github.com/DFWallet/anatha/x/ibc/02-client/types"
	connectionsims "github.com/DFWallet/anatha/x/ibc/03-connection/simulation"
	connectiontypes "github.com/DFWallet/anatha/x/ibc/03-connection/types"
	channelsims "github.com/DFWallet/anatha/x/ibc/04-channel/simulation"
	channeltypes "github.com/DFWallet/anatha/x/ibc/04-channel/types"
	host "github.com/DFWallet/anatha/x/ibc/24-host"
	"github.com/DFWallet/anatha/x/ibc/types"
)

// Simulation parameter constants
const (
	clientGenesis     = "client_genesis"
	connectionGenesis = "connection_genesis"
	channelGenesis    = "channel_genesis"
)

// RandomizedGenState generates a random GenesisState for evidence
func RandomizedGenState(simState *module.SimulationState) {
	var (
		clientGenesisState     clienttypes.GenesisState
		connectionGenesisState connectiontypes.GenesisState
		channelGenesisState    channeltypes.GenesisState
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, clientGenesis, &clientGenesisState, simState.Rand,
		func(r *rand.Rand) { clientGenesisState = clientsims.GenClientGenesis(r, simState.Accounts) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, connectionGenesis, &connectionGenesisState, simState.Rand,
		func(r *rand.Rand) { connectionGenesisState = connectionsims.GenConnectionGenesis(r, simState.Accounts) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, channelGenesis, &channelGenesisState, simState.Rand,
		func(r *rand.Rand) { channelGenesisState = channelsims.GenChannelGenesis(r, simState.Accounts) },
	)

	ibcGenesis := types.GenesisState{
		ClientGenesis:     clientGenesisState,
		ConnectionGenesis: connectionGenesisState,
		ChannelGenesis:    channelGenesisState,
	}

	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", host.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, ibcGenesis))
	simState.GenState[host.ModuleName] = simState.Cdc.MustMarshalJSON(ibcGenesis)
}