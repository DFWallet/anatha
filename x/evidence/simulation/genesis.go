package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/DFWallet/anatha/codec"

	"github.com/DFWallet/anatha/types/module"
	simtypes "github.com/DFWallet/anatha/types/simulation"
	"github.com/DFWallet/anatha/x/evidence/exported"
	"github.com/DFWallet/anatha/x/evidence/types"
)

// Simulation parameter constants
const evidence = "evidence"

// GenEvidences returns an empty slice of evidences.
func GenEvidences(_ *rand.Rand, _ []simtypes.Account) []exported.Evidence {
	return []exported.Evidence{}
}

// RandomizedGenState generates a random GenesisState for evidence
func RandomizedGenState(simState *module.SimulationState) {
	var ev []exported.Evidence

	simState.AppParams.GetOrGenerate(
		simState.Cdc, evidence, &ev, simState.Rand,
		func(r *rand.Rand) { ev = GenEvidences(r, simState.Accounts) },
	)

	evidenceGenesis := types.GenesisState{Evidence: ev}

	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, codec.MustMarshalJSONIndent(simState.Cdc, evidenceGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(evidenceGenesis)
}
