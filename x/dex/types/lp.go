package types

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func OrderTokensAndAmounts(token0, token1 string, amount0, amount1 uint64) (string, string, uint64, uint64) {
	if token0 >= token1 {
		return token0, token1, amount0, amount1
	} else {
		return token1, token0, amount1, amount0
	}
}

func GeneratePoolId(token0, token1 string) string {
	if token0 >= token1 {
		return fmt.Sprintf("%s-%s", token0, token1)
	}

	return fmt.Sprintf("%s-%s", token1, token0)
}

func CalculateShares(amount0, amount1, totalShares uint64) uint64 {
	shares := math.Sqrt(float64(amount0 * amount1))

	if totalShares == 0 {
		return uint64(shares)
	}

	return uint64(shares / float64(totalShares) * float64(totalShares))
}

func CalculateK(amount0, amount1 uint64) uint64 {
	return amount0 * amount1
}

func CalculateSharesPercentage(shares, totalShares uint64) float64 {
	return float64(shares) / float64(totalShares)
}

func CalculateSwapAmount(k, amount0, amount1, amountIn0, amountIn1 uint64) (uint64, uint64) {
  if amountIn0 == 0 {
    out := amount0 - (k / (amount1 + amountIn1))
    return out, 0
  } else {
    out := amount1 - (k / (amount0 + amountIn0))
    return 0, out
  }
}

func PoolDenom(poolId string) string {
	return fmt.Sprintf("%s-shares", poolId)
}

func CreateLPCoins(token0, token1 string, amount0, amount1 uint64) (sdk.Coins, error) {
  coinsStr := fmt.Sprintf("%d%s,%d%s", amount0, token0, amount1, token1)
  return sdk.ParseCoinsNormalized(coinsStr)
}

// remove
func FormatCoinsStr(token0, token1 string, amount0, amount1 uint64) string {
	return fmt.Sprintf("%d%s,%d%s", amount0, token0, amount1, token1)
}

func CreateSharesCoins(poolId string, amount uint64) (sdk.Coins, error) {
	denom := PoolDenom(poolId)
  coinsStr := fmt.Sprintf("%d%s", amount, denom)
  return sdk.ParseCoinsNormalized(coinsStr)
}

// remove
func FormatShareCoinStr(poolId string, amount uint64) string {
	denom := PoolDenom(poolId)
	return fmt.Sprintf("%d%s", amount, denom)
}
