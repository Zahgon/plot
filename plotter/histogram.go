// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

// Histogram implements the Plotter interface,
// drawing a histogram of the data.
type Histogram struct {
	// Bins is the set of bins for this histogram.
	Bins []HistogramBin

	// Width is the width of each bin.
	Width float64

	// FillColor is the color used to fill each
	// bar of the histogram.  If the color is nil
	// then the bars are not filled.
	FillColor color.Color

	// LineStyle is the style of the outline of each
	// bar of the histogram.
	draw.LineStyle

	// LogY allows rendering with a log-scaled Y axis.
	// When enabled, histogram bins with no entries will be discarded from
	// the histogram's DataRange.
	// The lowest Y value for the DataRange will be corrected to leave an
	// arbitrary amount of height for the smallest bin entry so it is visible
	// on the final plot.
	LogY bool
}

// NewHistogram returns a new histogram
// that represents the distribution of values
// using the given number of bins.
//
// Each y value is assumed to be the frequency
// count for the corresponding x.
//
// If the number of bins is non-positive than
// a reasonable default is used.
func NewHistogram(xy XYer, n int) (*Histogram, error) { _ = "STUB: not implemented"; return nil, nil }

// NewHist returns a new histogram, as in
// NewHistogram, except that it accepts a Valuer
// instead of an XYer.
func NewHist(vs Valuer, n int) (*Histogram, error) { _ = "STUB: not implemented"; return nil, nil }

type unitYs struct {
	Valuer
}

func (u unitYs) XY(i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0,

		// Plot implements the Plotter interface, drawing a line
		// that connects each point in the Line.
		0
}

func (h *Histogram) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum X and Y values
func (h *Histogram) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// ylow will hold the smallest non-zero y value.

// Reserve a bit of space for the smallest bin to be displayed still.

// Normalize normalizes the histogram so that the
// total area beneath it sums to a given value.
func (h *Histogram) Normalize(sum float64) { _ = "STUB: not implemented"; return }

// Thumbnail draws a rectangle in the given style of the histogram.
func (h *Histogram) Thumbnail(c *draw.Canvas) { _ = "STUB: not implemented"; return }

// binPoints returns a slice containing the
// given number of bins, and the width of
// each bin.
//
// If the given number of bins is not positive
// then a reasonable default is used.  The
// default is the square root of the sum of
// the y values.
func binPoints(xys XYer, n int) (bins []HistogramBin, width float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// A HistogramBin approximates the number of values
// within a range by a single number (the weight).
type HistogramBin struct {
	Min, Max float64
	Weight   float64
}
