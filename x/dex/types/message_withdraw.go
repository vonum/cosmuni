package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWithdraw{}

func NewMsgWithdraw(creator string, poolId string, shares uint64) *MsgWithdraw {
	return &MsgWithdraw{
		Creator: creator,
		PoolId:  poolId,
		Shares:  shares,
	}
}

func (msg *MsgWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Shares <= 0 {
		return errorsmod.Wrapf(ErrInvalidTokenAmount, "invalid shares amount %d", msg.Shares)
	}

	return nil
}
