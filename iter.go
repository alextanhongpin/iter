package iter

import (
	"golang.org/x/exp/constraints"
)

func Map[T any, R any](ts []T, fn func(T, int) R) []R {
	res := make([]R, len(ts))
	for i, t := range ts {
		res[i] = fn(t, i)
	}

	return res
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	res := make([]T, 0, len(ts))

	for _, t := range ts {
		if !fn(t) {
			continue
		}

		res = append(res, t)
	}

	return res
}

func Reduce[T any, R constraints.Ordered](ts []T, r R, fn func(T, R, int) R) R {
	for i, t := range ts {
		r = fn(t, r, i)
	}

	return r
}
