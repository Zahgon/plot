// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgpdf implements the vg.Canvas interface
// using gofpdf (github.com/phpdave11/gofpdf).
package vgpdf // import "gonum.org/v1/plot/vg/vgpdf"

import (
	_ "embed"
	"image"
	"image/color"
	"io"
	"sync"

	pdf "codeberg.org/go-pdf/fpdf"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// codePageEncoding holds informations about the characters encoding of TrueType
// font files, needed by gofpdf to embed fonts in a PDF document.
// We use cp1252 (code page 1252, Windows Western) to encode characters.
// See:
//   - https://en.wikipedia.org/wiki/Windows-1252
//
// TODO: provide a Canvas-level func option to embed fonts with a user provided
// code page schema?
//
//go:embed cp1252.map
var codePageEncoding []byte

func init() {
	draw.RegisterFormat("pdf", func(w, h vg.Length) vg.CanvasWriterTo {
		return New(w, h)
	})
}

// DPI is the nominal resolution of drawing in PDF.
const DPI = 72

// Canvas implements the vg.Canvas interface,
// drawing to a PDF.
type Canvas struct {
	doc  *pdf.Fpdf
	w, h vg.Length

	dpi       int
	numImages int
	stack     []context
	fonts     map[font.Font]struct{}

	// Switch to embed fonts in PDF file.
	// The default is to embed fonts.
	// This makes the PDF file more portable but also larger.
	embed bool
}

type context struct {
	fill  color.Color
	line  color.Color
	width vg.Length
}

// New creates a new PDF Canvas.
func New(w, h vg.Length) *Canvas { _ = "STUB: not implemented"; return nil }

// EmbedFonts specifies whether the resulting PDF canvas should
// embed the fonts or not.
// EmbedFonts returns the previous value before modification.
func (c *Canvas) EmbedFonts(v bool) bool { _ = "STUB: not implemented"; return false }

func (c *Canvas) DPI() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Canvas) context() *context { _ = "STUB: not implemented"; return nil }

func (c *Canvas) Size() (w, h vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length)
}

func (c *Canvas) SetLineWidth(w vg.Length) { _ = "STUB: not implemented"; return }

func (c *Canvas) SetLineDash(dashes []vg.Length, offs vg.Length) { _ = "STUB: not implemented"; return }

func (c *Canvas) SetColor(clr color.Color) { _ = "STUB: not implemented"; return }

func (c *Canvas) Rotate(r float64) { _ = "STUB: not implemented"; return }

func (c *Canvas) Translate(pt vg.Point) { _ = "STUB: not implemented"; return }

func (c *Canvas) Scale(x float64, y float64) { _ = "STUB: not implemented"; return }

func (c *Canvas) Push() { _ = "STUB: not implemented"; return }

func (c *Canvas) Pop() { _ = "STUB: not implemented"; return }

func (c *Canvas) Stroke(p vg.Path) { _ = "STUB: not implemented"; return }

func (c *Canvas) Fill(p vg.Path) { _ = "STUB: not implemented"; return }

func (c *Canvas) FillString(fnt font.Face, pt vg.Point, str string) {
	_ = "STUB: not implemented"
	return
}

// go-fpdf uses the top left corner as origin.

func (c *Canvas) sbounds(fnt font.Face, txt string) (left, top, right, bottom float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// not defined (standard font?), use average of 81%

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) { _ = "STUB: not implemented"; return }

// font registers a font and a size with the PDF canvas.
func (c *Canvas) font(fnt font.Face, pt vg.Point) { _ = "STUB: not implemented"; return }

// pdfPath processes a vg.Path and applies it to the canvas.
func (c *Canvas) pdfPath(path vg.Path, style string) { _ = "STUB: not implemented"; return }

func (c *Canvas) arc(comp vg.PathComp, style string) { _ = "STUB: not implemented"; return }

func (c *Canvas) pdfPointXY(x, y vg.Length) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (c *Canvas) pdfPoint(pt vg.Point) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// unit returns a fpdf.Unit, converted from a vg.Length.
func (c *Canvas) unit(l vg.Length) float64 { _ = "STUB: not implemented"; return 0 }

// imageName generates a unique image name for this PDF canvas
func (c *Canvas) imageName() string { _ = "STUB: not implemented"; return "" }

// WriterCounter implements the io.Writer interface, and counts
// the total number of bytes written.
type writerCounter struct {
	io.Writer
	n int64
}

func (w *writerCounter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteTo writes the Canvas to an io.Writer.
// After calling Write, the canvas is closed
// and may no longer be used for drawing.
func (c *Canvas) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// rgba converts a Go color into a gofpdf 3-tuple int + 1 float64
func rgba(c color.Color) (int, int, int, float64) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

type fontsCache struct {
	sync.RWMutex
	cache map[fontKey]fontVal
}

// fontKey represents a PDF font request.
// fontKey needs to know whether the font will be embedded or not,
// as gofpdf.MakeFont will generate different informations.
type fontKey struct {
	font  font.Face
	embed bool
}

type fontVal struct {
	z, j []byte
}

func (c *fontsCache) get(key fontKey) (fontVal, bool) {
	_ = "STUB: not implemented"
	return *new(fontVal), false
}

func (c *fontsCache) add(k fontKey, v fontVal) { _ = "STUB: not implemented"; return }

var pdfFonts = &fontsCache{
	cache: make(map[fontKey]fontVal),
}

func getFont(key fontKey, font, encoding []byte) (z, j []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func makeFont(key fontKey, font, encoding []byte) (val fontVal, err error) {
	_ = "STUB: not implemented"
	return *new(fontVal), nil
}

// NextPage creates a new page in the final PDF document.
// The new page is the new current page.
// Modifications applied to the canvas will only be applied to that new page.
func (c *Canvas) NextPage() { _ = "STUB: not implemented"; return }
