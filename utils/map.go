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

func ArrayToMapUniversal[K, R any, V Hashable](arr []K, defVal R, key func(K) V) map[V]R {
	res := make(map[V]R)
	for _, x := range arr {
		res[key(x)] = defVal
	}
	return res
}
