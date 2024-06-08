package utils

// ArrayToInterface is useful for MongoDB
func ArrayToInterface[K any](arr []K) []interface{} {
	var res []interface{}
	for _, x := range arr {
		res = append(res, x)
	}
	return res
}
