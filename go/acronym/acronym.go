package acronym

import "strings"

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) (res string) {

	s = strings.Replace(s, "-", " ", -1)
	s = strings.Replace(s, "_", " ", -1)

	for _, w := range strings.Fields(s) {
		res = res + strings.ToUpper(w[:1])
	}
	return res
}
