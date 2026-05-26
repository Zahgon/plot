// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package moreland

import (
	"image/color"
)

// rgb represents a physically linear RGB color.
type rgb struct {
	R, G, B float64
}

// cieXYZ returns a CIE XYZ color representation of the receiver.
func (c rgb) cieXYZ() cieXYZ { _ = "STUB: not implemented"; return *new(cieXYZ) }

// sRGBA returns an sRGB color representation of the receiver using the
// provided alpha which must be in [0, 1].
func (c rgb) sRGBA(alpha float64) sRGBA {
	_ = "STUB: not implemented"
	// f converts from a linear RGB component to an sRGB component.
	return *new(sRGBA)
}

// cieXYZ represents a color in CIE XYZ space.
// Y must be in the range [0,1]. X and Z must be greater than 0.
type cieXYZ struct {
	X, Y, Z float64
}

// rgb returns a linear RGB representation of the receiver.
func (c cieXYZ) rgb() rgb { _ = "STUB: not implemented"; return *new(rgb) }

// cieLAB returns a CIELAB color representation of the receiver.
func (c cieXYZ) cieLAB() cieLAB {
	_ = "STUB: not implemented"
	// f is an intermediate step in converting from CIE XYZ to CIE LAB.
	return *new(cieLAB)
}

// sRGBA represents a color within the sRGB color space, with an alpha channel
// but not premultiplied. All values must be in the range [0,1].
type sRGBA struct {
	R, G, B, A float64
}

// rgb returns a linear RGB representation of the receiver.
func (c sRGBA) rgb() rgb {
	_ = "STUB: not implemented"
	// f converts from an sRGB component to a linear RGB component.
	return *new(rgb)
}

// RGBA implements the color.Color interface.
func (c sRGBA) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

// cieLAB returns a CIE LAB representation of the receiver.
func (c sRGBA) cieLAB() cieLAB { _ = "STUB: not implemented"; return *new(cieLAB) }

// colorTosRGBA converts a color to an sRGBA.
func colorTosRGBA(c color.Color) sRGBA { _ = "STUB: not implemented"; return *new(sRGBA) }

// clamp forces all channels in c to be within the range [0, 1].
func (c *sRGBA) clamp() { _ = "STUB: not implemented"; return }

// cieLAB represents a color in CIE LAB space.
// L must be in the range [0, 100].
type cieLAB struct {
	L, A, B float64
}

// sRGBA return a linear RGB color representation of the receiver using the
// provided alpha which must be in [0, 1].
func (c cieLAB) sRGBA(alpha float64) sRGBA { _ = "STUB: not implemented"; return *new(sRGBA) }

// cieXYZ returns a CIE XYZ color representation of the receiver.
func (c cieLAB) cieXYZ() cieXYZ {
	_ = "STUB: not implemented"
	// f is an intermediate step in converting from CIE LAB to CIE XYZ.
	return *new(cieXYZ)
}

// Reference white-point D65

// MSH returns an MSH color representation of the receiver.
func (c cieLAB) MSH() msh { _ = "STUB: not implemented"; return *new(msh) }

// MSH represents a color in Magnitude-Saturation-Hue color space.
type msh struct {
	M, S, H float64
}

// colorToMSH converts a color to MSH space.
// TODO: If msh ever becomes exported, change this to implment color.Model
func colorToMSH(c color.Color) msh { _ = "STUB: not implemented"; return *new(msh) }

// cieLAB returns a CIELAB representation of the receiver.
func (c msh) cieLAB() cieLAB { _ = "STUB: not implemented"; return *new(cieLAB) }

// RGBA implements the color.Color interface.
func (c msh) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

// hueTwist returns the hue twist between color c and converge magnitude
// convergeM.
func hueTwist(c msh, convergeM float64) float64 { _ = "STUB: not implemented"; return 0 }
