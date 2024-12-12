package types

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"
)

var (
	ParamsKey = []byte("p_dex")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
