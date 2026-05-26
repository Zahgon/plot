// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotutil

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type combineXYs struct{ xs, ys plotter.Valuer }

func (c combineXYs) Len() int                    { _ = "STUB: not implemented"; return 0 }
func (c combineXYs) XY(i int) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

type item struct {
	name  string
	value plot.Thumbnailer
}

// AddStackedAreaPlots adds stacked area plot plotters to a plot.
// The variadic arguments must be either strings
// or plotter.Valuers.  Each valuer adds a stacked area
// plot to the plot below the stacked area plots added
// before it.  If a plotter.Valuer is immediately
// preceeded by a string then the string value is used to
// label the legend.
// Plots should be added in order of tallest to shortest,
// because they will be drawn in the order they are added
// (i.e. later plots will be painted over earlier plots).
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddStackedAreaPlots(plt *plot.Plot, xs plotter.Valuer, vs ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Make a line plotter and set its style.

// AddBoxPlots adds box plot plotters to a plot and
// sets the X axis of the plot to be nominal.
// The variadic arguments must be either strings
// or plotter.Valuers.  Each valuer adds a box plot
// to the plot at the X location corresponding to
// the number of box plots added before it.  If a
// plotter.Valuer is immediately preceeded by a
// string then the string value is used to label the
// tick mark for the box plot's X location.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddBoxPlots(plt *plot.Plot, width vg.Length, vs ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// AddScatters adds Scatter plotters to a plot.
// The variadic arguments must be either strings
// or plotter.XYers.  Each plotter.XYer is added to
// the plot using the next color, and glyph shape
// via the Color and Shape functions. If a
// plotter.XYer is immediately preceeded by
// a string then a legend entry is added to the plot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddScatters(plt *plot.Plot, vs ...any) error { _ = "STUB: not implemented"; return nil }

// AddLines adds Line plotters to a plot.
// The variadic arguments must be a string
// or one of a plotting type, plotter.XYers or *plotter.Function.
// Each plotting type is added to
// the plot using the next color and dashes
// shape via the Color and Dashes functions.
// If a plotting type is immediately preceeded by
// a string then a legend entry is added to the plot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddLines(plt *plot.Plot, vs ...any) error { _ = "STUB: not implemented"; return nil }

// AddLinePoints adds Line and Scatter plotters to a
// plot.  The variadic arguments must be either strings
// or plotter.XYers.  Each plotter.XYer is added to
// the plot using the next color, dashes, and glyph
// shape via the Color, Dashes, and Shape functions.
// If a plotter.XYer is immediately preceeded by
// a string then a legend entry is added to the plot
// using the string as the name.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddLinePoints(plt *plot.Plot, vs ...any) error { _ = "STUB: not implemented"; return nil }

// AddErrorBars adds XErrorBars and YErrorBars
// to a plot.  The variadic arguments must be
// of type plotter.XYer, and must be either a
// plotter.XErrorer, plotter.YErrorer, or both.
// Each errorer is added to the plot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddErrorBars(plt *plot.Plot, vs ...any) error { _ = "STUB: not implemented"; return nil }

// AddXErrorBars adds XErrorBars to a plot.
// The variadic arguments must be
// of type plotter.XYer, and plotter.XErrorer.
// Each errorer is added to the plot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddXErrorBars(plt *plot.Plot, es ...interface {
	plotter.XYer
	plotter.XErrorer
}) error {
	_ = "STUB: not implemented"
	return nil
}

// AddYErrorBars adds YErrorBars to a plot.
// The variadic arguments must be
// of type plotter.XYer, and plotter.YErrorer.
// Each errorer is added to the plot the color from
// the Colors function corresponding to its position
// in the argument list.
//
// If an error occurs then none of the plotters are added
// to the plot, and the error is returned.
func AddYErrorBars(plt *plot.Plot, es ...interface {
	plotter.XYer
	plotter.YErrorer
}) error {
	_ = "STUB: not implemented"
	return nil
}
