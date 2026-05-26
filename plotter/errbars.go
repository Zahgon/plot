// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// DefaultCapWidth is the default width of error bar caps.
var DefaultCapWidth = vg.Points(5)

// YErrorBars implements the plot.Plotter, plot.DataRanger,
// and plot.GlyphBoxer interfaces, drawing vertical error
// bars, denoting error in Y values.
type YErrorBars struct {
	XYs

	// YErrors is a copy of the Y errors for each point.
	YErrors

	// LineStyle is the style used to draw the error bars.
	draw.LineStyle

	// CapWidth is the width of the caps drawn at the top
	// of each error bar.
	CapWidth vg.Length
}

// NewYErrorBars returns a new YErrorBars plotter, or an error on failure.
// The error values from the YErrorer interface are interpreted as relative
// to the corresponding Y value. The errors for a given Y value are computed
// by taking the absolute value of the error returned by the YErrorer
// and subtracting the first and adding the second to the Y value.
func NewYErrorBars(yerrs interface {
	XYer
	YErrorer
}) (*YErrorBars, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Plot implements the Plotter interface, drawing labels.
func (e *YErrorBars) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// drawCap draws the cap if it is not clipped.
func (e *YErrorBars) drawCap(c *draw.Canvas, x, y vg.Length) { _ = "STUB: not implemented"; return }

// DataRange implements the plot.DataRanger interface.
func (e *YErrorBars) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes implements the plot.GlyphBoxer interface.
func (e *YErrorBars) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}

// XErrorBars implements the plot.Plotter, plot.DataRanger,
// and plot.GlyphBoxer interfaces, drawing horizontal error
// bars, denoting error in Y values.
type XErrorBars struct {
	XYs

	// XErrors is a copy of the X errors for each point.
	XErrors

	// LineStyle is the style used to draw the error bars.
	draw.LineStyle

	// CapWidth is the width of the caps drawn at the top
	// of each error bar.
	CapWidth vg.Length
}

// Returns a new XErrorBars plotter, or an error on failure. The error values
// from the XErrorer interface are interpreted as relative to the corresponding
// X value. The errors for a given X value are computed by taking the absolute
// value of the error returned by the XErrorer and subtracting the first and
// adding the second to the X value.
func NewXErrorBars(xerrs interface {
	XYer
	XErrorer
}) (*XErrorBars, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Plot implements the Plotter interface, drawing labels.
func (e *XErrorBars) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// drawCap draws the cap if it is not clipped.
func (e *XErrorBars) drawCap(c *draw.Canvas, x, y vg.Length) { _ = "STUB: not implemented"; return }

// DataRange implements the plot.DataRanger interface.
func (e *XErrorBars) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes implements the plot.GlyphBoxer interface.
func (e *XErrorBars) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}
