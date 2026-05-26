// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotutil

import (
	"gonum.org/v1/plot/plotter"
)

// ErrorPoints holds a set of x, y pairs along
// with their X and Y errors.
type ErrorPoints struct {
	plotter.XYs
	plotter.XErrors
	plotter.YErrors
}

// NewErrorPoints returns a new ErrorPoints where each
// point in the ErrorPoints is given by evaluating the
// center function on the Xs and Ys for the corresponding
// set of XY values in the pts parameter.  The XError
// and YError are computed likewise, using the err
// function.
//
// This function can be useful for summarizing sets of
// scatter points using a single point and error bars for
// each element of the scatter.
func NewErrorPoints(f func([]float64) (c, l, h float64), pts ...plotter.XYer) (*ErrorPoints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MeanAndConf95 returns the mean
// and the magnitude of the 95% confidence
// interval on the mean as low and high
// error values.
//
// MeanAndConf95 may be used as
// the f argument to NewErrorPoints.
func MeanAndConf95(vls []float64) (mean, lowerr, higherr float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// MedianAndMinMax returns the median
// value and error on the median given
// by the minimum and maximum data
// values.
//
// MedianAndMinMax may be used as
// the f argument to NewErrorPoints.
func MedianAndMinMax(vls []float64) (med, lowerr, higherr float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}
