// Copyright ©2013 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bezier implements 2D Bézier curve calculation.
package bezier // import "gonum.org/v1/plot/tools/bezier"

import "gonum.org/v1/plot/vg"

type point struct {
	Point, Control vg.Point
}

// Curve implements Bezier curve calculation according to the algorithm of Robert D. Miller.
//
// Graphics Gems 5, 'Quick and Simple Bézier Curve Drawing', pages 206-209.
type Curve []point

// NewCurve returns a Curve initialized with the control points in cp.
func New(cp ...vg.Point) Curve { _ = "STUB: not implemented"; return *new(Curve) }

// Point returns the point at t along the curve, where 0 ≤ t ≤ 1.
func (c Curve) Point(t float64) vg.Point { _ = "STUB: not implemented"; return *new(vg.Point) }

// Curve returns a slice of vg.Point, p, filled with points along the Bézier curve described by c.
// If the length of p is less than 2, the curve points are undefined. The length of p is not
// altered by the call.
func (c Curve) Curve(p []vg.Point) []vg.Point { _ = "STUB: not implemented"; return nil }
