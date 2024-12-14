package types

import (
  "fmt"
  "math"
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

func PoolDenom(poolId string) string {
  return fmt.Sprintf("%s-shares", poolId)
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

func FormatCoinsStr(token0, token1 string, amount0, amount1 uint64) string {
  return fmt.Sprintf("%d%s,%d%s", amount0, token0, amount1, token1)
}

func FormatShareCoinStr(poolId string, amount uint64) string {
  denom := PoolDenom(poolId)
  return fmt.Sprintf("%d%s", amount, denom)
}
