package iter

import (
	"sync"
)

type Result[T any] struct {
	Data  T
	Error error
}

func MapIndex[T any, V any](ts []T, fn func(int) V) []V {
	res := make([]V, len(ts))
	for i := 0; i < len(ts); i++ {
		res[i] = fn(i)
	}

	return res
}

func Map[T any, V any](ts []T, fn func(T) V) []V {
	return MapIndex(ts, func(i int) V {
		return fn(ts[i])
	})
}

// MapError maps T to V, and returns on first error.
func MapError[T any, V any](ts []T, fn func(T) (V, error)) ([]V, error) {
	res := make([]V, len(ts))

	for i := range ts {
		v, err := fn(ts[i])
		if err != nil {
			return nil, err
		}
		res[i] = v
	}

	return res, nil
}

// MapResult maps each item into Result type.
func MapResult[T any, V any](ts []T, fn func(T) (V, error)) []Result[V] {
	res := make([]Result[V], len(ts))

	for i := range ts {
		v, err := fn(ts[i])
		if err != nil {
			res[i] = Result[V]{Error: err}
		} else {
			res[i] = Result[V]{Data: v}
		}
	}

	return res
}

func EachIndex[T any](ts []T, fn func(int)) {
	for i := 0; i < len(ts); i++ {
		fn(i)
	}
}

func Each[T any](ts []T, fn func(T)) {
	EachIndex(ts, func(i int) {
		fn(ts[i])
	})
}

func GoEachIndex[T any](ts []T, fn func(int)) {
	var wg sync.WaitGroup
	wg.Add(len(ts))

	for i := 0; i < len(ts); i++ {
		go func(i int) {
			defer wg.Done()

			fn(i)
		}(i)
	}

	wg.Wait()
}

func GoEach[T any](ts []T, fn func(T)) {
	GoEachIndex(ts, func(i int) {
		fn(ts[i])
	})
}

func GoMapIndex[T any, V any](ts []T, fn func(int) V) []V {
	res := make([]V, len(ts))

	var wg sync.WaitGroup
	wg.Add(len(ts))

	for i := 0; i < len(ts); i++ {
		go func(i int) {
			defer wg.Done()

			res[i] = fn(i)
		}(i)
	}
	wg.Wait()

	return res
}

func GoMap[T any, V any](ts []T, fn func(T) V) []V {
	return GoMapIndex(ts, func(i int) V {
		return fn(ts[i])
	})
}

func GoMapIndexFlat[T any, V any](ts []T, fn func(int) []V) []V {
	res := make([][]V, len(ts))

	var wg sync.WaitGroup
	wg.Add(len(ts))

	for i := range ts {
		go func(i int) {
			defer wg.Done()

			res[i] = fn(i)
		}(i)
	}
	wg.Wait()

	return Flat(res)
}

func GoMapFlat[T any, V any](ts []T, fn func(T) []V) []V {
	return GoMapIndexFlat(ts, func(i int) []V {
		return fn(ts[i])
	})
}

func GoMapResult[T any, V any](ts []T, fn func(T) (V, error)) []Result[V] {
	res := make([]Result[V], len(ts))

	var wg sync.WaitGroup
	wg.Add(len(ts))

	for i := range ts {
		go func(i int) {
			defer wg.Done()

			v, err := fn(ts[i])
			if err != nil {
				res[i] = Result[V]{Error: err}
			} else {
				res[i] = Result[V]{Data: v}
			}
		}(i)
	}
	wg.Wait()

	return res
}

func FilterIndex[T any](ts []T, fn func(int) bool) []T {
	res := make([]T, 0, len(ts))

	for i := 0; i < len(ts); i++ {
		if !fn(i) {
			continue
		}

		res = append(res, ts[i])
	}

	return res
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	return FilterIndex(ts, func(i int) bool {
		return fn(ts[i])
	})
}

func ReduceIndex[T, V any](ts []T, r V, fn func(V, int) V) V {
	for i := 0; i < len(ts); i++ {
		r = fn(r, i)
	}

	return r
}

func Reduce[T, V any](ts []T, r V, fn func(V, T) V) V {
	return ReduceIndex(ts, r, func(v V, i int) V {
		return fn(v, ts[i])
	})
}

func Flat[T any](ts [][]T) []T {
	var res []T
	for _, t := range ts {
		res = append(res, t...)
	}
	return res
}
