// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reverese rune-wise left to right.
func Reverse(s string) string {
	return reverseTwo(s)
}