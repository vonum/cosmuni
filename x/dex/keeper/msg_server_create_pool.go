package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poolId := types.GeneratePoolId(msg.Token0, msg.Token1)
	_, found := k.Keeper.GetLiquidityPool(ctx, poolId)
	if found {
		return nil, errorsmod.Wrapf(types.ErrPoolExists, "pool %s exists", poolId)
	}

	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return nil, err
  }

	t0, t1, a0, a1 := types.OrderTokensAndAmounts(
		msg.Token0,
		msg.Token1,
		msg.Amount0,
		msg.Amount1,
	)

	sharesAmount := types.CalculateShares(a0, a1, 0)
	pool := types.LiquidityPool{
		Index:       poolId,
		Token0:      t0,
		Token1:      t1,
		Amount0:     a0,
		Amount1:     a1,
		TotalShares: sharesAmount,
		K:           types.CalculateK(a0, a1),
	}
	k.Keeper.SetLiquidityPool(ctx, pool)

  lpCoins, err := types.CreateLPCoins(t0, t1, a0, a1)
  if err != nil {
    return nil, err
  }

  shares, err := types.CreateSharesCoins(poolId, sharesAmount)
  if err != nil {
    return nil, err
  }

  err = k.ExecuteDeposit(ctx, senderAddr, lpCoins, shares)
  if err != nil {
    return nil, err
  }

	return &types.MsgCreatePoolResponse{}, nil
}
