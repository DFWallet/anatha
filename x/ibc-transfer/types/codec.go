package types

import (
	"github.com/DFWallet/anatha/codec"
	cdctypes "github.com/DFWallet/anatha/codec/types"
	sdk "github.com/DFWallet/anatha/types"
	channel "github.com/DFWallet/anatha/x/ibc/04-channel"
	commitmenttypes "github.com/DFWallet/anatha/x/ibc/23-commitment/types"
)

// RegisterCodec registers the IBC transfer types
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransfer{}, "cosmos-sdk/MsgTransfer", nil)
	cdc.RegisterConcrete(FungibleTokenPacketData{}, "cosmos-sdk/PacketDataTransfer", nil)
}

// RegisterInterfaces register the ibc transfer module interfaces to protobuf
// Any.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransfer{})
}

var (
	amino = codec.New()

	// ModuleCdc references the global x/transfer module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/transfer and
	// defined at the application level.
	ModuleCdc = codec.NewHybridCodec(amino, cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	channel.RegisterCodec(amino)
	commitmenttypes.RegisterCodec(amino)
	amino.Seal()
}
