package shared

import (
	"fmt"
	"hash"
	"hash/fnv"
	"moonlogs/internal/lib/serialize"
	"strconv"
	"sync"
)

var fnvHasherPool = sync.Pool{
	New: func() interface{} {
		return fnv.New64a()
	},
}

func HashQuery[T ~map[string]interface{}](query T) (string, error) {
	formattedQuery := make(T)

	for k, value := range query {
		switch v := value.(type) {
		case string:
			vInt, err := strconv.Atoi(v)
			if err != nil {
				formattedQuery[k] = v
				continue
			}
			formattedQuery[k] = vInt
		default:
			formattedQuery[k] = v
		}
	}

	bytes, err := serialize.JSONMarshal(formattedQuery)
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
