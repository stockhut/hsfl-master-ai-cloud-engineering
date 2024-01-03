package fun

import (
	"fmt"
)

func Map[T any, U any](ts []T, f func(T) U) []U {

	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}

	return us
}

func MapOrErr[T any, U any](ts []T, f func(T) (U, error)) ([]U, error) {

	us := make([]U, len(ts))
	for i, t := range ts {
		u, err := f(t)
		if err != nil {
			return nil, fmt.Errorf("failed to map item %v: %w", t, err)
		}
		us[i] = u
	}

	return us, nil
}
