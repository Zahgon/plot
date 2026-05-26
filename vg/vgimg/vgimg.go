// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vgimg implements the vg.Canvas interface using
// git.sr.ht/~sbinet/gg as a backend to output raster images.
package vgimg // import "gonum.org/v1/plot/vg/vgimg"

import (
	"image"
	"image/color"
	"image/draw"
	"io"

	"git.sr.ht/~sbinet/gg"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
	vgdraw "gonum.org/v1/plot/vg/draw"
)

func init() {
	vgdraw.RegisterFormat("png", func(w, h vg.Length) vg.CanvasWriterTo {
		return PngCanvas{Canvas: New(w, h)}
	})

	vgdraw.RegisterFormat("jpg", func(w, h vg.Length) vg.CanvasWriterTo {
		return JpegCanvas{Canvas: New(w, h)}
	})

	vgdraw.RegisterFormat("jpeg", func(w, h vg.Length) vg.CanvasWriterTo {
		return JpegCanvas{Canvas: New(w, h)}
	})

	vgdraw.RegisterFormat("tif", func(w, h vg.Length) vg.CanvasWriterTo {
		return TiffCanvas{Canvas: New(w, h)}
	})

	vgdraw.RegisterFormat("tiff", func(w, h vg.Length) vg.CanvasWriterTo {
		return TiffCanvas{Canvas: New(w, h)}
	})
}

// Canvas implements the vg.Canvas interface,
// drawing to an image.Image using draw2d.
type Canvas struct {
	ctx   *gg.Context
	img   draw.Image
	w, h  vg.Length
	color []color.Color

	// dpi is the number of dots per inch for this canvas.
	dpi int

	// width is the current line width.
	width vg.Length

	// backgroundColor is the background color, set by
	// UseBackgroundColor.
	backgroundColor color.Color
}

const (
	// DefaultDPI is the default dot resolution for image
	// drawing in dots per inch.
	DefaultDPI = 96

	// DefaultWidth and DefaultHeight are the default canvas
	// dimensions.
	DefaultWidth  = 4 * vg.Inch
	DefaultHeight = 4 * vg.Inch
)

// New returns a new image canvas.
func New(w, h vg.Length) *Canvas { _ = "STUB: not implemented"; return nil }

// NewWith returns a new image canvas created according to the specified
// options. The currently accepted options are UseWH,
// UseDPI, UseImage, and UseImageWithContext.
// Each of the options specifies the size of the canvas (UseWH, UseImage),
// the resolution of the canvas (UseDPI), or both (useImageWithContext).
// If size or resolution are not specified, defaults are used.
// It panics if size and resolution are overspecified (i.e., too many options are
// passed).
func NewWith(o ...option) *Canvas { _ = "STUB: not implemented"; return nil }

// h should also == 0.

// These constants are used to ensure that the options
// used when initializing a canvas are compatible with
// each other.
const (
	setsDPI uint32 = 1 << iota
	setsSize
	setsBackground
)

type option func(*Canvas) uint32

// UseWH specifies the width and height of the canvas.
// The size is rounded up to the nearest pixel.
func UseWH(w, h vg.Length) option { _ = "STUB: not implemented"; return *new(option) }

// UseDPI sets the dots per inch of a canvas. It should only be
// used as an option argument when initializing a new canvas.
func UseDPI(dpi int) option { _ = "STUB: not implemented"; return *new(option) }

// UseImage specifies an image to create
// the canvas from. The
// minimum point of the given image
// should probably be 0,0.
//
// Note that a copy of the input image is performed.
// This means that modifications applied to the canvas are not reflected
// on the original image.
func UseImage(img draw.Image) option { _ = "STUB: not implemented"; return *new(option) }

// UseImageWithContext specifies both an image
// and a graphic context to create the canvas from.
// The minimum point of the given image
// should probably be 0,0.
func UseImageWithContext(img draw.Image, ctx *gg.Context) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// UseBackgroundColor specifies the image background color.
// Without UseBackgroundColor, the default color is white.
func UseBackgroundColor(c color.Color) option { _ = "STUB: not implemented"; return *new(option) }

// Image returns the image the canvas is drawing to.
//
// The dimensions of the returned image must not be modified.
func (c *Canvas) Image() draw.Image { _ = "STUB: not implemented"; return *new(draw.Image) }

func (c *Canvas) Size() (w, h vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length), *new(vg.Length)
}

func (c *Canvas) SetLineWidth(w vg.Length) { _ = "STUB: not implemented"; return }

func (c *Canvas) SetLineDash(ds []vg.Length, offs vg.Length) { _ = "STUB: not implemented"; return }

func (c *Canvas) SetColor(clr color.Color) { _ = "STUB: not implemented"; return }

func (c *Canvas) Rotate(t float64) { _ = "STUB: not implemented"; return }

func (c *Canvas) Translate(pt vg.Point) { _ = "STUB: not implemented"; return }

func (c *Canvas) Scale(x, y float64) { _ = "STUB: not implemented"; return }

func (c *Canvas) Push() { _ = "STUB: not implemented"; return }

func (c *Canvas) Pop() { _ = "STUB: not implemented"; return }

func (c *Canvas) Stroke(p vg.Path) { _ = "STUB: not implemented"; return }

func (c *Canvas) Fill(p vg.Path) { _ = "STUB: not implemented"; return }

func (c *Canvas) outline(p vg.Path) { _ = "STUB: not implemented"; return }

// DPI returns the resolution of the receiver in pixels per inch.
func (c *Canvas) DPI() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Canvas) FillString(font font.Face, pt vg.Point, str string) {
	_ = "STUB: not implemented"
	return
}

// DrawImage implements the vg.Canvas.DrawImage method.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) { _ = "STUB: not implemented"; return }

// WriterCounter implements the io.Writer interface, and counts
// the total number of bytes written.
type writerCounter struct {
	io.Writer
	n int64
}

func (w *writerCounter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// A JpegCanvas is an image canvas with a WriteTo method
// that writes a jpeg image.
type JpegCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a jpeg image.
func (c JpegCanvas) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// A PngCanvas is an image canvas with a WriteTo method that
// writes a png image.
type PngCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a png image.
func (c PngCanvas) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// A TiffCanvas is an image canvas with a WriteTo method that
// writes a tiff image.
type TiffCanvas struct {
	*Canvas
}

// WriteTo implements the io.WriterTo interface, writing a tiff image.
func (c TiffCanvas) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
