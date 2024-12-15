package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poolId := types.GeneratePoolId(msg.Token0, msg.Token1)
	t0, t1, a0, a1 := types.OrderTokensAndAmounts(
		msg.Token0,
		msg.Token1,
		msg.Amount0,
		msg.Amount1,
	)

	pool, found := k.Keeper.GetLiquidityPool(ctx, poolId)
	if !found {
		return nil, errorsmod.Wrapf(types.ErrPoolNonExistant, "pool %s does not exist", poolId)
	}

	senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	sharesAmount := types.CalculateShares(a0, a1, pool.TotalShares)

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

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, shares)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrMintingShares, "error: %s", err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		senderAddr,
		shares,
	)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrTransferingShares, "error: %s", err)
	}

	pool.TotalShares += sharesAmount
	pool.Amount0 += a0
	pool.Amount1 += a1
	k.SetLiquidityPool(ctx, pool)

	return &types.MsgDepositResponse{}, nil
}
