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
  t0, t1, a0, a1 := types.OrderTokensAndAmounts(
    msg.Token0,
    msg.Token1,
    msg.Amount0,
    msg.Amount1,
  )

  pool, found := k.GetLiquidityPool(ctx, poolId)
  if !found {
		return nil, errorsmod.Wrapf(types.ErrPoolNonExistant, "pool %s does not exist", poolId)
  }

  senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)

  out0, out1 := types.CalculateSwapAmount(pool.K, pool.Amount0, pool.Amount1, a0, a1)

  // send tokens to pool
  coinsIn, _ := sdk.ParseCoinsNormalized(types.FormatCoinsStr(t0, t1, a0, a1))
  err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, types.ModuleName, coinsIn)
  if err != nil {
    return nil, err
  }

  // send tokens from pool
  coinsOut, _ := sdk.ParseCoinsNormalized(types.FormatCoinsStr(t0, t1, out0, out1))
  err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, senderAddr, coinsOut)
  if err != nil {
    return nil, err
  }

  // update pool values
  pool.Amount0 = pool.Amount0 + a0 - out0
  pool.Amount1 = pool.Amount1 + a1 - out1
  k.SetLiquidityPool(ctx, pool)

	return &types.MsgSwapResponse{}, nil
}
