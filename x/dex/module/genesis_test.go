package dex_test

import (
	"testing"

	keepertest "cosmuni/testutil/keeper"
	"cosmuni/testutil/nullify"
	dex "cosmuni/x/dex/module"
	"cosmuni/x/dex/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LiquidityPoolList: []types.LiquidityPool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, k, genesisState)
	got := dex.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LiquidityPoolList, got.LiquidityPoolList)
	// this line is used by starport scaffolding # genesis/test/assert
}
