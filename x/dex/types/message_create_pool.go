package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, token0 string, token1 string, amount0 uint64, amount1 uint64) *MsgCreatePool {
	return &MsgCreatePool{
		Creator: creator,
		Token0:  token0,
		Token1:  token1,
		Amount0: amount0,
		Amount1: amount1,
	}
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Token0 == msg.Token1 {
		return errorsmod.Wrapf(ErrSameTokenPool, "token provided twice %s", msg.Token0)
	}

	if err := sdk.ValidateDenom(msg.Token0); err != nil {
		return errorsmod.Wrapf(ErrInvalidToken, "invalid token0 %s", msg.Token0)
	}

	if err := sdk.ValidateDenom(msg.Token1); err != nil {
		return errorsmod.Wrapf(ErrInvalidToken, "invalid token1 %s", msg.Token1)
	}

	if msg.Amount0 <= 0 {
		return errorsmod.Wrapf(ErrInvalidTokenAmount, "invalid token0 amount %d", msg.Amount0)
	}

	if msg.Amount1 <= 0 {
		return errorsmod.Wrapf(ErrInvalidTokenAmount, "invalid token1 amount %d", msg.Amount1)
	}

	return nil
}
