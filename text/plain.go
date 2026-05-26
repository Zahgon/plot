// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text // import "gonum.org/v1/plot/text"

import (
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
)

// Plain is a text/plain handler.
type Plain struct {
	Fonts *font.Cache
}

var _ Handler = (*Plain)(nil)

// Cache returns the cache of fonts used by the text handler.
func (hdlr Plain) Cache() *font.Cache {
	_ = "STUB: not implemented"

	// Extents returns the Extents of a font.
	return nil
}

func (hdlr Plain) Extents(fnt font.Font) font.Extents {
	_ = "STUB: not implemented"
	return *new(font.Extents)
}

// Lines splits a given block of text into separate lines.
func (hdlr Plain) Lines(txt string) []string { _ = "STUB: not implemented"; return nil }

// Box returns the bounding box of the given non-multiline text where:
//   - width is the horizontal space from the origin.
//   - height is the vertical space above the baseline.
//   - depth is the vertical space below the baseline, a positive number.
func (hdlr Plain) Box(txt string, fnt font.Font) (width, height, depth vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length), *new(vg.Length)
}

// Draw renders the given text with the provided style and position
// on the canvas.
func (hdlr Plain) Draw(c vg.Canvas, txt string, sty Style, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}
