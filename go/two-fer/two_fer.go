package twofer

import "fmt"

// ShareWith returns string "One for {name}, one for me" if name is present.
// Otherwise it returns "One for you, one for me".
func ShareWith(name string) string {
	var who string
	if name == "" {
		who = "you"
	} else {
		who = name
	}
	return fmt.Sprintf("One for %s, one for me.", who)
}
