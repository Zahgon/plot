// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

// GlyphBoxes implements the Plotter interface, drawing
// all of the glyph boxes of the plot.  This is intended for
// debugging.
type GlyphBoxes struct {
	draw.LineStyle
}

func NewGlyphBoxes() *GlyphBoxes { _ = "STUB: not implemented"; return nil }

func (g GlyphBoxes) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }
