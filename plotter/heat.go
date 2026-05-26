// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/vg/draw"
)

// GridXYZ describes three dimensional data where the X and Y
// coordinates are arranged on a rectangular grid.
type GridXYZ interface {
	// Dims returns the dimensions of the grid.
	Dims() (c, r int)

	// Z returns the value of a grid value at (c, r).
	// It will panic if c or r are out of bounds for the grid.
	Z(c, r int) float64

	// X returns the coordinate for the column at the index c.
	// It will panic if c is out of bounds for the grid.
	X(c int) float64

	// Y returns the coordinate for the row at the index r.
	// It will panic if r is out of bounds for the grid.
	Y(r int) float64
}

// HeatMap implements the Plotter interface, drawing
// a heat map of the values in the GridXYZ field.
type HeatMap struct {
	GridXYZ GridXYZ

	// Palette is the color palette used to render
	// the heat map. Palette must not be nil or
	// return a zero length []color.Color.
	Palette palette.Palette

	// Underflow and Overflow are colors used to fill
	// heat map elements outside the dynamic range
	// defined by Min and Max.
	Underflow color.Color
	Overflow  color.Color

	// NaN is the color used to fill heat map elements
	// that are NaN or do not map to a unique palette
	// color.
	NaN color.Color

	// Min and Max define the dynamic range of the
	// heat map.
	Min, Max float64

	// Rasterized indicates whether the heatmap
	// should be produced using raster-based drawing.
	Rasterized bool
}

// NewHeatMap creates as new heat map plotter for the given data,
// using the provided palette. If g has Min and Max methods that return
// a float, those returned values are used to set the respective HeatMap
// fields. If the returned HeatMap is used when Min is greater than Max,
// the Plot method will panic.
func NewHeatMap(g GridXYZ, p palette.Palette) *HeatMap { _ = "STUB: not implemented"; return nil }

// Plot implements the Plot method of the plot.Plotter interface.
func (h *HeatMap) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// plotRasterized plots the heatmap using raster-based drawing.
func (h *HeatMap) plotRasterized(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// Apply palette scaling.

// plotVectorized plots the heatmap using vector-based drawing.
func (h *HeatMap) plotVectorized(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// ps scales the palette uniformly across the data range.

// Apply palette scaling.

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (h *HeatMap) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// Make a unit length when there is no neighbour.

// Make a unit length when there is no neighbour.

// GlyphBoxes implements the GlyphBoxes method
// of the plot.GlyphBoxer interface.
func (h *HeatMap) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }
