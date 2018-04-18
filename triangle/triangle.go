// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "not a triangle" // not a triangle
	Equ = "equilateral"    // equilateral
	Iso = "isosceles"      // isosceles
	Sca = "scalene"        // scalene
)

// ShareWith should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	// sides are now sorted
	switch {
	case math.IsInf(c, 1): // largest side is +Inf, guards against (3, +Inf, +Inf)
		return NaT
	case a <= 0:
		return NaT
	case a+b < c: // triangle inequality
		return NaT
	case a == c:
		return Equ
	case a == b || b == c:
		return Iso
	}
	return Sca
}
