package simapp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	tmkv "github.com/tendermint/tendermint/libs/kv"

	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/auth"
)

func TestGetSimulationLog(t *testing.T) {
	cdc := MakeCodec()

	decoders := make(sdk.StoreDecoderRegistry)
	decoders[auth.StoreKey] = func(cdc *codec.Codec, kvAs, kvBs tmkv.Pair) string { return "10" }

	tests := []struct {
		store       string
		kvPairs     []tmkv.Pair
		expectedLog string
	}{
		{
			"Empty",
			[]tmkv.Pair{{}},
			"",
		},
		{
			auth.StoreKey,
			[]tmkv.Pair{{Key: auth.GlobalAccountNumberKey, Value: cdc.MustMarshalBinaryLengthPrefixed(uint64(10))}},
			"10",
		},
		{
			"OtherStore",
			[]tmkv.Pair{{Key: []byte("key"), Value: []byte("value")}},
			fmt.Sprintf("store A %X => %X\nstore B %X => %X\n", []byte("key"), []byte("value"), []byte("key"), []byte("value")),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.store, func(t *testing.T) {
			require.Equal(t, tt.expectedLog, GetSimulationLog(tt.store, decoders, cdc, tt.kvPairs, tt.kvPairs), tt.store)
		})
	}
}
