// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vg

import (
	"image"
	"image/color"

	"gonum.org/v1/plot/font"
)

// MultiCanvas creates a canvas that duplicates its drawing operations to all
// the provided canvases, similar to the Unix tee(1) command.
//
// Each drawing operation is sent to each listed canvas, one at a time.
func MultiCanvas(cs ...Canvas) Canvas { _ = "STUB: not implemented"; return *new(Canvas) }

type teeCanvas struct {
	cs []Canvas
}

// SetLineWidth sets the width of stroked paths.
// If the width is not positive then stroked lines
// are not drawn.
func (tee teeCanvas) SetLineWidth(w Length) { _ = "STUB: not implemented"; return }

// SetLineDash sets the dash pattern for lines.
// The pattern slice specifies the lengths of
// alternating dashes and gaps, and the offset
// specifies the distance into the dash pattern
// to start the dash.
func (tee teeCanvas) SetLineDash(pattern []Length, offset Length) {
	_ = "STUB: not implemented"
	return
}

// SetColor sets the current drawing color.
// Note that fill color and stroke color are
// the same, so if you want different fill
// and stroke colors then you must set a color,
// draw fills, set a new color and then draw lines.
func (tee teeCanvas) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

// Rotate applies a rotation transform to the context.
// The parameter is specified in radians.
func (tee teeCanvas) Rotate(rad float64) { _ = "STUB: not implemented"; return }

// Translate applies a translational transform to the context.
func (tee teeCanvas) Translate(pt Point) { _ = "STUB: not implemented"; return }

// Scale applies a scaling transform to the context.
func (tee teeCanvas) Scale(x, y float64) { _ = "STUB: not implemented"; return }

// Push saves the current line width, the
// current dash pattern, the current
// transforms, and the current color
// onto a stack so that the state can later
// be restored by calling Pop().
func (tee teeCanvas) Push() { _ = "STUB: not implemented"; return }

// Pop restores the context saved by the
// corresponding call to Push().
func (tee teeCanvas) Pop() { _ = "STUB: not implemented"; return }

// Stroke strokes the given path.
func (tee teeCanvas) Stroke(p Path) { _ = "STUB: not implemented"; return }

// Fill fills the given path.
func (tee teeCanvas) Fill(p Path) { _ = "STUB: not implemented"; return }

// FillString fills in text at the specified
// location using the given font.
// If the font size is zero, the text is not drawn.
func (tee teeCanvas) FillString(f font.Face, pt Point, text string) {
	_ = "STUB: not implemented"
	return
}

// DrawImage draws the image, scaled to fit
// the destination rectangle.
func (tee teeCanvas) DrawImage(rect Rectangle, img image.Image) { _ = "STUB: not implemented"; return }

var _ Canvas = (*teeCanvas)(nil)
