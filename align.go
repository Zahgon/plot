// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"gonum.org/v1/plot/vg/draw"
)

// Align returns a two-dimensional row-major array of Canvases which will
// produce tiled plots with DataCanvases that are evenly sized and spaced.
// The arguments to the function are a two-dimensional row-major array
// of plots, a tile configuration, and the canvas to which the tiled
// plots are to be drawn.
func Align(plots [][]*Plot, t draw.Tiles, dc draw.Canvas) [][]draw.Canvas {
	_ = "STUB: not implemented"
	return nil
}

// Create the initial tiles.

// Calculate the maximum spacing between data canvases
// for each row and column.

// Calculate the total row and column spacing.

// Adjust the horizontal and vertical spacing between
// canvases to match the maximum for each column and row,
// respectively.

// Adjust the canvas so that the height and width of the
// DataCanvas is the same for all plots.
