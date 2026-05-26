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

var (
	// DefaultQuartMedianStyle is a fat dot.
	DefaultQuartMedianStyle = draw.GlyphStyle{
		Color:  color.Black,
		Radius: vg.Points(1.5),
		Shape:  draw.CircleGlyph{},
	}

	// DefaultQuartWhiskerStyle is a hairline.
	DefaultQuartWhiskerStyle = draw.LineStyle{
		Color:    color.Black,
		Width:    vg.Points(0.5),
		Dashes:   []vg.Length{},
		DashOffs: 0,
	}
)

// QuartPlot implements the Plotter interface, drawing
// a plot to represent the distribution of values.
//
// This style of the plot appears in Tufte's "The Visual
// Display of Quantitative Information".
type QuartPlot struct {
	fiveStatPlot

	// Offset is added to the x location of each plot.
	// When the Offset is zero, the plot is drawn
	// centered at its x location.
	Offset vg.Length

	// MedianStyle is the line style for the median point.
	MedianStyle draw.GlyphStyle

	// WhiskerStyle is the line style used to draw the
	// whiskers.
	WhiskerStyle draw.LineStyle

	// Horizontal dictates whether the QuartPlot should be in the vertical
	// (default) or horizontal direction.
	Horizontal bool
}

// NewQuartPlot returns a new QuartPlot that represents
// the distribution of the given values.
//
// An error is returned if the plot is created with
// no values.
//
// The fence values are 1.5x the interquartile before
// the first quartile and after the third quartile.  Any
// value that is outside of the fences are drawn as
// Outside points.  The adjacent values (to which the
// whiskers stretch) are the minimum and maximum
// values that are not outside the fences.
func NewQuartPlot(loc float64, values Valuer) (*QuartPlot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Plot draws the QuartPlot on Canvas c and Plot plt.
func (b *QuartPlot) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum x
// and y values, implementing the plot.DataRanger
// interface.
func (b *QuartPlot) DataRange() (float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes returns a slice of GlyphBoxes for the plot,
// implementing the plot.GlyphBoxer interface.
func (b *QuartPlot) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}

// OutsideLabels returns a *Labels that will plot
// a label for each of the outside points.  The
// labels are assumed to correspond to the
// points used to create the plot.
func (b *QuartPlot) OutsideLabels(labels Labeller) (*Labels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type quartPlotOutsideLabels struct {
	qp     *QuartPlot
	labels []string
}

func (o quartPlotOutsideLabels) Len() int { _ = "STUB: not implemented"; return 0 }

func (o quartPlotOutsideLabels) XY(i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (o quartPlotOutsideLabels) Label(i int) string {
	_ = "STUB: not implemented"

	// horizQuartPlot is like a regular QuartPlot, however,
	// it draws horizontally instead of Vertically.
	return ""
}

type horizQuartPlot struct{ *QuartPlot }

func (b horizQuartPlot) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum x
// and y values, implementing the plot.DataRanger
// interface.
func (b horizQuartPlot) DataRange() (float64, float64, float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes returns a slice of GlyphBoxes for the plot,
// implementing the plot.GlyphBoxer interface.
func (b horizQuartPlot) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}

// OutsideLabels returns a *Labels that will plot
// a label for each of the outside points.  The
// labels are assumed to correspond to the
// points used to create the plot.
func (b *horizQuartPlot) OutsideLabels(labels Labeller) (*Labels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type horizQuartPlotOutsideLabels struct {
	quartPlotOutsideLabels
}

func (o horizQuartPlotOutsideLabels) XY(i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}
