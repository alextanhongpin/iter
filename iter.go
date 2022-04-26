package iter

import (
	"golang.org/x/exp/constraints"
)

func Map[T any, R any](ts []T, fn func(T) R) []R {
	res := make([]R, len(ts))
	for i, t := range ts {
		res[i] = fn(t)
	}
	return res
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	var res []T
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
