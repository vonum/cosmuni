package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "cosmuni/testutil/keeper"
	"cosmuni/testutil/nullify"
	"cosmuni/x/dex/keeper"
	"cosmuni/x/dex/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNLiquidityPool(keeper keeper.Keeper, ctx context.Context, n int) []types.LiquidityPool {
	items := make([]types.LiquidityPool, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetLiquidityPool(ctx, items[i])
	}
	return items
}

func TestLiquidityPoolGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLiquidityPool(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetLiquidityPool(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestLiquidityPoolRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLiquidityPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLiquidityPool(ctx,
			item.Index,
		)
		_, found := keeper.GetLiquidityPool(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestLiquidityPoolGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNLiquidityPool(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLiquidityPool(ctx)),
	)
}
