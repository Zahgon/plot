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
	// DefaultGridLineStyle is the default style for grid lines.
	DefaultGridLineStyle = draw.LineStyle{
		Color: color.Gray{128},
		Width: vg.Points(0.25),
	}
)

// Grid implements the plot.Plotter interface, drawing
// a set of grid lines at the major tick marks.
type Grid struct {
	// Vertical is the style of the vertical lines.
	Vertical draw.LineStyle

	// Horizontal is the style of the horizontal lines.
	Horizontal draw.LineStyle
}

// NewGrid returns a new grid with both vertical and
// horizontal lines using the default grid line style.
func NewGrid() *Grid { _ = "STUB: not implemented"; return nil }

// Plot implements the plot.Plotter interface.
func (g *Grid) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }
