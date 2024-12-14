package types

import "math"

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
