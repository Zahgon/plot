// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.12
// +build go1.12

package plot

const root = "gonum.org/v1/plot"

// Version returns the version of Gonum/plot and its checksum. The returned
// values are only valid in binaries built with module support.
//
// If a replace directive exists in the Gonum/plot go.mod, the replace will
// be reported in the version in the following format:
//
//	"version=>[replace-path] [replace-version]"
//
// and the replace sum will be returned in place of the original sum.
//
// The exact version format returned by Version may change in future.
func Version() (version, sum string) { _ = "STUB: not implemented"; return "", "" }
