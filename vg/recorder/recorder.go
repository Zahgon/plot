// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package recorder provides support for vector graphics serialization.
package recorder // import "gonum.org/v1/plot/vg/recorder"

import (
	"image"
	"image/color"

	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/vg"
)

var _ vg.Canvas = (*Canvas)(nil)

// Canvas implements vg.Canvas operation serialization.
type Canvas struct {
	// Actions holds a log of all methods called on
	// the canvas.
	Actions []Action

	// KeepCaller indicates whether the Canvas will
	// retain runtime caller location for the actions.
	// This includes source filename and line number.
	KeepCaller bool

	// fonts holds a collection of font/size descriptions.
	fonts map[fontID]font.Face
	cache *font.Cache
}

type fontID struct {
	name string
	size vg.Length
}

// Action is a vector graphics action as defined by the
// vg.Canvas interface. Each method of Canvas has a
// corresponding Action type.
type Action interface {
	Call() string
	ApplyTo(vg.Canvas)
	callerLocation() *callerLocation
}

type callerLocation struct {
	haveCaller bool
	file       string
	line       int
}

func (l *callerLocation) set() { _ = "STUB: not implemented"; return }

func (l callerLocation) String() string { _ = "STUB: not implemented"; return "" }

// Reset resets the Canvas to the base state.
func (c *Canvas) Reset() { _ = "STUB: not implemented"; return }

// ReplayOn applies the set of Actions recorded by the Canvas onto
// the destination Canvas.
func (c *Canvas) ReplayOn(dst vg.Canvas) error { _ = "STUB: not implemented"; return nil }

func (c *Canvas) append(a Action) { _ = "STUB: not implemented"; return }

// SetLineWidth corresponds to the vg.Canvas.SetWidth method.
type SetLineWidth struct {
	Width vg.Length

	l callerLocation
}

// SetLineWidth implements the SetLineWidth method of the vg.Canvas interface.
func (c *Canvas) SetLineWidth(w vg.Length) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *SetLineWidth) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *SetLineWidth) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *SetLineWidth) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// SetLineDash corresponds to the vg.Canvas.SetLineDash method.
	return nil
}

type SetLineDash struct {
	Dashes  []vg.Length
	Offsets vg.Length

	l callerLocation
}

// SetLineDash implements the SetLineDash method of the vg.Canvas interface.
func (c *Canvas) SetLineDash(dashes []vg.Length, offs vg.Length) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *SetLineDash) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *SetLineDash) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *SetLineDash) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// SetColor corresponds to the vg.Canvas.SetColor method.
	return nil
}

type SetColor struct {
	Color color.Color

	l callerLocation
}

// SetColor implements the SetColor method of the vg.Canvas interface.
func (c *Canvas) SetColor(col color.Color) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *SetColor) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *SetColor) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *SetColor) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Rotate corresponds to the vg.Canvas.Rotate method.
	return nil
}

type Rotate struct {
	Angle float64

	l callerLocation
}

// Rotate implements the Rotate method of the vg.Canvas interface.
func (c *Canvas) Rotate(a float64) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Rotate) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Rotate) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Rotate) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Translate corresponds to the vg.Canvas.Translate method.
	return nil
}

type Translate struct {
	Point vg.Point

	l callerLocation
}

// Translate implements the Translate method of the vg.Canvas interface.
func (c *Canvas) Translate(pt vg.Point) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Translate) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Translate) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Translate) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Scale corresponds to the vg.Canvas.Scale method.
	return nil
}

type Scale struct {
	X, Y float64

	l callerLocation
}

// Scale implements the Scale method of the vg.Canvas interface.
func (c *Canvas) Scale(x, y float64) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Scale) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Scale) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Scale) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Push corresponds to the vg.Canvas.Push method.
	return nil
}

type Push struct {
	l callerLocation
}

// Push implements the Push method of the vg.Canvas interface.
func (c *Canvas) Push() {
	_ = "STUB: not implemented"

	// Call returns the method call that generated the action.
	return
}

func (a *Push) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Push) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Push) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Pop corresponds to the vg.Canvas.Pop method.
	return nil
}

type Pop struct {
	l callerLocation
}

// Pop implements the Pop method of the vg.Canvas interface.
func (c *Canvas) Pop() {
	_ = "STUB: not implemented"

	// Call returns the method call that generated the action.
	return
}

func (a *Pop) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Pop) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Pop) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Stroke corresponds to the vg.Canvas.Stroke method.
	return nil
}

type Stroke struct {
	Path vg.Path

	l callerLocation
}

// Stroke implements the Stroke method of the vg.Canvas interface.
func (c *Canvas) Stroke(path vg.Path) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Stroke) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Stroke) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Stroke) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Fill corresponds to the vg.Canvas.Fill method.
	return nil
}

type Fill struct {
	Path vg.Path

	l callerLocation
}

// Fill implements the Fill method of the vg.Canvas interface.
func (c *Canvas) Fill(path vg.Path) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Fill) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Fill) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Fill) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// FillString corresponds to the vg.Canvas.FillString method.
	return nil
}

type FillString struct {
	Font   font.Font
	Size   vg.Length
	Point  vg.Point
	String string

	l callerLocation

	fonts map[fontID]font.Face
}

// FillString implements the FillString method of the vg.Canvas interface.
func (c *Canvas) FillString(font font.Face, pt vg.Point, str string) {
	_ = "STUB: not implemented"
	return
}

// ApplyTo applies the action to the given vg.Canvas.
func (a *FillString) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

// Call returns the pseudo method call that generated the action.
func (a *FillString) Call() string { _ = "STUB: not implemented"; return "" }

func (a *FillString) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// DrawImage corresponds to the vg.Canvas.DrawImage method
	return nil
}

type DrawImage struct {
	Rectangle vg.Rectangle
	Image     image.Image

	l callerLocation
}

// DrawImage implements the DrawImage method of the vg.Canvas interface.
func (c *Canvas) DrawImage(rect vg.Rectangle, img image.Image) { _ = "STUB: not implemented"; return }

// ApplyTo applies the action to the given vg.Canvas.
func (a *DrawImage) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

// Call returns the pseudo method call that generated the action.
func (a *DrawImage) Call() string { _ = "STUB: not implemented"; return "" }

func (a *DrawImage) callerLocation() *callerLocation {
	_ = "STUB: not implemented"

	// Commenter defines types that can record comments.
	return nil
}

type Commenter interface {
	Comment(string)
}

var _ Commenter = (*Canvas)(nil)

// Comment implements a comment mechanism.
type Comment struct {
	Text string

	l callerLocation
}

// Comment adds a comment to a list of Actions..
func (c *Canvas) Comment(text string) { _ = "STUB: not implemented"; return }

// Call returns the method call that generated the action.
func (a *Comment) Call() string { _ = "STUB: not implemented"; return "" }

// ApplyTo applies the action to the given vg.Canvas.
func (a *Comment) ApplyTo(c vg.Canvas) { _ = "STUB: not implemented"; return }

func (a *Comment) callerLocation() *callerLocation { _ = "STUB: not implemented"; return nil }
