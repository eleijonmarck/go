package main

import (
	"fmt"
	"github.com/eleijonmarck/adventofcode5/stringchecker"
)

type testString struct {
	thestring string
	expected  bool
}

func main() {

	niceString := testString{
		thestring: "ugknbfddgicrmopn",
		expected:  true,
	}

	// anotherNiceString := testString{
	// 	thestring: "aaa",
	// 	expected:  true,
	// }

	badString := testString{
		thestring: "jchzalrnumimnmhpA",
		expected:  false,
	}

	sc := stringchecker.NewStringchecker()

	isitnice := sc.IsNice(niceString.thestring)
	fmt.Printf("The sc says that %s is %v and it should have been %v \n", niceString.thestring, isitnice, niceString.expected)

	isitnice2 := sc.IsNice(badString.thestring)
	fmt.Printf("The sc says that %s is %v and it should have been %v \n", badString.thestring, isitnice2, badString.expected)
}
