// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgeps implements the vg.Canvas interface using
// encapsulated postscript.
package vgeps // import "gonum.org/v1/plot/vg/vgeps"

import (
	"bytes"
	"image"
	"image/color"
	"io"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func init() {
	draw.RegisterFormat("eps", func(w, h vg.Length) vg.CanvasWriterTo {
		return New(w, h)
	})
}

// DPI is the nominal resolution of drawing in EPS.
const DPI = 72

type Canvas struct {
	stack []context
	w, h  vg.Length
	buf   *bytes.Buffer
}

type context struct {
	color  color.Color
	width  vg.Length
	dashes []vg.Length
	offs   vg.Length
	font   string
	fsize  vg.Length
}

// pr is the amount of precision to use when outputting float64s.
const pr = 5

// New returns a new Canvas.
func New(w, h vg.Length) *Canvas { _ = "STUB: not implemented"; return nil }

// NewTitle returns a new Canvas with the given title string.
func NewTitle(w, h vg.Length, title string) *Canvas { _ = "STUB: not implemented"; return nil }

func (c *Canvas) Size() (w, h vg.Length) {
	_ = "STUB: not implemented"

	// context returns the top context on the stack.
	return *new(vg.Length), *new(vg.Length)
}

func (e *Canvas) context() *context { _ = "STUB: not implemented"; return nil }

func (e *Canvas) SetLineWidth(w vg.Length) { _ = "STUB: not implemented"; return }

func (e *Canvas) SetLineDash(dashes []vg.Length, o vg.Length) { _ = "STUB: not implemented"; return }

func (e *Canvas) SetColor(c color.Color) { _ = "STUB: not implemented"; return }

func (e *Canvas) Rotate(r float64) { _ = "STUB: not implemented"; return }

func (e *Canvas) Translate(pt vg.Point) { _ = "STUB: not implemented"; return }

func (e *Canvas) Scale(x, y float64) { _ = "STUB: not implemented"; return }

func (e *Canvas) Push() { _ = "STUB: not implemented"; return }

func (e *Canvas) Pop() { _ = "STUB: not implemented"; return }

func (e *Canvas) Stroke(path vg.Path) { _ = "STUB: not implemented"; return }

func (e *Canvas) Fill(path vg.Path) { _ = "STUB: not implemented"; return }

func (e *Canvas) trace(path vg.Path) { _ = "STUB: not implemented"; return }

func (e *Canvas) FillString(fnt font.Face, pt vg.Point, str string) {
	_ = "STUB: not implemented"
	return
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) {
	_ = "STUB: not implemented"
	// FIXME: https://github.com/gonum/plot/issues/271
	return
}

// WriteTo writes the canvas to an io.Writer.
func (e *Canvas) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
