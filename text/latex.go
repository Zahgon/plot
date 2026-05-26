// Copyright ©2020 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"image/color"

	"codeberg.org/go-latex/latex/drawtex"
	"codeberg.org/go-latex/latex/font/ttf"
	"codeberg.org/go-latex/latex/mtex"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
)

// Latex parses, formats and renders LaTeX.
type Latex struct {
	// Fonts is the cache of font faces used by this text handler.
	Fonts *font.Cache

	// DPI is the dot-per-inch controlling the font resolution used by LaTeX.
	// If zero, the resolution defaults to 72.
	DPI float64
}

var _ Handler = (*Latex)(nil)

// Cache returns the cache of fonts used by the text handler.
func (hdlr Latex) Cache() *font.Cache {
	_ = "STUB: not implemented"

	// Extents returns the Extents of a font.
	return nil
}

func (hdlr Latex) Extents(fnt font.Font) font.Extents {
	_ = "STUB: not implemented"
	return *new(font.Extents)
}

// Lines splits a given block of text into separate lines.
func (hdlr Latex) Lines(txt string) []string { _ = "STUB: not implemented"; return nil }

// Box returns the bounding box of the given non-multiline text where:
//   - width is the horizontal space from the origin.
//   - height is the vertical space above the baseline.
//   - depth is the vertical space below the baseline, a positive number.
func (hdlr Latex) Box(txt string, fnt font.Font) (width, height, depth vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length), *new(vg.Length)
}

// Add a bit of space, with a linegap as mtex.Box is returning
// a very tight bounding box.
// See gonum/plot#661.

// Draw renders the given text with the provided style and position
// on the canvas.
func (hdlr Latex) Draw(c vg.Canvas, txt string, sty Style, pt vg.Point) {
	_ = "STUB: not implemented"
	return
}

func (hdlr *Latex) fontsFor(fnt font.Font) *ttf.Fonts { _ = "STUB: not implemented"; return nil }

// latexDPI is the default LaTeX resolution used for computing the LaTeX
// layout of equations and regular text.
// Dimensions are then rescaled to the desired resolution.
const latexDPI = 72.0

func (hdlr Latex) dpi() float64 { _ = "STUB: not implemented"; return 0 }

type latex struct {
	cnv   vg.Canvas
	fonts *font.Cache
	sty   Style
	pt    vg.Point

	w vg.Length
	h vg.Length

	cos vg.Length
	sin vg.Length

	xoff vg.Length
	yoff vg.Length
}

var _ mtex.Renderer = (*latex)(nil)

func (r *latex) Render(width, height, dpi float64, c *drawtex.Canvas) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *latex) drawGlyph(dpi float64, op drawtex.GlyphOp) { _ = "STUB: not implemented"; return }

func (r *latex) drawRect(dpi float64, op drawtex.RectOp) { _ = "STUB: not implemented"; return }

func (r *latex) rotate(x, y vg.Length) (vg.Length, vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length)
}

// FillPolygon fills a polygon with the given color.
func fillPolygon(c vg.Canvas, clr color.Color, pts []vg.Point) { _ = "STUB: not implemented"; return }
