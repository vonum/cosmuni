package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeposit{}

func NewMsgDeposit(creator string, token0 string, token1 string, amount0 uint64, amount1 uint64) *MsgDeposit {
	return &MsgDeposit{
		Creator: creator,
		Token0:  token0,
		Token1:  token1,
		Amount0: amount0,
		Amount1: amount1,
	}
}

func (msg *MsgDeposit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
