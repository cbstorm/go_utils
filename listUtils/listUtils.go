package listUtils

import "fmt"

func Mapping[K any, V any](list []K, f func(e K, idx int) V) []V {
	n := make([]V, len(list))
	for i, e := range list {
		n[i] = f(e, i)
	}
	return n
}

func ForEach[K any](list []K, f func(e K, idx int)) {
	for i, e := range list {
		f(e, i)
	}
}

func Filter[K any](list []K, f func(e K, idx int) bool) []K {
	n := make([]K, 0)
	for i, e := range list {
		if f(e, i) {
			n = append(n, e)
		}
	}
	return n
}

func Find[K any](list []K, f func(e K) bool) (K, bool) {
	var n K
	for _, e := range list {
		if f(e) {
			return e, true
		}
	}

	return n, false
}

func FindIndex[K any](list []K, f func(e K) bool) int {
	for i, e := range list {
		if f(e) {
			return i
		}
	}
	return -1
}

func Includes[K int | string | bool](list []K, compare K) bool {
	for _, e := range list {
		if e == compare {
			return true
		}
	}
	return false
}

func Join(list []string, separator string) string {
	var n string
	for i, e := range list {
		if i == 0 {
			n = e
		} else {
			n = fmt.Sprintf("%s%s%s", n, separator, e)
		}
	}
	return n
}

func Reduce[K any, V any](list []K, f func(a *V, b K), inititalValue V) V {

	for _, e := range list {
		f(&inititalValue, e)
	}
	return inititalValue
}
