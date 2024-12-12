package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

  poolId, t0, t1, a0, a1 := types.GeneratePoolId(
    msg.Token0,
    msg.Token1,
    msg.Amount0,
    msg.Amount1,
  )
  _, found := k.Keeper.GetLiquidityPool(ctx, poolId)
  if found {
    return nil, errorsmod.Wrapf(types.ErrPoolExists, "pool %s exists", poolId)
  }

  pool := types.LiquidityPool{
    Index: poolId,
    Token0: t0,
    Token1: t1,
    Amount0: a0,
    Amount1: a1,
  }
  k.Keeper.SetLiquidityPool(ctx, pool)

  senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
  coins0 := sdk.NewCoin(msg.Token0, math.NewIntFromUint64(msg.Amount0))
  coins1 := sdk.NewCoin(msg.Token1, math.NewIntFromUint64(msg.Amount1))

  err := k.bankKeeper.SendCoinsFromAccountToModule(
    ctx,
    senderAddr,
    types.ModuleName,
    sdk.Coins{coins0, coins1},
  )

  if err != nil {
    return nil, errorsmod.Wrapf(types.ErrProvidingLiquidity, "error: %s", err)
  }

	return &types.MsgCreatePoolResponse{}, nil
}