// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package draw // import "gonum.org/v1/plot/vg/draw"

import (
	"image/color"
	"sync"

	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
)

// formats holds the registered canvas image formats
var formats = struct {
	sync.RWMutex
	m map[string]func(w, h vg.Length) vg.CanvasWriterTo
}{
	m: make(map[string]func(w, h vg.Length) vg.CanvasWriterTo),
}

// Formats returns the sorted list of registered vg formats.
func Formats() []string { _ = "STUB: not implemented"; return nil }

// RegisterFormat registers an image format for use by NewFormattedCanvas.
// name is the name of the format, like "jpeg" or "png".
// fn is the construction function to call for the format.
//
// RegisterFormat panics if fn is nil.
func RegisterFormat(name string, fn func(w, h vg.Length) vg.CanvasWriterTo) {
	_ = "STUB: not implemented"
	return
}

// A Canvas is a vector graphics canvas along with
// an associated Rectangle defining a section of the canvas
// to which drawing should take place.
type Canvas struct {
	vg.Canvas
	vg.Rectangle
}

// XAlignment specifies text alignment in the X direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using XAlignment(desired number).
type XAlignment = text.XAlignment

const (
	// XLeft aligns the left edge of the text with the specified location.
	XLeft = text.XLeft
	// XCenter aligns the horizontal center of the text with the specified location.
	XCenter = text.XCenter
	// XRight aligns the right edge of the text with the specified location.
	XRight = text.XRight
)

// YAlignment specifies text alignment in the Y direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using YAlignment(desired number).
type YAlignment = text.YAlignment

const (
	// YTop aligns the top of of the text with the specified location.
	YTop = text.YTop
	// YCenter aligns the vertical center of the text with the specified location.
	YCenter = text.YCenter
	// YBottom aligns the bottom of the text with the specified location.
	YBottom = text.YBottom
)

// Position specifies the text position.
const (
	PosLeft   = text.PosLeft
	PosBottom = text.PosBottom
	PosCenter = text.PosCenter
	PosTop    = text.PosTop
	PosRight  = text.PosRight
)

// LineStyle describes what a line will look like.
type LineStyle struct {
	// Color is the color of the line.
	Color color.Color

	// Width is the width of the line.
	Width vg.Length

	Dashes   []vg.Length
	DashOffs vg.Length
}

// A GlyphStyle specifies the look of a glyph used to draw
// a point on a plot.
type GlyphStyle struct {
	// Color is the color used to draw the glyph.
	color.Color

	// Radius specifies the size of the glyph's radius.
	Radius vg.Length

	// Shape draws the shape of the glyph.
	Shape GlyphDrawer
}

// A GlyphDrawer wraps the DrawGlyph function.
type GlyphDrawer interface {
	// DrawGlyph draws the glyph at the given
	// point, with the given color and radius.
	DrawGlyph(*Canvas, GlyphStyle, vg.Point)
}

// DrawGlyph draws the given glyph to the draw
// area.  If the point is not within the Canvas
// or the sty.Shape is nil then nothing is drawn.
func (c *Canvas) DrawGlyph(sty GlyphStyle, pt vg.Point) { _ = "STUB: not implemented"; return }

// DrawGlyphNoClip draws the given glyph to the draw
// area.  If the sty.Shape is nil then nothing is drawn.
func (c *Canvas) DrawGlyphNoClip(sty GlyphStyle, pt vg.Point) { _ = "STUB: not implemented"; return }

// Rectangle returns the rectangle surrounding this glyph,
// assuming that it is drawn centered at 0,0
func (g GlyphStyle) Rectangle() vg.Rectangle { _ = "STUB: not implemented"; return *new(vg.Rectangle) }

// CircleGlyph is a glyph that draws a solid circle.
type CircleGlyph struct{}

// DrawGlyph implements the GlyphDrawer interface.
func (CircleGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// RingGlyph is a glyph that draws the outline of a circle.
type RingGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (RingGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

const (
	cosπover4 = vg.Length(.707106781202420)
	sinπover6 = vg.Length(.500000000025921)
	cosπover6 = vg.Length(.866025403769473)
)

// SquareGlyph is a glyph that draws the outline of a square.
type SquareGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (SquareGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// BoxGlyph is a glyph that draws a filled square.
type BoxGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (BoxGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// TriangleGlyph is a glyph that draws the outline of a triangle.
type TriangleGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (TriangleGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// PyramidGlyph is a glyph that draws a filled triangle.
type PyramidGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (PyramidGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// PlusGlyph is a glyph that draws a plus sign
type PlusGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (PlusGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// CrossGlyph is a glyph that draws a big X.
type CrossGlyph struct{}

// DrawGlyph implements the Glyph interface.
func (CrossGlyph) DrawGlyph(c *Canvas, sty GlyphStyle, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

// New returns a new (bounded) draw.Canvas.
func New(c vg.CanvasSizer) Canvas { _ = "STUB: not implemented"; return *new(Canvas) }

// NewFormattedCanvas creates a new vg.CanvasWriterTo with the specified
// image format. Supported formats need to be registered by importing one or
// more of the following packages:
//
//   - gonum.org/v1/plot/vg/vgeps: provides eps
//   - gonum.org/v1/plot/vg/vgimg: provides png, jpg|jpeg, tif|tiff
//   - gonum.org/v1/plot/vg/vgpdf: provides pdf
//   - gonum.org/v1/plot/vg/vgsvg: provides svg
//   - gonum.org/v1/plot/vg/vgtex: provides tex
func NewFormattedCanvas(w, h vg.Length, format string) (vg.CanvasWriterTo, error) {
	_ = "STUB: not implemented"
	return *new(vg.CanvasWriterTo), nil
}

// NewCanvas returns a new (bounded) draw.Canvas of the given size.
func NewCanvas(c vg.Canvas, w, h vg.Length) Canvas { _ = "STUB: not implemented"; return *new(Canvas) }

// Center returns the center point of the area
func (c *Canvas) Center() vg.Point { _ = "STUB: not implemented"; return *new(vg.Point) }

// Contains returns true if the Canvas contains the point.
func (c *Canvas) Contains(p vg.Point) bool { _ = "STUB: not implemented"; return false }

// ContainsX returns true if the Canvas contains the
// x coordinate.
func (c *Canvas) ContainsX(x vg.Length) bool { _ = "STUB: not implemented"; return false }

// ContainsY returns true if the Canvas contains the
// y coordinate.
func (c *Canvas) ContainsY(y vg.Length) bool { _ = "STUB: not implemented"; return false }

// X returns the value of x, given in the unit range,
// in the drawing coordinates of this draw area.
// A value of 0, for example, will return the minimum
// x value of the draw area and a value of 1 will
// return the maximum.
func (c *Canvas) X(x float64) vg.Length { _ = "STUB: not implemented"; return *new(vg.Length) }

// Y returns the value of x, given in the unit range,
// in the drawing coordinates of this draw area.
// A value of 0, for example, will return the minimum
// y value of the draw area and a value of 1 will
// return the maximum.
func (c *Canvas) Y(y float64) vg.Length { _ = "STUB: not implemented"; return *new(vg.Length) }

// Crop returns a new Canvas corresponding to the Canvas
// c with the given lengths added to the minimum
// and maximum x and y values of the Canvas's Rectangle.
// Note that cropping the right and top sides of the canvas
// requires specifying negative values of right and top.
func Crop(c Canvas, left, right, bottom, top vg.Length) Canvas {
	_ = "STUB: not implemented"
	return *new(Canvas)
}

// Tiles creates regular subcanvases from a Canvas.
type Tiles struct {
	// Cols and Rows specify the number of rows and columns of tiles.
	Cols, Rows int
	// PadTop, PadBottom, PadRight, and PadLeft specify the padding
	// on the corresponding side of each tile.
	PadTop, PadBottom, PadRight, PadLeft vg.Length
	// PadX and PadY specify the padding between columns and rows
	// of tiles respectively..
	PadX, PadY vg.Length
}

// At returns the subcanvas within c that corresponds to the
// tile at column x, row y.
func (ts Tiles) At(c Canvas, x, y int) Canvas { _ = "STUB: not implemented"; return *new(Canvas) }

// SetLineStyle sets the current line style
func (c *Canvas) SetLineStyle(sty LineStyle) { _ = "STUB: not implemented"; return }

// StrokeLines draws a line connecting a set of points
// in the given Canvas.
func (c *Canvas) StrokeLines(sty LineStyle, lines ...[]vg.Point) { _ = "STUB: not implemented"; return }

// StrokeLine2 draws a line between two points in the given
// Canvas.
func (c *Canvas) StrokeLine2(sty LineStyle, x0, y0, x1, y1 vg.Length) {
	_ = "STUB: not implemented"
	return
}

// ClipLinesXY returns a slice of lines that
// represent the given line clipped in both
// X and Y directions.
func (c *Canvas) ClipLinesXY(lines ...[]vg.Point) [][]vg.Point {
	_ = "STUB: not implemented"
	return nil
}

// ClipLinesX returns a slice of lines that
// represent the given line clipped in the
// X direction.
func (c *Canvas) ClipLinesX(lines ...[]vg.Point) (clipped [][]vg.Point) {
	_ = "STUB: not implemented"
	return nil
}

// ClipLinesY returns a slice of lines that
// represent the given line clipped in the
// Y direction.
func (c *Canvas) ClipLinesY(lines ...[]vg.Point) (clipped [][]vg.Point) {
	_ = "STUB: not implemented"
	return nil
}

// clipLine performs clipping of a line by a single
// clipping line specified by the norm, clip point,
// and in function.
func clipLine(in func(vg.Point, vg.Point) bool, clip, norm vg.Point, pts []vg.Point) (lines [][]vg.Point) {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

// !curIn && nextIn

// FillPolygon fills a polygon with the given color.
func (c *Canvas) FillPolygon(clr color.Color, pts []vg.Point) { _ = "STUB: not implemented"; return }

// ClipPolygonXY returns a slice of lines that
// represent the given polygon clipped in both
// X and Y directions.
func (c *Canvas) ClipPolygonXY(pts []vg.Point) []vg.Point { _ = "STUB: not implemented"; return nil }

// ClipPolygonX returns a slice of lines that
// represent the given polygon clipped in the
// X direction.
func (c *Canvas) ClipPolygonX(pts []vg.Point) []vg.Point { _ = "STUB: not implemented"; return nil }

// ClipPolygonY returns a slice of lines that
// represent the given polygon clipped in the
// Y direction.
func (c *Canvas) ClipPolygonY(pts []vg.Point) []vg.Point { _ = "STUB: not implemented"; return nil }

// clipPoly performs clipping of a polygon by a single
// clipping line specified by the norm, clip point,
// and in function.
func clipPoly(in func(vg.Point, vg.Point) bool, clip, norm vg.Point, pts []vg.Point) (clipped []vg.Point) {
	_ = "STUB: not implemented"
	return nil
}

// do nothing

// !curIn && nextIn

// slop is some slop for floating point equality
const slop = 3e-8 // ≈ √1⁻¹⁵

func isLeft(p, clip vg.Point) bool { _ = "STUB: not implemented"; return false }

func isRight(p, clip vg.Point) bool { _ = "STUB: not implemented"; return false }

func isBelow(p, clip vg.Point) bool { _ = "STUB: not implemented"; return false }

func isAbove(p, clip vg.Point) bool { _ = "STUB: not implemented"; return false }

// isect returns the intersection of a line p0→p1 with the
// clipping line specified by the clip point and normal.
func isect(p0, p1, clip, norm vg.Point) vg.Point {
	_ = "STUB: not implemented"
	// t = (norm · (p0 - clip)) / (norm · (p0 - p1))
	return *new(vg.Point)
}

// p = p0 + t*(p1 - p0)

// FillText fills lines of text in the draw area.
// pt specifies the location where the text is to be drawn.
func (c *Canvas) FillText(sty TextStyle, pt vg.Point, txt string) {
	_ = "STUB: not implemented"
	return
}
