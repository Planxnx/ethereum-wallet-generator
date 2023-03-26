package utils

import "fmt"

func MustError[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func Must[T any](val T, err any) T {
	switch e := err.(type) {
	case bool:
		if !e {
			panic("not ok")
		}
	case error:
		if e != nil {
			panic(e)
		}
	default:
		if err != nil {
			panic(fmt.Sprintf("invalid error type, must be bool or error, got %T(%v)", err, err))
		}
	}
	return val
}
