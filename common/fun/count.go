package fun

func Count[T any](ts []T, pred func(T) bool) int {

	count := 0
	for _, t := range ts {
		if pred(t) {
			count++
		}
	}
	return count
}
