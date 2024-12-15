package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/dex module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSameTokenPool        = sdkerrors.Register(ModuleName, 1101, "same token in pool")
	ErrInvalidToken         = sdkerrors.Register(ModuleName, 1102, "invalid token")
	ErrInvalidTokenAmount   = sdkerrors.Register(ModuleName, 1103, "invalid token amount")

	ErrPoolExists           = sdkerrors.Register(ModuleName, 1104, "pool exists")
	ErrPoolNonExistant      = sdkerrors.Register(ModuleName, 1105, "pool doesn't exists")

	ErrProvidingLiquidity   = sdkerrors.Register(ModuleName, 1106, "failed to provide liquidity")
	ErrWitdhrawingLiquidity = sdkerrors.Register(ModuleName, 1107, "failed to provide liquidity")

	ErrSwapping             = sdkerrors.Register(ModuleName, 1108, "insufficient shares")
)
