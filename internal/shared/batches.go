package shared

func BatchSlice[T any](slice []T, batchSize int) [][]T {
	var batches [][]T

	for batchSize < len(slice) {
		slice, batches = slice[batchSize:], append(batches, slice[0:batchSize:batchSize])
	}
	batches = append(batches, slice)

	return batches
}
