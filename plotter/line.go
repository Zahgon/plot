// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

// StepKind specifies a form of a connection of two consecutive points.
type StepKind int

const (
	// NoStep connects two points by simple line
	NoStep StepKind = iota

	// PreStep connects two points by following lines: vertical, horizontal.
	PreStep

	// MidStep connects two points by following lines: horizontal, vertical, horizontal.
	// Vertical line is placed in the middle of the interval.
	MidStep

	// PostStep connects two points by following lines: horizontal, vertical.
	PostStep
)

// Line implements the Plotter interface, drawing a line.
type Line struct {
	// XYs is a copy of the points for this line.
	XYs

	// StepStyle is the kind of the step line.
	StepStyle StepKind

	// LineStyle is the style of the line connecting the points.
	// Use zero width to disable lines.
	draw.LineStyle

	// FillColor is the color to fill the area below the plot.
	// Use nil to disable the filling. This is the default.
	FillColor color.Color
}

// NewLine returns a Line that uses the default line style and
// does not draw glyphs.
func NewLine(xys XYer) (*Line, error) { _ = "STUB: not implemented"; return nil, nil }

// Plot draws the Line, implementing the plot.Plotter interface.
func (pts *Line) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum
// x and y values, implementing the plot.DataRanger interface.
func (pts *Line) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0,

		// GlyphBoxes implements the plot.GlyphBoxer interface.
		0, 0, 0
}

func (pts *Line) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }

// Thumbnail returns the thumbnail for the Line, implementing the plot.Thumbnailer interface.
func (pts *Line) Thumbnail(c *draw.Canvas) { _ = "STUB: not implemented"; return }

// NewLinePoints returns both a Line and a
// Points for the given point data.
func NewLinePoints(xys XYer) (*Line, *Scatter, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
