package types

import "fmt"


func GeneratePoolId(token0, token1 string, amount0, amount1 uint64) (string, string, string, uint64, uint64) {
  if token0 >= token1 {
    return fmt.Sprintf("%s-%s", token0, token1), token0, token1, amount0, amount1
  } else {
    return fmt.Sprintf("%s-%s", token1, token0), token1, token0, amount1, amount0
  }
}
