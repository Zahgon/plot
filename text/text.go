// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"image/color"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
)

// Handler parses, formats and renders text.
type Handler interface {
	// Cache returns the cache of fonts used by the text handler.
	Cache() *font.Cache

	// Extents returns the Extents of a font.
	Extents(fnt font.Font) font.Extents

	// Lines splits a given block of text into separate lines.
	Lines(txt string) []string

	// Box returns the bounding box of the given non-multiline text where:
	//  - width is the horizontal space from the origin.
	//  - height is the vertical space above the baseline.
	//  - depth is the vertical space below the baseline, a positive number.
	Box(txt string, fnt font.Font) (width, height, depth vg.Length)

	// Draw renders the given text with the provided style and position
	// on the canvas.
	Draw(c vg.Canvas, txt string, sty Style, pt vg.Point)
}

// XAlignment specifies text alignment in the X direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using XAlignment(desired number).
type XAlignment float64

const (
	// XLeft aligns the left edge of the text with the specified location.
	XLeft XAlignment = 0
	// XCenter aligns the horizontal center of the text with the specified location.
	XCenter XAlignment = -0.5
	// XRight aligns the right edge of the text with the specified location.
	XRight XAlignment = -1
)

// YAlignment specifies text alignment in the Y direction. Three preset
// options are available, but an arbitrary alignment
// can also be specified using YAlignment(desired number).
type YAlignment float64

const (
	// YTop aligns the top of of the text with the specified location.
	YTop YAlignment = -1
	// YCenter aligns the vertical center of the text with the specified location.
	YCenter YAlignment = -0.5
	// YBottom aligns the bottom of the text with the specified location.
	YBottom YAlignment = 0
)

// Position specifies the text position.
const (
	PosLeft   = -1
	PosBottom = -1
	PosCenter = 0
	PosTop    = +1
	PosRight  = +1
)

// Style describes what text will look like.
type Style struct {
	// Color is the text color.
	Color color.Color

	// Font is the font description.
	Font font.Font

	// Rotation is the text rotation in radians, performed around the axis
	// defined by XAlign and YAlign.
	Rotation float64

	// XAlign and YAlign specify the alignment of the text.
	XAlign XAlignment
	YAlign YAlignment

	// Handler parses and formats text according to a given
	// dialect (Markdown, LaTeX, plain, ...)
	// The default is a plain text handler.
	Handler Handler
}

// FontExtents returns the extents of this Style's font.
func (s Style) FontExtents() font.Extents { _ = "STUB: not implemented"; return *new(font.Extents) }

// Width returns the width of lines of text
// when using the given font before any text rotation is applied.
func (s Style) Width(txt string) (max vg.Length) { _ = "STUB: not implemented"; return *new(vg.Length) }

// Height returns the height of the text when using
// the given font before any text rotation is applied.
func (s Style) Height(txt string) vg.Length { _ = "STUB: not implemented"; return *new(vg.Length) }

// box returns the bounding box of a possibly multi-line text.
func (s Style) box(txt string) (w, h vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length)
}

// Rectangle returns a rectangle giving the bounds of
// this text assuming that it is drawn at (0, 0).
func (s Style) Rectangle(txt string) vg.Rectangle {
	_ = "STUB: not implemented"
	return *new(vg.Rectangle)
}

// descent + linegap

// lower left corner

// upper left corner

// lower right corner

// upper right corner

// rotatePoint applies rotation theta (in radians) about the origin to point p.
func rotatePoint(theta float64, p vg.Point) vg.Point {
	_ = "STUB: not implemented"
	return *new(vg.Point)
}

func max(d ...vg.Length) vg.Length { _ = "STUB: not implemented"; return *new(vg.Length) }

func min(d ...vg.Length) vg.Length { _ = "STUB: not implemented"; return *new(vg.Length) }
