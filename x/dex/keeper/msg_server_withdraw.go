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
  senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
  senderShares := k.bankKeeper.SpendableCoin(ctx, senderAddr, poolDenom)

  if msg.Shares > senderShares.Amount.Uint64() {
    return nil, errorsmod.Wrapf(types.ErrInsufficientShhares, "%s has %d shares avaliable", senderAddr, senderShares.Amount.Uint64())
  }

  shareRatio := types.CalculateSharesPercentage(msg.Shares, pool.TotalShares)
  a0 := uint64(float64(pool.Amount0) * shareRatio)
  a1 := uint64(float64(pool.Amount1) * shareRatio)

	lpCoins, _ := sdk.ParseCoinsNormalized(
    types.FormatCoinsStr(pool.Token0, pool.Token1, a0, a1),
  )
  err := k.bankKeeper.SendCoinsFromModuleToAccount(
    ctx,
    types.ModuleName,
    senderAddr,
    lpCoins,
  )
  if err != nil {
    return nil, err
  }

	shareCoins, err := sdk.ParseCoinsNormalized(
    types.FormatShareCoinStr(msg.PoolId, msg.Shares),
  )
  err = k.bankKeeper.SendCoinsFromAccountToModule(
    ctx,
    senderAddr,
    types.ModuleName,
    shareCoins,
  )
  if err != nil {
    return nil, err
  }

  err =  k.bankKeeper.BurnCoins(ctx, types.ModuleName, shareCoins)
  if err != nil {
    return nil, err
  }

  pool.TotalShares -= msg.Shares
  pool.Amount0 -= a0
  pool.Amount1 -= a1
  k.SetLiquidityPool(ctx, pool)

	return &types.MsgWithdrawResponse{}, nil
}
