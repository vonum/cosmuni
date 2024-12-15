package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

  pool, found := k.GetLiquidityPool(ctx, msg.PoolId)
  if !found {
		return nil, errorsmod.Wrapf(types.ErrPoolNonExistant, "pool %s does not exist", msg.PoolId)
  }
  poolDenom := types.PoolDenom(msg.PoolId)

  senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    return nil, err
  }

  senderShares := k.bankKeeper.SpendableCoin(ctx, senderAddr, poolDenom)
  if msg.Shares > senderShares.Amount.Uint64() {
    return nil, errorsmod.Wrapf(
      types.ErrWitdhrawingLiquidity,
      "%s has %d shares avaliable",
      senderAddr,
      senderShares.Amount.Uint64(),
    )
  }

  shareRatio := types.CalculateSharesPercentage(msg.Shares, pool.TotalShares)
  a0 := uint64(float64(pool.Amount0) * shareRatio)
  a1 := uint64(float64(pool.Amount1) * shareRatio)

  lpCoins, err := types.CreateLPCoins(pool.Token0, pool.Token1, a0, a1)
  if err != nil {
    return nil, err
  }

  shares, err := types.CreateSharesCoins(msg.PoolId, msg.Shares)
  if err != nil {
    return nil, err
  }

  err = k.ExecuteWithdrawal(ctx, senderAddr, lpCoins, shares)
  if err != nil {
    return nil, err
  }

  pool.TotalShares -= msg.Shares
  pool.Amount0 -= a0
  pool.Amount1 -= a1
  k.SetLiquidityPool(ctx, pool)

	return &types.MsgWithdrawResponse{}, nil
}
