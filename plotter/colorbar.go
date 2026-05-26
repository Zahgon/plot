// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/palette"
	"gonum.org/v1/plot/vg/draw"
)

// ColorBar is a plot.Plotter that draws a color bar legend for a ColorMap.
type ColorBar struct {
	ColorMap palette.ColorMap

	// Vertical determines wether the legend will be
	// plotted vertically or horizontally.
	// The default is false (horizontal).
	Vertical bool

	// Colors specifies the number of colors to be
	// shown in the legend. If Colors is not specified,
	// a default will be used.
	Colors int
}

// colors returns the number of colors to be shown
// in the legend, substituting invalid values
// with the default of one color per point.
func (l *ColorBar) colors(c draw.Canvas) int { _ = "STUB: not implemented"; return 0 }

// check determines whether the ColorBar is
// valid in its current configuration.
func (l *ColorBar) check() { _ = "STUB: not implemented"; return }

// Plot implements the Plot method of the plot.Plotter interface.
func (l *ColorBar) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (l *ColorBar) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}
