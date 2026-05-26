// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vg

// A Point is a location in 2d space.
//
// Points are used for drawing, not for data.  For
// data, see the XYer interface.
type Point struct {
	X, Y Length
}

// Dot returns the dot product of two points.
func (p Point) Dot(q Point) Length {
	_ = "STUB: not implemented"
	return *

	// Add returns the component-wise sum of two points.
	new(Length)
}

func (p Point) Add(q Point) Point { _ = "STUB: not implemented"; return *new(Point) }

// Sub returns the component-wise difference of two points.
func (p Point) Sub(q Point) Point { _ = "STUB: not implemented"; return *new(Point) }

// Scale returns the component-wise product of a point and a scalar.
func (p Point) Scale(s Length) Point { _ = "STUB: not implemented"; return *new(Point) }

// A Rectangle represents a rectangular region of 2d space.
type Rectangle struct {
	Min Point
	Max Point
}

// Size returns the width and height of a Rectangle.
func (r Rectangle) Size() Point { _ = "STUB: not implemented"; return *new(Point) }

// Add returns the rectangle r translated by p.
func (r Rectangle) Add(p Point) Rectangle { _ = "STUB: not implemented"; return *new(Rectangle) }

// Path returns the path of a Rect specified by its
// upper left corner, width and height.
func (r Rectangle) Path() (p Path) { _ = "STUB: not implemented"; return *new(Path) }
