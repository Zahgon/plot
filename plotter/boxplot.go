// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// fiveStatPlot contains the shared fields for quartile
// and box-whisker plots.
type fiveStatPlot struct {
	// Values is a copy of the values of the values used to
	// create this box plot.
	Values

	// Location is the location of the box along its axis.
	Location float64

	// Median is the median value of the data.
	Median float64

	// Quartile1 and Quartile3 are the first and
	// third quartiles of the data respectively.
	Quartile1, Quartile3 float64

	// AdjLow and AdjHigh are the `adjacent' values
	// on the low and high ends of the data.  The
	// adjacent values are the points to which the
	// whiskers are drawn.
	AdjLow, AdjHigh float64

	// Min and Max are the extreme values of the data.
	Min, Max float64

	// Outside are the indices of Vs for the outside points.
	Outside []int
}

// BoxPlot implements the Plotter interface, drawing
// a boxplot to represent the distribution of values.
type BoxPlot struct {
	fiveStatPlot

	// Offset is added to the x location of each box.
	// When the Offset is zero, the boxes are drawn
	// centered at their x location.
	Offset vg.Length

	// Width is the width used to draw the box.
	Width vg.Length

	// CapWidth is the width of the cap used to top
	// off a whisker.
	CapWidth vg.Length

	// GlyphStyle is the style of the outside point glyphs.
	GlyphStyle draw.GlyphStyle

	// FillColor is the color used to fill the box.
	// The default is no fill.
	FillColor color.Color

	// BoxStyle is the line style for the box.
	BoxStyle draw.LineStyle

	// MedianStyle is the line style for the median line.
	MedianStyle draw.LineStyle

	// WhiskerStyle is the line style used to draw the
	// whiskers.
	WhiskerStyle draw.LineStyle

	// Horizontal dictates whether the BoxPlot should be in the vertical
	// (default) or horizontal direction.
	Horizontal bool
}

// NewBoxPlot returns a new BoxPlot that represents
// the distribution of the given values.  The style of
// the box plot is that used for Tukey's schematic
// plots in “Exploratory Data Analysis.”
//
// An error is returned if the boxplot is created with
// no values.
//
// The fence values are 1.5x the interquartile before
// the first quartile and after the third quartile.  Any
// value that is outside of the fences are drawn as
// Outside points.  The adjacent values (to which the
// whiskers stretch) are the minimum and maximum
// values that are not outside the fences.
func NewBoxPlot(w vg.Length, loc float64, values Valuer) (*BoxPlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFiveStat(w vg.Length, loc float64, values Valuer) (fiveStatPlot, error) {
	_ = "STUB: not implemented"
	return *new(fiveStatPlot), nil
}

// median returns the median value from a
// sorted Values.
func median(vs Values) float64 { _ = "STUB: not implemented"; return 0 }

// Plot draws the BoxPlot on Canvas c and Plot plt.
func (b *BoxPlot) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum x
// and y values, implementing the plot.DataRanger
// interface.
func (b *BoxPlot) DataRange() (float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes returns a slice of GlyphBoxes for the
// points and for the median line of the boxplot,
// implementing the plot.GlyphBoxer interface
func (b *BoxPlot) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }

// OutsideLabels returns a *Labels that will plot
// a label for each of the outside points.  The
// labels are assumed to correspond to the
// points used to create the box plot.
func (b *BoxPlot) OutsideLabels(labels Labeller) (*Labels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type boxPlotOutsideLabels struct {
	box    *BoxPlot
	labels []string
}

func (o boxPlotOutsideLabels) Len() int { _ = "STUB: not implemented"; return 0 }

func (o boxPlotOutsideLabels) XY(i int) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

func (o boxPlotOutsideLabels) Label(i int) string {
	_ = "STUB: not implemented"

	// horizBoxPlot is like a regular BoxPlot, however,
	// it draws horizontally instead of Vertically.
	// TODO: Merge code for horizontal and vertical box plots as has been done for
	// bar charts.
	return ""
}

type horizBoxPlot struct{ *BoxPlot }

func (b horizBoxPlot) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum x
// and y values, implementing the plot.DataRanger
// interface.
func (b horizBoxPlot) DataRange() (float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes returns a slice of GlyphBoxes for the
// points and for the median line of the boxplot,
// implementing the plot.GlyphBoxer interface
func (b horizBoxPlot) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}

// OutsideLabels returns a *Labels that will plot
// a label for each of the outside points.  The
// labels are assumed to correspond to the
// points used to create the box plot.
func (b *horizBoxPlot) OutsideLabels(labels Labeller) (*Labels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type horizBoxPlotOutsideLabels struct {
	boxPlotOutsideLabels
}

func (o horizBoxPlotOutsideLabels) XY(i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

// ValueLabels implements both the Valuer
// and Labeller interfaces.
type ValueLabels []struct {
	Value float64
	Label string
}

// Len returns the number of items.
func (vs ValueLabels) Len() int {
	_ = "STUB: not implemented"

	// Value returns the value of item i.
	return 0
}

func (vs ValueLabels) Value(i int) float64 {
	_ = "STUB: not implemented"

	// Label returns the label of item i.
	return 0
}

func (vs ValueLabels) Label(i int) string { _ = "STUB: not implemented"; return "" }
