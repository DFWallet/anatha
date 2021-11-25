package types

import (
	sdkerrors "github.com/DFWallet/anatha/types/errors"
)

// IBC channel sentinel errors
var (
	ErrInvalidPacketTimeout    = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrOnlyOneDenomAllowed     = sdkerrors.Register(ModuleName, 3, "only one denom allowed")
	ErrInvalidDenomForTransfer = sdkerrors.Register(ModuleName, 4, "invalid denomination for cross-chain transfer")
)