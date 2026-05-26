// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"time"

	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// Ticker creates Ticks in a specified range
type Ticker interface {
	// Ticks returns Ticks in a specified range
	Ticks(min, max float64) []Tick
}

// Normalizer rescales values from the data coordinate system to the
// normalized coordinate system.
type Normalizer interface {
	// Normalize transforms a value x in the data coordinate system to
	// the normalized coordinate system.
	Normalize(min, max, x float64) float64
}

// An Axis represents either a horizontal or vertical
// axis of a plot.
type Axis struct {
	// Min and Max are the minimum and maximum data
	// values represented by the axis.
	Min, Max float64

	Label struct {
		// Text is the axis label string.
		Text string

		// Padding is the distance between the label and the axis.
		Padding vg.Length

		// TextStyle is the style of the axis label text.
		// For the vertical axis, one quarter turn
		// counterclockwise will be added to the label
		// text before drawing.
		TextStyle text.Style

		// Position is where the axis label string should be drawn.
		// The default value is draw.PosCenter, displaying the label
		// at the center of the axis.
		// Valid values are [-1,+1], with +1 being the far right/top
		// of the axis, and -1 the far left/bottom of the axis.
		Position float64
	}

	// LineStyle is the style of the axis line.
	draw.LineStyle

	// Padding between the axis line and the data.  Having
	// non-zero padding ensures that the data is never drawn
	// on the axis, thus making it easier to see.
	Padding vg.Length

	Tick struct {
		// Label is the TextStyle on the tick labels.
		Label text.Style

		// LineStyle is the LineStyle of the tick lines.
		draw.LineStyle

		// Length is the length of a major tick mark.
		// Minor tick marks are half of the length of major
		// tick marks.
		Length vg.Length

		// Marker returns the tick marks.  Any tick marks
		// returned by the Marker function that are not in
		// range of the axis are not drawn.
		Marker Ticker
	}

	// Scale transforms a value given in the data coordinate system
	// to the normalized coordinate system of the axis—its distance
	// along the axis as a fraction of the axis range.
	Scale Normalizer

	// AutoRescale enables an axis to automatically adapt its minimum
	// and maximum boundaries, according to its underlying Ticker.
	AutoRescale bool
}

// makeAxis returns a default Axis.
//
// The default range is (∞, ­∞), and thus any finite
// value is less than Min and greater than Max.
func makeAxis(o orientation) Axis { _ = "STUB: not implemented"; return *new(Axis) }

// sanitizeRange ensures that the range of the
// axis makes sense.
func (a *Axis) sanitizeRange() { _ = "STUB: not implemented"; return }

// LinearScale an be used as the value of an Axis.Scale function to
// set the axis to a standard linear scale.
type LinearScale struct{}

var _ Normalizer = LinearScale{}

// Normalize returns the fractional distance of x between min and max.
func (LinearScale) Normalize(min, max, x float64) float64 { _ = "STUB: not implemented"; return 0 }

// LogScale can be used as the value of an Axis.Scale function to
// set the axis to a log scale.
type LogScale struct{}

var _ Normalizer = LogScale{}

// Normalize returns the fractional logarithmic distance of
// x between min and max.
func (LogScale) Normalize(min, max, x float64) float64 { _ = "STUB: not implemented"; return 0 }

// InvertedScale can be used as the value of an Axis.Scale function to
// invert the axis using any Normalizer.
type InvertedScale struct{ Normalizer }

var _ Normalizer = InvertedScale{}

// Normalize returns a normalized [0, 1] value for the position of x.
func (is InvertedScale) Normalize(min, max, x float64) float64 { _ = "STUB: not implemented"; return 0 }

// Norm returns the value of x, given in the data coordinate
// system, normalized to its distance as a fraction of the
// range of this axis.  For example, if x is a.Min then the return
// value is 0, and if x is a.Max then the return value is 1.
func (a Axis) Norm(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// drawTicks returns true if the tick marks should be drawn.
func (a Axis) drawTicks() bool { _ = "STUB: not implemented"; return false }

// A horizontalAxis draws horizontally across the bottom
// of a plot.
type horizontalAxis struct {
	Axis
}

// size returns the height of the axis.
func (a horizontalAxis) size() (h vg.Length) {
	_ = "STUB: not implemented"
	return *
	// We assume that the label isn't rotated.
	new(vg.Length)
}

// draw draws the axis along the lower edge of a draw.Canvas.
func (a horizontalAxis) draw(c draw.Canvas) { _ = "STUB: not implemented"; return }

// GlyphBoxes returns the GlyphBoxes for the tick labels.
func (a horizontalAxis) GlyphBoxes(p *Plot) []GlyphBox { _ = "STUB: not implemented"; return nil }

// FIXME(sbinet): want data coordinates

// A verticalAxis is drawn vertically up the left side of a plot.
type verticalAxis struct {
	Axis
}

// size returns the width of the axis.
func (a verticalAxis) size() (w vg.Length) {
	_ = "STUB: not implemented"
	return *
	// We assume that the label isn't rotated.
	new(vg.Length)
}

// draw draws the axis along the left side of a draw.Canvas.
func (a verticalAxis) draw(c draw.Canvas) { _ = "STUB: not implemented"; return }

// GlyphBoxes returns the GlyphBoxes for the tick labels
func (a verticalAxis) GlyphBoxes(p *Plot) []GlyphBox { _ = "STUB: not implemented"; return nil }

// FIXME(sbinet): want data coordinates

// descent + linegap

// DefaultTicks is suitable for the Tick.Marker field of an Axis,
// it returns a reasonable default set of tick marks.
type DefaultTicks struct{}

var _ Ticker = DefaultTicks{}

// Ticks returns Ticks in the specified range.
func (DefaultTicks) Ticks(min, max float64) []Tick { _ = "STUB: not implemented"; return nil }

// Simple fall back was chosen, so
// majorDelta is the label distance.

// Choose a reasonable, but ad
// hoc formatting for labels.

// See talbotLinHanrahan for the values used here.

// Find the first minor tick not greater
// than the lowest data value.

// Add ticks at minorDelta intervals when
// they are not within minorDelta/2 of a
// labelled tick.

func minInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func maxInt(a, b int) int {
	_ = "STUB: not implemented"

	// LogTicks is suitable for the Tick.Marker field of an Axis,
	// it returns tick marks suitable for a log-scale axis.
	return 0
}

type LogTicks struct {
	// Prec specifies the precision of tick rendering
	// according to the documentation for strconv.FormatFloat.
	Prec int
}

var _ Ticker = LogTicks{}

// Ticks returns Ticks in a specified range
func (t LogTicks) Ticks(min, max float64) []Tick { _ = "STUB: not implemented"; return nil }

// ConstantTicks is suitable for the Tick.Marker field of an Axis.
// This function returns the given set of ticks.
type ConstantTicks []Tick

var _ Ticker = ConstantTicks{}

// Ticks returns Ticks in a specified range
func (ts ConstantTicks) Ticks(float64, float64) []Tick {
	_ = "STUB: not implemented"

	// UnixTimeIn returns a time conversion function for the given location.
	return nil
}

func UnixTimeIn(loc *time.Location) func(t float64) time.Time {
	_ = "STUB: not implemented"
	return nil
}

// UTCUnixTime is the default time conversion for TimeTicks.
var UTCUnixTime = UnixTimeIn(time.UTC)

// TimeTicks is suitable for axes representing time values.
type TimeTicks struct {
	// Ticker is used to generate a set of ticks.
	// If nil, DefaultTicks will be used.
	Ticker Ticker

	// Format is the textual representation of the time value.
	// If empty, time.RFC3339 will be used
	Format string

	// Time takes a float64 value and converts it into a time.Time.
	// If nil, UTCUnixTime is used.
	Time func(t float64) time.Time
}

var _ Ticker = TimeTicks{}

// Ticks implements plot.Ticker.
func (t TimeTicks) Ticks(min, max float64) []Tick { _ = "STUB: not implemented"; return nil }

// A Tick is a single tick mark on an axis.
type Tick struct {
	// Value is the data value marked by this Tick.
	Value float64

	// Label is the text to display at the tick mark.
	// If Label is an empty string then this is a minor
	// tick mark.
	Label string
}

// IsMinor returns true if this is a minor tick mark.
func (t Tick) IsMinor() bool { _ = "STUB: not implemented"; return false }

// lengthOffset returns an offset that should be added to the
// tick mark's line to accout for its length.  I.e., the start of
// the line for a minor tick mark must be shifted by half of
// the length.
func (t Tick) lengthOffset(len vg.Length) vg.Length {
	_ = "STUB: not implemented"
	return *new(vg.Length)
}

// tickLabelHeight returns height of the tick mark labels.
func tickLabelHeight(sty text.Style, ticks []Tick) vg.Length {
	_ = "STUB: not implemented"
	return *new(vg.Length)
}

// tickLabelWidth returns the width of the widest tick mark label.
func tickLabelWidth(sty text.Style, ticks []Tick) vg.Length {
	_ = "STUB: not implemented"
	return *new(vg.Length)
}

// formatFloatTick returns a g-formated string representation of v
// to the specified precision.
func formatFloatTick(v float64, prec int) string { _ = "STUB: not implemented"; return "" }

// TickerFunc is suitable for the Tick.Marker field of an Axis.
// It is an adapter which allows to quickly setup a Ticker using a function with an appropriate signature.
type TickerFunc func(min, max float64) []Tick

var _ Ticker = TickerFunc(nil)

// Ticks implements plot.Ticker.
func (f TickerFunc) Ticks(min, max float64) []Tick { _ = "STUB: not implemented"; return nil }
