// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

// Scatter implements the Plotter interface, drawing
// a glyph for each of a set of points.
type Scatter struct {
	// XYs is a copy of the points for this scatter.
	XYs

	// GlyphStyleFunc, if not nil, specifies GlyphStyles
	// for individual points
	GlyphStyleFunc func(int) draw.GlyphStyle

	// GlyphStyle is the style of the glyphs drawn
	// at each point.
	draw.GlyphStyle
}

// NewScatter returns a Scatter that uses the
// default glyph style.
func NewScatter(xys XYer) (*Scatter, error) { _ = "STUB: not implemented"; return nil, nil }

// Plot draws the Scatter, implementing the plot.Plotter
// interface.
func (pts *Scatter) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum
// x and y values, implementing the plot.DataRanger
// interface.
func (pts *Scatter) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0,

		// GlyphBoxes returns a slice of plot.GlyphBoxes,
		// implementing the plot.GlyphBoxer interface.
		0, 0, 0
}

func (pts *Scatter) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"
	return nil
}

// Thumbnail the thumbnail for the Scatter,
// implementing the plot.Thumbnailer interface.
func (pts *Scatter) Thumbnail(c *draw.Canvas) { _ = "STUB: not implemented"; return }
