package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/dex module sentinel errors
var (
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample             = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrSameTokenPool      = sdkerrors.Register(ModuleName, 1102, "same token in pool")
	ErrInvalidToken       = sdkerrors.Register(ModuleName, 1103, "invalid token")
	ErrInvalidTokenAmount = sdkerrors.Register(ModuleName, 1104, "invalid token amount")

	ErrPoolExists         = sdkerrors.Register(ModuleName, 1105, "pool exists")
	ErrProvidingLiquidity = sdkerrors.Register(ModuleName, 1106, "failed to provide liquidity")
	ErrMintingShares      = sdkerrors.Register(ModuleName, 1107, "failed to mint shares")
	ErrTransferingShares  = sdkerrors.Register(ModuleName, 1108, "failed to transfer shares")
)
