package simulation

import (
	"math/rand"

	simtypes "github.com/DFWallet/anatha/types/simulation"
	"github.com/DFWallet/anatha/x/ibc/04-channel/types"
)

// GenChannelGenesis returns the default channel genesis state.
func GenChannelGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
