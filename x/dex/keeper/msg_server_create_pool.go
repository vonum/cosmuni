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
	t0, t1, a0, a1 := types.OrderTokensAndAmounts(
		msg.Token0,
		msg.Token1,
		msg.Amount0,
		msg.Amount1,
	)
	_, found := k.Keeper.GetLiquidityPool(ctx, poolId)
	if found {
		return nil, errorsmod.Wrapf(types.ErrPoolExists, "pool %s exists", poolId)
	}

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

	senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	lpCoins, _ := sdk.ParseCoinsNormalized(types.FormatCoinsStr(t0, t1, a0, a1))

	err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		senderAddr,
		types.ModuleName,
		lpCoins,
	)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrProvidingLiquidity, "error: %s", err)
	}

	shares, err := sdk.ParseCoinsNormalized(
    types.FormatShareCoinStr(poolId, sharesAmount),
  )
	if err != nil {
		return nil, errorsmod.Wrapf(err, "failed to parse share denom")
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, shares)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrMintingShares, "error: %s", err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, senderAddr, shares)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrTransferingShares, "error: %s", err)
	}

	return &types.MsgCreatePoolResponse{}, nil
}
