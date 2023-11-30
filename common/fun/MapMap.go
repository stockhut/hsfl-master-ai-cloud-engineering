package fun

import "golang.org/x/exp/maps"

func MapToSlice[K comparable, V any, T any](m map[K]V, f func(k K, v V) T) []T {

	ts := make([]T, len(maps.Keys(m)))
	i := 0
	for k, v := range m {
		ts[i] = f(k, v)
		i++
	}

	return ts
}
