// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package liberation exports the Liberation fonts as a font.Collection.
package liberation // import "gonum.org/v1/plot/font/liberation"

import (
	"sync"

	"gonum.org/v1/plot/font"
)

var (
	once       sync.Once
	collection font.Collection
)

func Collection() font.Collection { _ = "STUB: not implemented"; return *new(font.Collection) }

// mono variant

// sans-serif variant

func addColl(fnt font.Font, ttf []byte) { _ = "STUB: not implemented"; return }
