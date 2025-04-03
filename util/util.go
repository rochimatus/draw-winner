package util

func SliceFilter[T interface{}](filterFunc func(T) bool, collections []T) []T {
	data := []T{}
	for _, v := range collections {
		if filterFunc(v) {
			data = append(data, v)
		}
	}
	return data
}
