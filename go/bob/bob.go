// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	remark = strings.TrimSpace(remark)
	isQuestion, _ := regexp.MatchString("\\A.*\\?\\z", remark)
	isYelling, _ := regexp.MatchString("\\A[A-Z1-9\\s[:punct:]]+[A-Z][A-Z1-9\\s[:punct:]]+\\z", remark)
	switch {
	case isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isYelling:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
