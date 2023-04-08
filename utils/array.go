package utils

// forked this methods from core-js
func Some[T any](arr []T, check func(T) bool) bool {
	for _, v := range arr {
		if check(v) {
			return true
		}
	}
	return false
}

func Have[T any](arr []T, check func(T) bool) bool {
	for _, v := range arr {
		if !check(v) {
			return false
		}
	}
	return true
}
