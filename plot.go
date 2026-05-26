// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"image/color"
	"io"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/font/liberation"
	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var (
	// DefaultFont is the name of the default font for plot text.
	DefaultFont = font.Font{
		Typeface: "Liberation",
		Variant:  "Serif",
	}

	// DefaultTextHandler is the default text handler used for text processing.
	DefaultTextHandler text.Handler
)

// Plot is the basic type representing a plot.
type Plot struct {
	Title struct {
		// Text is the text of the plot title.  If
		// Text is the empty string then the plot
		// will not have a title.
		Text string

		// Padding is the amount of padding
		// between the bottom of the title and
		// the top of the plot.
		Padding vg.Length

		// TextStyle specifies how the plot title text should be displayed.
		TextStyle text.Style
	}

	// BackgroundColor is the background color of the plot.
	// The default is White.
	BackgroundColor color.Color

	// X and Y are the horizontal and vertical axes
	// of the plot respectively.
	X, Y Axis

	// Legend is the plot's legend.
	Legend Legend

	// TextHandler parses and formats text according to a given
	// dialect (Markdown, LaTeX, plain, ...)
	// The default is a plain text handler.
	TextHandler text.Handler

	// plotters are drawn by calling their Plot method
	// after the axes are drawn.
	plotters []Plotter
}

// Plotter is an interface that wraps the Plot method.
// Some standard implementations of Plotter can be
// found in the gonum.org/v1/plot/plotter
// package, documented here:
// https://godoc.org/gonum.org/v1/plot/plotter
type Plotter interface {
	// Plot draws the data to a draw.Canvas.
	Plot(draw.Canvas, *Plot)
}

// DataRanger wraps the DataRange method.
type DataRanger interface {
	// DataRange returns the range of X and Y values.
	DataRange() (xmin, xmax, ymin, ymax float64)
}

// orientation describes whether an axis is horizontal or vertical.
type orientation byte

const (
	horizontal orientation = iota
	vertical
)

// New returns a new plot with some reasonable default settings.
func New() *Plot { _ = "STUB: not implemented"; return nil }

// Add adds a Plotters to the plot.
//
// If the plotters implements DataRanger then the
// minimum and maximum values of the X and Y
// axes are changed if necessary to fit the range of
// the data.
//
// When drawing the plot, Plotters are drawn in the
// order in which they were added to the plot.
func (p *Plot) Add(ps ...Plotter) { _ = "STUB: not implemented"; return }

// Draw draws a plot to a draw.Canvas.
//
// Plotters are drawn in the order in which they were
// added to the plot.  Plotters that  implement the
// GlyphBoxer interface will have their GlyphBoxes
// taken into account when padding the plot so that
// none of their glyphs are clipped.
func (p *Plot) Draw(c draw.Canvas) { _ = "STUB: not implemented"; return }

// DataCanvas returns a new draw.Canvas that
// is the subset of the given draw area into which
// the plot data will be drawn.
func (p *Plot) DataCanvas(da draw.Canvas) draw.Canvas {
	_ = "STUB: not implemented"
	return *new(draw.Canvas)
}

// DrawGlyphBoxes draws red outlines around the plot's
// GlyphBoxes.  This is intended for debugging.
func (p *Plot) DrawGlyphBoxes(c draw.Canvas) { _ = "STUB: not implemented"; return }

// padX returns a draw.Canvas that is padded horizontally
// so that glyphs will no be clipped.
func padX(p *Plot, c draw.Canvas) draw.Canvas { _ = "STUB: not implemented"; return *new(draw.Canvas) }

// rightMost returns the right-most GlyphBox.
func rightMost(c *draw.Canvas, boxes []GlyphBox) GlyphBox {
	_ = "STUB: not implemented"
	return *new(GlyphBox)
}

// leftMost returns the left-most GlyphBox.
func leftMost(c *draw.Canvas, boxes []GlyphBox) GlyphBox {
	_ = "STUB: not implemented"
	return *new(GlyphBox)
}

// padY returns a draw.Canvas that is padded vertically
// so that glyphs will no be clipped.
func padY(p *Plot, c draw.Canvas) draw.Canvas { _ = "STUB: not implemented"; return *new(draw.Canvas) }

// topMost returns the top-most GlyphBox.
func topMost(c *draw.Canvas, boxes []GlyphBox) GlyphBox {
	_ = "STUB: not implemented"
	return *new(GlyphBox)
}

// bottomMost returns the bottom-most GlyphBox.
func bottomMost(c *draw.Canvas, boxes []GlyphBox) GlyphBox {
	_ = "STUB: not implemented"
	return *new(GlyphBox)
}

// Transforms returns functions to transfrom
// from the x and y data coordinate system to
// the draw coordinate system of the given
// draw area.
func (p *Plot) Transforms(c *draw.Canvas) (x, y func(float64) vg.Length) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GlyphBoxer wraps the GlyphBoxes method.
// It should be implemented by things that meet
// the Plotter interface that draw glyphs so that
// their glyphs are not clipped if drawn near the
// edge of the draw.Canvas.
//
// When computing padding, the plot ignores
// GlyphBoxes as follows:
// If the Size.X > 0 and the X value is not in range
// of the X axis then the box is ignored.
// If Size.Y > 0 and the Y value is not in range of
// the Y axis then the box is ignored.
//
// Also, GlyphBoxes with Size.X <= 0 are ignored
// when computing horizontal padding and
// GlyphBoxes with Size.Y <= 0 are ignored when
// computing vertical padding.  This is useful
// for things like box plots and bar charts where
// the boxes and bars are considered to be glyphs
// in the X direction (and thus need padding), but
// may be clipped in the Y direction (and do not
// need padding).
type GlyphBoxer interface {
	GlyphBoxes(*Plot) []GlyphBox
}

// A GlyphBox describes the location of a glyph
// and the offset/size of its bounding box.
//
// If the Rectangle.Size().X is non-positive (<= 0) then
// the GlyphBox is ignored when computing the
// horizontal padding, and likewise with
// Rectangle.Size().Y and the vertical padding.
type GlyphBox struct {
	// The glyph location in normalized coordinates.
	X, Y float64

	// Rectangle is the offset of the glyph's minimum drawing
	// point relative to the glyph location and its size.
	vg.Rectangle
}

// GlyphBoxes returns the GlyphBoxes for all plot
// data that meet the GlyphBoxer interface.
func (p *Plot) GlyphBoxes(*Plot) (boxes []GlyphBox) { _ = "STUB: not implemented"; return nil }

// NominalX configures the plot to have a nominal X
// axis—an X axis with names instead of numbers.  The
// X location corresponding to each name are the integers,
// e.g., the x value 0 is centered above the first name and
// 1 is above the second name, etc.  Labels for x values
// that do not end up in range of the X axis will not have
// tick marks.
func (p *Plot) NominalX(names ...string) { _ = "STUB: not implemented"; return }

// HideX configures the X axis so that it will not be drawn.
func (p *Plot) HideX() { _ = "STUB: not implemented"; return }

// HideY configures the Y axis so that it will not be drawn.
func (p *Plot) HideY() { _ = "STUB: not implemented"; return }

// HideAxes hides the X and Y axes.
func (p *Plot) HideAxes() { _ = "STUB: not implemented"; return }

// NominalY is like NominalX, but for the Y axis.
func (p *Plot) NominalY(names ...string) { _ = "STUB: not implemented"; return }

// WriterTo returns an io.WriterTo that will write the plot as
// the specified image format.
//
// Supported formats are:
//
//   - .eps
//   - .jpg|.jpeg
//   - .pdf
//   - .png
//   - .svg
//   - .tex
//   - .tif|.tiff
func (p *Plot) WriterTo(w, h vg.Length, format string) (io.WriterTo, error) {
	_ = "STUB: not implemented"
	return *new(io.WriterTo), nil
}

// Save saves the plot to an image file.  The file format is determined
// by the extension.
//
// Supported extensions are:
//
//   - .eps
//   - .jpg|.jpeg
//   - .pdf
//   - .png
//   - .svg
//   - .tex
//   - .tif|.tiff
func (p *Plot) Save(w, h vg.Length, file string) (err error) { _ = "STUB: not implemented"; return nil }

func init() {
	font.DefaultCache.Add(liberation.Collection())
	DefaultTextHandler = text.Plain{
		Fonts: font.DefaultCache,
	}
}
