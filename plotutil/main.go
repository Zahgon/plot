// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

var examples = []struct {
	name   string
	mkplot func() *plot.Plot
}{
	{"example_errpoints", Example_errpoints},
	{"example_stackedAreaChart", Example_stackedAreaChart},
}

func main() {
	for _, ex := range examples {
		drawEps(ex.name, ex.mkplot)
		drawSvg(ex.name, ex.mkplot)
		drawPng(ex.name, ex.mkplot)
		drawTiff(ex.name, ex.mkplot)
		drawJpg(ex.name, ex.mkplot)
		drawPdf(ex.name, ex.mkplot)
	}
}

func drawEps(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

func drawPdf(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

func drawSvg(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

func drawPng(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

func drawTiff(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

func drawJpg(name string, mkplot func() *plot.Plot) { _ = "STUB: not implemented"; return }

// Example_errpoints draws some error points.
func Example_errpoints() *plot.Plot { _ = "STUB: not implemented"; return nil }

// Get some random data.

type stackValues struct{ vs []plotter.Values }

func (n stackValues) Len() int            { _ = "STUB: not implemented"; return 0 }
func (n stackValues) Value(i int) float64 { _ = "STUB: not implemented"; return 0 }

// An example of making a stacked area chart.
func Example_stackedAreaChart() *plot.Plot { _ = "STUB: not implemented"; return nil }
