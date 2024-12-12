package keeper

import (
	"context"

	"cosmuni/x/dex/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LiquidityPoolAll(ctx context.Context, req *types.QueryAllLiquidityPoolRequest) (*types.QueryAllLiquidityPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var liquidityPools []types.LiquidityPool

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	liquidityPoolStore := prefix.NewStore(store, types.KeyPrefix(types.LiquidityPoolKeyPrefix))

	pageRes, err := query.Paginate(liquidityPoolStore, req.Pagination, func(key []byte, value []byte) error {
		var liquidityPool types.LiquidityPool
		if err := k.cdc.Unmarshal(value, &liquidityPool); err != nil {
			return err
		}

		liquidityPools = append(liquidityPools, liquidityPool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLiquidityPoolResponse{LiquidityPool: liquidityPools, Pagination: pageRes}, nil
}

func (k Keeper) LiquidityPool(ctx context.Context, req *types.QueryGetLiquidityPoolRequest) (*types.QueryGetLiquidityPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetLiquidityPool(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetLiquidityPoolResponse{LiquidityPool: val}, nil
}
