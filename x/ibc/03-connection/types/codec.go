package types

import (
	"github.com/DFWallet/anatha/codec"
	cdctypes "github.com/DFWallet/anatha/codec/types"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/x/ibc/03-connection/exported"
)

// RegisterCodec registers the necessary x/ibc/03-connection interfaces and concrete types
// on the provided Amino codec. These types are used for Amino JSON serialization.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterInterface((*exported.ConnectionI)(nil), nil)
	cdc.RegisterInterface((*exported.CounterpartyI)(nil), nil)

	cdc.RegisterConcrete(ConnectionEnd{}, "ibc/connection/ConnectionEnd", nil)
	cdc.RegisterConcrete(MsgConnectionOpenInit{}, "ibc/connection/MsgConnectionOpenInit", nil)
	cdc.RegisterConcrete(MsgConnectionOpenTry{}, "ibc/connection/MsgConnectionOpenTry", nil)
	cdc.RegisterConcrete(MsgConnectionOpenAck{}, "ibc/connection/MsgConnectionOpenAck", nil)
	cdc.RegisterConcrete(MsgConnectionOpenConfirm{}, "ibc/connection/MsgConnectionOpenConfirm", nil)
}

// RegisterInterfaces register the ibc interfaces submodule implementations to protobuf
// Any.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConnectionOpenInit{},
		&MsgConnectionOpenTry{},
		&MsgConnectionOpenAck{},
		&MsgConnectionOpenConfirm{},
	)
}

var (
	amino = codec.New()

	// SubModuleCdc references the global x/ibc/03-connectionl module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/ibc/03-connectionl and
	// defined at the application level.
	SubModuleCdc = codec.NewHybridCodec(amino, cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}