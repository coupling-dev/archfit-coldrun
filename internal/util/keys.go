// Package util holds small helpers shared by the store.
package util

// ValidKey reports whether key may be stored.
func ValidKey(key string) bool {
	return key != ""
}
