// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg/draw"
)

// Image is a plotter that draws a scaled, raster image.
type Image struct {
	img            image.Image
	cols           int
	rows           int
	xmin, xmax, dx float64
	ymin, ymax, dy float64
}

// NewImage creates a new image plotter.
// Image will plot img inside the rectangle defined by the
// (xmin, ymin) and (xmax, ymax) points given in the data space.
// The img will be scaled to fit inside the rectangle.
func NewImage(img image.Image, xmin, ymin, xmax, ymax float64) *Image {
	_ = "STUB: not implemented"
	return nil
}

// Plot implements the Plot method of the plot.Plotter interface.
func (img *Image) Plot(c draw.Canvas, p *plot.Plot) { _ = "STUB: not implemented"; return }

// DataRange implements the DataRange method
// of the plot.DataRanger interface.
func (img *Image) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes implements the GlyphBoxes method
// of the plot.GlyphBoxer interface.
func (img *Image) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox {
	_ = "STUB: not implemented"

	// transform warps the image to align with non-linear axes.
	return nil
}

func (img *Image) transformFor(p *plot.Plot) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

// Find the equivalent image column after applying axis transforms.

// Find the equivalent column of the previous image column after applying
// axis transforms.

// Find the equivalent image row after applying axis transforms.

// Find the equivalent row of the previous image row after applying
// axis transforms.

// Set all the pixels in the new image between (cPrevTrans, rPrevTrans)
// and (cTrans, rTrans) to the color at (c,r) in the original image.
// TODO: Improve interpolation.

func (img *Image) x(c int) float64 { _ = "STUB: not implemented"; return 0 }

func (img *Image) y(r int) float64 { _ = "STUB: not implemented"; return 0 }

// uniform is a cropped uniform image.
type uniform struct {
	*image.Uniform
	rect image.Rectangle
}

func (img uniform) Bounds() image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}
