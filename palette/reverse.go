// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette

import (
	"image/color"
)

// Reverse reverses the direction of ColorMap c.
func Reverse(c ColorMap) ColorMap { _ = "STUB: not implemented"; return *new(ColorMap) }

// reverse is a ColorMap that reverses the direction of the ColorMap it
// contains.
type reverse struct {
	ColorMap
}

// At implements the ColorMap interface for a Reversed ColorMap.
func (r reverse) At(v float64) (color.Color, error) {
	_ = "STUB: not implemented"
	return *new(color.Color), nil
}

// Palette implements the ColorMap interface for a Reversed ColorMap.
func (r reverse) Palette(colors int) Palette { _ = "STUB: not implemented"; return *new(Palette) }
