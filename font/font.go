// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package font

import (
	"sync"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// DefaultCache is the global cache for fonts.
var DefaultCache *Cache = NewCache(nil)

// Font represents a font face.
type Font struct {
	// Typeface identifies the Font.
	Typeface Typeface

	// TODO(sbinet): Gio@v0.2.0 has dropped font.Font.Variant
	// we should probably follow suit.

	// Variant is the variant of a font, such as "Mono" or "Smallcaps".
	Variant Variant

	// Style is the style of a font, such as Regular or Italic.
	Style font.Style

	// Weight is the weight of a font, such as Normal or Bold.
	Weight font.Weight

	// Size is the size of the font.
	Size Length
}

// Name returns a fully qualified name for the given font.
func (f *Font) Name() string { _ = "STUB: not implemented"; return "" }

// From returns a copy of the provided font with its size set.
func From(fnt Font, size Length) Font { _ = "STUB: not implemented"; return *new(Font) }

// Typeface identifies a particular typeface design.
// The empty string denotes the default typeface.
type Typeface string

// Variant denotes a typeface variant, such as "Mono", "Smallcaps" or "Math".
type Variant string

// Extents contains font metric information.
type Extents struct {
	// Ascent is the distance that the text
	// extends above the baseline.
	Ascent Length

	// Descent is the distance that the text
	// extends below the baseline. The descent
	// is given as a positive value.
	Descent Length

	// Height is the distance from the lowest
	// descending point to the highest ascending
	// point.
	Height Length
}

// Face holds a font descriptor and the associated font face.
type Face struct {
	Font Font
	Face *opentype.Font
}

// Name returns a fully qualified name for the given font.
func (f *Face) Name() string { _ = "STUB: not implemented"; return "" }

// FontFace returns the opentype font face for the requested
// dots-per-inch resolution.
func (f *Face) FontFace(dpi float64) font.Face { _ = "STUB: not implemented"; return *new(font.Face) }

// default hinting for OpenType fonts
const defaultHinting = font.HintingNone

// Extents returns the FontExtents for a font.
func (f *Face) Extents() Extents {
	_ = "STUB: not implemented"

	// TODO(sbinet): re-use a Font-level sfnt.Buffer instead?
	return *new(Extents)
}

// Width returns width of a string when drawn using the font.
func (f *Face) Width(s string) Length { _ = "STUB: not implemented"; return *new(Length) }

// scale converts sfnt.Unit to float64

// no-op

// Collection is a collection of fonts, regrouped under a common typeface.
type Collection []Face

// Cache collects font faces.
type Cache struct {
	mu    sync.RWMutex
	def   Typeface
	faces map[Font]*opentype.Font
}

// We make Cache implement dummy GobDecoder and GobEncoder interfaces
// to allow plot.Plot (or any other type holding a Cache) to be (de)serialized
// with encoding/gob.
// As Cache holds opentype.Font, the reflect-based gob (de)serialization can not
// work: gob isn't happy with opentype.Font having no exported field:
//
//   error: gob: type font.Cache has no exported fields
//
// FIXME(sbinet): perhaps encode/decode Cache.def typeface?

func (c *Cache) GobEncode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
func (c *Cache) GobDecode([]byte) error     { _ = "STUB: not implemented"; return nil }

// NewCache creates a new cache of fonts from the provided collection of
// font Faces.
// The first font Face in the collection is set to be the default one.
func NewCache(coll Collection) *Cache { _ = "STUB: not implemented"; return nil }

// Add adds a whole collection of font Faces to the font cache.
// If the cache is empty, the first font Face in the collection is set
// to be the default one.
func (c *Cache) Add(coll Collection) { _ = "STUB: not implemented"; return }

// store all font descriptors with the same size.

// Lookup returns the font Face corresponding to the provided Font descriptor,
// with the provided font size set.
//
// If no matching font Face could be found, the one corresponding to
// the default typeface is selected and returned.
func (c *Cache) Lookup(fnt Font, size Length) Face { _ = "STUB: not implemented"; return *new(Face) }

// Has returns whether the cache contains the exact font descriptor.
func (c *Cache) Has(fnt Font) bool { _ = "STUB: not implemented"; return false }

func (c *Cache) lookup(key Font) *opentype.Font { _ = "STUB: not implemented"; return nil }

func weightName(w font.Weight) string { _ = "STUB: not implemented"; return "" }

func styleName(sty font.Style) string { _ = "STUB: not implemented"; return "" }
