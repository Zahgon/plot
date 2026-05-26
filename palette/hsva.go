// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright ©2011-2013 The bíogo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package palette

import (
	"image/color"
)

// HSVA represents a Hue/Saturation/Value/Alpha color.
// H, S, V and A are valid within [0, 1].
type HSVA struct {
	H, S, V, A float64
}

// HSVAModel converts any color.Color to an HSVA color.
var HSVAModel = color.ModelFunc(hsvaModel)

func hsvaModel(c color.Color) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

// Convert r, g, b, a to HSVA
func rgbaToHsva(r, g, b, a uint32) HSVA { _ = "STUB: not implemented"; return *new(HSVA) }

// This should really be math.NaN() since we have a 0 length vector,
// but 0 seems to be the convention and it may simplify imports in
// dependent packages.

// RGBA allows HSVAColor to satisfy the color.Color interface.
func (c HSVA) RGBA() (r, g, b, a uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }
