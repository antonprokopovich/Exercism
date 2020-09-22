package acronym

import "strings"

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	var res string

	for _, w := range strings.Fields(s) {
		res = res + strings.ToUpper(w[:1])
	}
	return res
}
