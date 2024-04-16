package utils

type Hashable interface {
	int | string
}

func ArrayToMap[K Hashable, V any](arr []K, defVal V) map[K]V {
	res := make(map[K]V)
	for _, x := range arr {
		res[x] = defVal
	}
	return res
}
