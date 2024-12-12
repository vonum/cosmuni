package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LiquidityPoolKeyPrefix is the prefix to retrieve all LiquidityPool
	LiquidityPoolKeyPrefix = "LiquidityPool/value/"
)

// LiquidityPoolKey returns the store key to retrieve a LiquidityPool from the index fields
func LiquidityPoolKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
