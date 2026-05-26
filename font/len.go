// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package font

// Length is a unit-independent representation of length.
// Internally, the length is stored in postscript points.
type Length float64

// Dots returns the length in dots for the given resolution.
func (l Length) Dots(dpi float64) float64 { _ = "STUB: not implemented"; return 0 }

// Points returns the length in postscript points.
func (l Length) Points() float64 {
	_ = "STUB: not implemented"

	// Common lengths.
	return 0
}

const (
	Inch       Length = 72
	Centimeter        = Inch / 2.54
	Millimeter        = Centimeter / 10
)

// Points returns a length for the given number of points.
func Points(pt float64) Length {
	_ = "STUB: not implemented"

	// ParseLength parses a Length string.
	// A Length string is a possible signed floating number with a unit.
	// e.g. "42cm" "2.4in" "66pt"
	// If no unit was given, ParseLength assumes it was (postscript) points.
	// Currently valid units are:
	//
	//   - mm (millimeter)
	//   - cm (centimeter)
	//   - in (inch)
	//   - pt (point)
	return *new(Length)
}

func ParseLength(value string) (Length, error) { _ = "STUB: not implemented"; return *new(Length), nil }
