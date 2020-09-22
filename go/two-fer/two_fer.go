package twofer

import "strings"

const pattern = "One for you, one for me."

// ShareWith returns string "One for {name}, one for me." if name is present.
// Otherwise returns "One for you, one for me."
func ShareWith(name string) string {
	result := pattern
	if name != "" {
		result = strings.Replace(result, "you", name, -1)
	}
	return result
}
