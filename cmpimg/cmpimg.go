// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmpimg compares the raw representation of images taking into account
// idiosyncracies related to their underlying format (SVG, PDF, PNG, ...).
package cmpimg // import "gonum.org/v1/plot/cmpimg"

import (
	"image"
	"image/color"
	"image/draw"

	"rsc.io/pdf"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"
)

// Equal takes the raw representation of two images, raw1 and raw2,
// together with the underlying image type ("eps", "jpeg", "jpg", "pdf", "png", "svg", "tiff"),
// and returns whether the two images are equal or not.
//
// Equal may return an error if the decoding of the raw image somehow failed.
func Equal(typ string, raw1, raw2 []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// EqualApprox takes the raw representation of two images, raw1 and raw2,
// together with the underlying image type ("eps", "jpeg", "jpg", "pdf", "png", "svg", "tiff"),
// a normalized delta parameter to describe how close the matching should be
// performed (delta=0: perfect match, delta=1, loose match)
// and returns whether the two images are equal or not.
//
// EqualApprox may return an error if the decoding of the raw image somehow failed.
// EqualApprox only uses the normalized delta parameter for "jpeg", "jpg", "png",
// and "tiff" images. It ignores that parameter for other document types.
func EqualApprox(typ string, raw1, raw2 []byte, delta float64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func cmpPdf(pdf1, pdf2 *pdf.Reader) bool { _ = "STUB: not implemented"; return false }

func cmpImg(v1, v2 image.Image, delta float64) bool { _ = "STUB: not implemented"; return false }

// yiqEqApprox compares the colors of 2 pixels, in the NTSC YIQ color space,
// as described in:
//
//	Measuring perceived color difference using YIQ NTSC
//	transmission color space in mobile applications.
//	Yuriy Kotsarenko, Fernando Ramos.
//
// An electronic version is available at:
//
// - http://www.progmat.uaem.mx:8080/artVol2Num2/Articulo3Vol2Num2.pdf
func yiqEqApprox(c1, c2 color.RGBA, d2 float64) bool {
	_ = "STUB: not implemented"
	// difference between 2 maximally different pixels.
	return false
}

func newRGBAFrom(src image.Image) *image.RGBA { _ = "STUB: not implemented"; return nil }

// Diff calculates an intensity-scaled difference between images a and b
// and places the result in dst, returning the intersection of a, b and
// dst. It is the responsibility of the caller to construct dst so that
// it will overlap with a and b. For the purposes of Diff, alpha is not
// considered.
//
// Diff is not intended to be used for quantitative analysis of the
// difference between the input images, but rather to highlight differences
// between them for testing purposes, so the calculation is rather naive.
func Diff(dst draw.Image, a, b image.Image) image.Rectangle {
	_ = "STUB: not implemented"
	return *new(image.Rectangle)
}

// Determine greyscale dynamic range.

// Render intensity-scaled difference.

type diffColor struct {
	a, b color.Color
}

func (c diffColor) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func diff(a, b uint32) uint32 { _ = "STUB: not implemented"; return 0 }

type scaledColor struct {
	min, max uint32
	c        color.Color
}

func (c scaledColor) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }
