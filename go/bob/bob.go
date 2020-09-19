// Solution for "Bob" problem from https://exercism.io/tracks/go/exercises/bob

package bob

import (
	"strings"
	"unicode"
)

func IsNumber(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) && !unicode.IsSpace(c) && !unicode.IsPunct(c) {
			return false
		}
	}
	return true
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsYelling(remark string) bool {
	return IsUpper(strings.TrimSpace(remark)) && !IsNumber(remark)
}

func IsQuestion(remark string) bool {
	trimmed := strings.TrimSpace(remark)
	return len(trimmed) > 0 && trimmed[len(trimmed)-1:] == "?"
}

func IsJustAddressing(remark string) bool {
	return strings.TrimSpace(remark) == ""
}

const (
	ToAddressing      string = "Fine. Be that way!"
	ToYellingQuestion string = "Calm down, I know what I'm doing!"
	ToYelling         string = "Whoa, chill out!"
	ToQuestion        string = "Sure."
	ToOther           string = "Whatever."
)

// Returns standard response string for each remark type.
func Hey(remark string) string {
	switch {
	case IsJustAddressing(remark):
		return ToAddressing
	case IsYelling(remark) && IsQuestion(remark):
		return ToYellingQuestion
	case IsYelling(remark):
		return ToYelling
	case IsQuestion(remark):
		return ToQuestion
	default:
		return ToOther
	}
}
