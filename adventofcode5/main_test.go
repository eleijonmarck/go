package main

import (
	"github.com/eleijonmarck/adventofcode5/stringchecker"
	"testing"
)

func TestStringCheckerForNiceString(t *testing.T) {

	niceString := testString{
		thestring: "ugknbfddgicrmopn",
		expected:  true,
	}

	sc := stringchecker.NewStringchecker()

	if actual, expected := sc.IsNice(niceString.thestring), niceString.expected; actual != expected {
		t.Errorf("Expected %s but found %s", expected, actual)
	} else {
		t.Log("Success")
	}
}

func TestStringCheckerForBadString(t *testing.T) {

	var badStrings = make([]testString, 3)

	badStrings = append(badStrings, testString{thestring: "jchzalrnumimnmhp", expected: false})
	badStrings = append(badStrings, testString{thestring: "haegwjzuvuyypxyu", expected: false})
	badStrings = append(badStrings, testString{thestring: "dvszwmarrgswjxmb", expected: false})

	sc := stringchecker.NewStringchecker()

	for _, bad := range badStrings {
		if actual, expected := sc.IsNice(bad.thestring), bad.expected; actual != expected {
			t.Errorf("Expected %s but found %s", expected, actual)
		} else {
			t.Log("Success")
		}
	}
}
