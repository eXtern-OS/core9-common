package utils

func Ternary[K any](condition bool, v1, v2 K) K {
	if condition {
		return v1
	} else {
		return v2
	}
}
