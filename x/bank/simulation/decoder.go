package simulation

import (
	"bytes"
	"fmt"

	"github.com/DFWallet/anatha/x/bank/exported"

	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/DFWallet/anatha/x/bank/types"
)

type SupplyUnmarshaller interface {
	UnmarshalSupply([]byte) (exported.SupplyI, error)
}

// NewDecodeStore returns a function closure that unmarshals the KVPair's values
// to the corresponding types.
func NewDecodeStore(cdc SupplyUnmarshaller) func(kvA, kvB tmkv.Pair) string {
	return func(kvA, kvB tmkv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.SupplyKey):
			supplyA, err := cdc.UnmarshalSupply(kvA.Value)
			if err != nil {
				panic(err)
			}

			supplyB, err := cdc.UnmarshalSupply(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("%v\n%v", supplyA, supplyB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}