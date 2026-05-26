// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright ©2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package palette provides basic color palette handling.
package palette // import "gonum.org/v1/plot/palette"

import (
	"errors"
	"image/color"
)

// Palette is a collection of colors ordered into a palette.
type Palette interface {
	Colors() []color.Color
}

// DivergingPalette is a collection of colors ordered into a palette with
// a critical class or break in the middle of the color range.
type DivergingPalette interface {
	Palette

	// CriticalIndex returns the indices of the lightest
	// (median) color or colors in the DivergingPalette.
	// The low and high index values will be equal when
	// there is a single median color.
	CriticalIndex() (low, high int)
}

// A ColorMap maps scalar values to colors.
type ColorMap interface {
	// At returns the color associated with the given value.
	// If the value is not between Max() and Min(), an error is returned.
	At(float64) (color.Color, error)

	// Max returns the current maximum value of the ColorMap.
	Max() float64

	// SetMax sets the maximum value of the ColorMap.
	SetMax(float64)

	// Min returns the current minimum value of the ColorMap.
	Min() float64

	// SetMin sets the minimum value of the ColorMap.
	SetMin(float64)

	// Alpha returns the opacity value of the ColorMap.
	Alpha() float64

	// SetAlpha sets the opacity value of the ColorMap. Zero is transparent
	// and one is completely opaque. The default value of alpha should be
	// expected to be one. The function should be expected to panic
	// if alpha is not between zero and one.
	SetAlpha(float64)

	// Palette creates a Palette with the specified number of colors
	// from the ColorMap.
	Palette(colors int) Palette
}

// DivergingColorMap maps scalar values to colors that diverge
// from a central value.
type DivergingColorMap interface {
	ColorMap

	// SetConvergePoint sets the value where the diverging colors
	// should meet. The default value should be expected to be
	// (Min() + Max()) / 2. It should be expected that calling either
	// SetMax() or SetMin() will set a new default value, so for a
	// custom convergence point this function should be called after
	// SetMax() and SetMin(). The function should be expected to panic
	// if the value is not between Min() and Max().
	SetConvergePoint(float64)

	// ConvergePoint returns the value where the diverging colors meet.
	ConvergePoint() float64
}

// Hue represents a hue in HSV color space. Valid Hues are within [0, 1].
type Hue float64

const (
	Red Hue = Hue(iota) / 6
	Yellow
	Green
	Cyan
	Blue
	Magenta
)

var (
	// ErrOverflow is the error returned by ColorMaps when the specified
	// value is greater than the maximum value.
	ErrOverflow = errors.New("palette: specified value > maximum")

	// ErrUnderflow is the error returned by ColorMaps when the specified
	// value is less than the minimum value.
	ErrUnderflow = errors.New("palette: specified value < minimum")

	// ErrNaN is the error returned by ColorMaps when the specified
	// value is NaN.
	ErrNaN = errors.New("palette: specified value == NaN")
)

// Complement returns the complementary hue of a Hue.
func (h Hue) Complement() Hue { _ = "STUB: not implemented"; return *new(Hue) }

type palette []color.Color

func (p palette) Colors() []color.Color { _ = "STUB: not implemented"; return nil }

type divergingPalette []color.Color

func (p divergingPalette) Colors() []color.Color { _ = "STUB: not implemented"; return nil }

func (d divergingPalette) CriticalIndex() (low, high int) { _ = "STUB: not implemented"; return 0, 0 }

// Rainbow returns a rainbow palette with the specified number of colors, saturation
// value and alpha, and hues in the specified range.
func Rainbow(colors int, start, end Hue, sat, val, alpha float64) Palette {
	_ = "STUB: not implemented"
	return *new(Palette)
}

// Heat returns a red to yellow palette with the specified number of colors and alpha.
func Heat(colors int, alpha float64) Palette { _ = "STUB: not implemented"; return *new(Palette) }

// Radial return a diverging palette across the specified range, through white and with
// the specified alpha.
func Radial(colors int, start, end Hue, alpha float64) DivergingPalette {
	_ = "STUB: not implemented"
	return *new(DivergingPalette)
}
