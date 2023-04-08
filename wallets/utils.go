package wallets

import "unsafe"

// b2s converts a byte slice to a string without memory allocation.
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
