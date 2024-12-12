package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetLiquidityPool set a specific liquidityPool in the store from its index
func (k Keeper) SetLiquidityPool(ctx context.Context, liquidityPool types.LiquidityPool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidityPoolKeyPrefix))
	b := k.cdc.MustMarshal(&liquidityPool)
	store.Set(types.LiquidityPoolKey(
		liquidityPool.Index,
	), b)
}

// GetLiquidityPool returns a liquidityPool from its index
func (k Keeper) GetLiquidityPool(
	ctx context.Context,
	index string,

) (val types.LiquidityPool, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidityPoolKeyPrefix))

	b := store.Get(types.LiquidityPoolKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLiquidityPool removes a liquidityPool from the store
func (k Keeper) RemoveLiquidityPool(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidityPoolKeyPrefix))
	store.Delete(types.LiquidityPoolKey(
		index,
	))
}

// GetAllLiquidityPool returns all liquidityPool
func (k Keeper) GetAllLiquidityPool(ctx context.Context) (list []types.LiquidityPool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LiquidityPoolKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.LiquidityPool
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
