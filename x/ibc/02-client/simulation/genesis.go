package simulation

import (
	"math/rand"

	simtypes "github.com/DFWallet/anatha/types/simulation"
	"github.com/DFWallet/anatha/x/ibc/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
