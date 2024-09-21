package shared

import (
	"fmt"
	"hash"
	"hash/fnv"
	"moonlogs/internal/lib/serialize"
	"sync"
)

var fnvHasherPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

func HashQuery[T ~map[string]interface{}](query T) (string, error) {
	bytes, err := serialize.JSONMarshal(query)
	if err != nil {
		return "", fmt.Errorf("failed creating record: %v", err)
	}

	FNV64Hasher := fnvHasherPool.Get().(hash.Hash64)
	defer fnvHasherPool.Put(FNV64Hasher)

	FNV64Hasher.Write(bytes)
	hashSum := FNV64Hasher.Sum64()
	FNV64Hasher.Reset()

	return fmt.Sprint(hashSum), nil
}
