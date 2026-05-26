// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmpimg

import (
	"flag"
	"testing"
)

var GenerateTestData = flag.Bool("regen", false, "Uses the current state to regenerate the test data.")

func goldenPath(path string) string { _ = "STUB: not implemented"; return "" }

// CheckPlot checks a generated plot against a previously created reference.
// If GenerateTestData = true, it regenerates the reference.
// For image.Image formats, a base64 encoded png representation is output to
// the testing log when a difference is identified.
func CheckPlot(ExampleFunc func(), t *testing.T, filenames ...string) {
	_ = "STUB: not implemented"
	return
}

// CheckPlotApprox checks a generated plot against a previously created reference.
// The normalized delta parameter describes how tight the matching should be
// performed, where delta=0 expresses a perfect match, and delta=1 a very loose match.
// If GenerateTestData = true, it regenerates the reference.
// For image.Image formats, a base64 encoded png representation is output to
// the testing log when a difference is identified.
func CheckPlotApprox(ExampleFunc func(), t *testing.T, delta float64, filenames ...string) {
	_ = "STUB: not implemented"
	return
}

// Recreate Golden images and exit.

// Run the example.

// Read the images we've just generated and check them against the
// Golden Images.

// remove the dot in e.g. ".pdf"
