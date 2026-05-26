// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// FieldXY describes a two dimensional vector field where the
// X and Y coordinates are arranged on a rectangular grid.
type FieldXY interface {
	// Dims returns the dimensions of the grid.
	Dims() (c, r int)

	// Vector returns the value of a vector field at (c, r).
	// It will panic if c or r are out of bounds for the field.
	Vector(c, r int) XY

	// X returns the coordinate for the column at the index c.
	// It will panic if c is out of bounds for the grid.
	X(c int) float64

	// Y returns the coordinate for the row at the index r.
	// It will panic if r is out of bounds for the grid.
	Y(r int) float64
}

// Field implements the Plotter interface, drawing
// a vector field of the values in the FieldXY field.
type Field struct {
	FieldXY FieldXY

	// DrawGlyph is the user hook to draw a field
	// vector glyph. The function should draw a unit
	// vector to (1, 0) on the vg.Canvas, c with the
	// sty LineStyle. The Field plotter will rotate
	// and scale the unit vector appropriately.
	// If the magnitude of v is zero, no scaling or
	// rotation is performed.
	//
	// The direction and magnitude of v can be used
	// to determine properties of the glyph drawing
	// but should not be used to determine size or
	// directions of the glyph.
	//
	// If DrawGlyph is nil, a simple arrow will be
	// drawn.
	DrawGlyph func(c vg.Canvas, sty draw.LineStyle, v XY)

	// LineStyle is the style of the line used to
	// render vectors when DrawGlyph is nil.
	// Otherwise it is passed to DrawGlyph.
	LineStyle draw.LineStyle

	// max define the dynamic range of the field.
	max float64
}

// NewField creates a new vector field plotter.
func NewField(f FieldXY) *Field { _ = "STUB: not implemented"; return nil }

// Plot implements the Plot method of the plot.Plotter interface.
func (f *Field) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// Do not scale when the vector is zero, otherwise the
// user cannot render special-case glyphs for that case.

func drawVector(c vg.Canvas, v XY) { _ = "STUB: not implemented"; return }

// TODO(kortschak): Improve this arrow.

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (f *Field) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Make a unit length when there is no neighbour.

// Make a unit length when there is no neighbour.

// GlyphBoxes implements the GlyphBoxes method
// of the plot.GlyphBoxer interface.
func (f *Field) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }
