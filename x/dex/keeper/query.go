package keeper

import (
	"cosmuni/x/dex/types"
)

var _ types.QueryServer = Keeper{}
