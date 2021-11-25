package simulation

import (
	"math/rand"

	simtypes "github.com/DFWallet/anatha/types/simulation"
	"github.com/DFWallet/anatha/x/ibc/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
