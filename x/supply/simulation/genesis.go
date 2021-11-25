package simulation

// DONTCOVER

import (
	"fmt"

	"github.com/DFWallet/anatha/codec"

	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/types/module"
	"github.com/DFWallet/anatha/x/supply/internal/types"
)

// RandomizedGenState generates a random GenesisState for supply
func RandomizedGenState(simState *module.SimulationState) {
	numAccs := int64(len(simState.Accounts))
	totalSupply := sdk.NewInt(simState.InitialStake * (numAccs + simState.NumBonded))
	supplyGenesis := types.NewGenesisState(sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, totalSupply)))

	fmt.Printf("Generated supply parameters:\n%s\n", codec.MustMarshalJSONIndent(simState.Cdc, supplyGenesis))
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(supplyGenesis)
}
