// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var (
	// DefaultFont is the default font for label text.
	DefaultFont = plot.DefaultFont

	// DefaultFontSize is the default font.
	DefaultFontSize = vg.Points(10)
)

// Labels implements the Plotter interface,
// drawing a set of labels at specified points.
type Labels struct {
	XYs

	// Labels is the set of labels corresponding
	// to each point.
	Labels []string

	// TextStyle is the style of the label text. Each label
	// can have a different text style.
	TextStyle []text.Style

	// Offset is added directly to the final label location.
	Offset vg.Point
}

// NewLabels returns a new Labels using the DefaultFont and
// the DefaultFontSize.
func NewLabels(d XYLabeller) (*Labels, error) { _ = "STUB: not implemented"; return nil, nil }

// Plot implements the Plotter interface, drawing labels.
func (l *Labels) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange returns the minimum and maximum X and Y values
func (l *Labels) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"

	// GlyphBoxes returns a slice of GlyphBoxes,
	// one for each of the labels, implementing the
	// plot.GlyphBoxer interface.
	return 0, 0, 0, 0
}

func (l *Labels) GlyphBoxes(p *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }

// XYLabeller combines the XYer and Labeller types.
type XYLabeller interface {
	XYer
	Labeller
}

// XYLabels holds XY data with labels.
// The ith label corresponds to the ith XY.
type XYLabels struct {
	XYs
	Labels []string
}

// Label returns the label for point index i.
func (l XYLabels) Label(i int) string { _ = "STUB: not implemented"; return "" }

var _ XYLabeller = (*XYLabels)(nil)
