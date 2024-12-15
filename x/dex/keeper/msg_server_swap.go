package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

  poolId := types.GeneratePoolId(msg.Token0, msg.Token1)
  pool, found := k.GetLiquidityPool(ctx, poolId)
  if !found {
		return nil, errorsmod.Wrapf(types.ErrPoolNonExistant, "pool %s does not exist", poolId)
  }

  t0, t1, a0, a1 := types.OrderTokensAndAmounts(
    msg.Token0,
    msg.Token1,
    msg.Amount0,
    msg.Amount1,
  )

  senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return nil, err
  }
  out0, out1 := types.CalculateSwapAmount(pool.K, pool.Amount0, pool.Amount1, a0, a1)

  coinsIn, err := types.CreateLPCoins(t0, t1, a0, a1)
  if err != nil {
    return nil, err
  }

  coinsOut, err := types.CreateLPCoins(t0, t1, out0, out1)
  if err != nil {
    return nil, err
  }

  err = k.ExecuteSwap(ctx, senderAddr, coinsIn, coinsOut)
  if err != nil {
    return nil, err
  }

  pool.Amount0 = pool.Amount0 + a0 - out0
  pool.Amount1 = pool.Amount1 + a1 - out1
  k.SetLiquidityPool(ctx, pool)

	return &types.MsgSwapResponse{}, nil
}
