package stringchecker

import (
	"strings"
)

// Stringchecker checks whether a string is nice or bad
type Stringchecker struct {
	niceRules
	badRules
}

type niceRules struct {
}

type badRules struct {
	BadStrings []string
}

func (sc *Stringchecker) containsAtLeastThreeVows(s string) bool {
	var count int
	if strings.ContainsAny(s, "a") {
		count++
	}
	if strings.ContainsAny(s, "e") {
		count++
	}
	if strings.ContainsAny(s, "i") {
		count++
	}
	if strings.ContainsAny(s, "u") {
		count++
	}
	if strings.ContainsAny(s, "o") {
		count++
	}

	if count >= 3 {
		return true
	}
	return false
}

func (r *niceRules) containsCertainLettersThatAppearTwice(s string) bool {
	if strings.Contains(s, "aa") || strings.Contains(s, "bb") || strings.Contains(s, "cc") || strings.Contains(s, "dd") {
		return true
	}
	return false
}

func (r *badRules) containsBadStrings(s ...string) bool {
	for _, sub := range s {
		for _, bad := range r.BadStrings {
			if strings.Contains(sub, bad) {
				return true
			}
		}
	}
	return false
}

// IsNice return true if the string is nice, or false if it is bad
func (sc *Stringchecker) IsNice(s string) bool {

	if sc.containsBadStrings(s) {
		return false
	}
	if sc.containsAtLeastThreeVows(s) && sc.containsCertainLettersThatAppearTwice(s) {
		return true
	}
	return false
}

// NewStringchecker returns a new Stringchecker
func NewStringchecker() *Stringchecker {

	// here we can change what bad substrings we want
	var badStrings = []string{
		"ab",
		"cd",
		"pq",
		"xy",
		// can be added to see that if we choose string; it fails running main
		//		"ug",
	}

	stringchecker := Stringchecker{
		niceRules{},
		badRules{
			BadStrings: badStrings,
		},
	}
	return &stringchecker
}
