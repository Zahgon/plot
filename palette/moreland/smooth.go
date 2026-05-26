// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moreland

import (
	"image/color"

	"gonum.org/v1/plot/palette"
)

// smoothDiverging is a smooth diverging color palette as described in
// "Diverging Color Maps for Scientific Visualization." by Kenneth Moreland,
// in Proceedings of the 5th International Symposium on Visual Computing,
// December 2009. DOI 10.1007/978-3-642-10520-3_9.
type smoothDiverging struct {
	// start and end are the beginning and ending colors
	start, end msh

	// convergeM is the MSH magnitude of the convergence point.
	// It is 88 by default.
	convergeM float64

	// alpha represents the opacity of the returned
	// colors in the range (0,1). It is 1 by default.
	alpha float64

	// min and max are the minimum and maximum values of the range of
	// scalars that can be mapped to colors using this palette.
	min, max float64

	// convergePoint is a number between min and max where the colors
	// should converge.
	convergePoint float64
}

// NewSmoothDiverging creates a new smooth diverging ColorMap as described in
// "Diverging Color Maps for Scientific Visualization." by Kenneth Moreland,
// in Proceedings of the 5th International Symposium on Visual Computing,
// December 2009. DOI 10.1007/978-3-642-10520-3_9.
//
// start and end are the start- and end-point colors and
// convergeM is the magnitude of the convergence point in
// magnitude-saturation-hue (MSH) color space. Note that
// convergeM specifies the color of the convergence point; it does not
// specify the location of the convergence point.
func NewSmoothDiverging(start, end color.Color, convergeM float64) palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// newSmoothDiverging creates a new smooth diverging ColorMap
// where start and end are the start and end point colors in MSH space and
// convergeM is the MSH magnitude of the convergence point. Note that
// convergeM specifies the color of the convergence point; it does not
// specify the location of the convergence point.
func newSmoothDiverging(start, end msh, convergeM float64) palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// At implements the palette.ColorMap interface.
func (p *smoothDiverging) At(v float64) (color.Color, error) {
	_ = "STUB: not implemented"
	return *new(color.Color), nil
}

func inUnitRange(v float64) bool { _ = "STUB: not implemented"; return false }

// SetMax implements the palette.ColorMap interface.
func (p *smoothDiverging) SetMax(v float64) { _ = "STUB: not implemented"; return }

// SetMin implements the palette.ColorMap interface.
func (p *smoothDiverging) SetMin(v float64) { _ = "STUB: not implemented"; return }

// Max implements the palette.ColorMap interface.
func (p *smoothDiverging) Max() float64 {
	_ = "STUB: not implemented"

	// Min implements the palette.ColorMap interface.
	return 0
}

func (p *smoothDiverging) Min() float64 {
	_ = "STUB: not implemented"

	// SetAlpha sets the opacity value of this color map. Zero is transparent
	// and one is completely opaque.
	// The function will panic is alpha is not between zero and one.
	return 0
}

func (p *smoothDiverging) SetAlpha(alpha float64) { _ = "STUB: not implemented"; return }

// Alpha returns the opacity value of this color map.
func (p *smoothDiverging) Alpha() float64 {
	_ = "STUB: not implemented"

	// SetConvergePoint sets the value where the diverging colors
	// should meet.
	return 0
}

func (p *smoothDiverging) SetConvergePoint(val float64) { _ = "STUB: not implemented"; return }

// ConvergePoint returns the value where the diverging colors meet.
func (p *smoothDiverging) ConvergePoint() float64 { _ = "STUB: not implemented"; return 0 }

// interpolateMSHDiverging performs a color interpolation through MSH space,
// where scalar is a number between 0 and 1 that the
// color should be evaluated at, and convergePoint is a number between 0 and
// 1 where the colors should converge.
func (p *smoothDiverging) interpolateMSHDiverging(scalar, convergePoint float64) msh {
	_ = "STUB: not implemented"
	return *new(msh)
}

// interpolation factor

// interpolation factors

// Palette returns a palette.Palette with the specified number of colors.
func (p smoothDiverging) Palette(n int) palette.Palette {
	_ = "STUB: not implemented"
	return *new(palette.Palette)
}

// Avoid potential overflow on last element
// due to floating point error.

// SmoothBlueRed is a SmoothDiverging-class ColorMap ranging from blue to red.
func SmoothBlueRed() palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// SmoothPurpleOrange is a SmoothDiverging-class ColorMap ranging from purple to orange.
func SmoothPurpleOrange() palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// SmoothGreenPurple is a SmoothDiverging-class ColorMap ranging from green to purple.
func SmoothGreenPurple() palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// SmoothBlueTan is a SmoothDiverging-class ColorMap ranging from blue to tan.
func SmoothBlueTan() palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}

// SmoothGreenRed is a SmoothDiverging-class ColorMap ranging from green to red.
func SmoothGreenRed() palette.DivergingColorMap {
	_ = "STUB: not implemented"
	return *new(palette.DivergingColorMap)
}
