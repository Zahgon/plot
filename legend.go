// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// A Legend gives a description of the meaning of different
// data elements of the plot.  Each legend entry has a name
// and a thumbnail, where the thumbnail shows a small
// sample of the display style of the corresponding data.
type Legend struct {
	// TextStyle is the style given to the legend
	// entry texts.
	TextStyle text.Style

	// Padding is the amount of padding to add
	// between each entry in the legend.  If Padding
	// is zero then entries are spaced based on the
	// font size.
	Padding vg.Length

	// Top and Left specify the location of the legend.
	// If Top is true the legend is located along the top
	// edge of the plot, otherwise it is located along
	// the bottom edge.  If Left is true then the legend
	// is located along the left edge of the plot, and the
	// text is positioned after the icons, otherwise it is
	// located along the right edge and the text is
	// positioned before the icons.
	Top, Left bool

	// XOffs and YOffs are added to the legend's
	// final position.
	XOffs, YOffs vg.Length

	// YPosition specifies the vertical position of a legend entry.
	// Valid values are [-1,+1], with +1 being the top of the
	// entry vertical space, and -1 the bottom.
	YPosition float64

	// ThumbnailWidth is the width of legend thumbnails.
	ThumbnailWidth vg.Length

	// entries are all of the legendEntries described
	// by this legend.
	entries []legendEntry
}

// A legendEntry represents a single line of a legend, it
// has a name and an icon.
type legendEntry struct {
	// text is the text associated with this entry.
	text string

	// thumbs is a slice of all of the thumbnails styles
	thumbs []Thumbnailer
}

// Thumbnailer wraps the Thumbnail method, which
// draws the small image in a legend representing the
// style of data.
type Thumbnailer interface {
	// Thumbnail draws an thumbnail representing
	// a legend entry.  The thumbnail will usually show
	// a smaller representation of the style used
	// to plot the corresponding data.
	Thumbnail(c *draw.Canvas)
}

// NewLegend returns a legend with the default parameter settings.
func NewLegend() Legend { _ = "STUB: not implemented"; return *new(Legend) }

func newLegend(hdlr text.Handler) Legend { _ = "STUB: not implemented"; return *new(Legend) }

// Draw draws the legend to the given draw.Canvas.
func (l *Legend) Draw(c draw.Canvas) { _ = "STUB: not implemented"; return }

// Rectangle returns the extent of the Legend.
func (l *Legend) Rectangle(c draw.Canvas) vg.Rectangle {
	_ = "STUB: not implemented"
	return *new(vg.Rectangle)
}

// entryHeight returns the height of the tallest legend
// entry text.
func (l *Legend) entryHeight() (height vg.Length) {
	_ = "STUB: not implemented"
	return *new(vg.Length)
}

// Add adds an entry to the legend with the given name.
// The entry's thumbnail is drawn as the composite of all of the
// thumbnails.
func (l *Legend) Add(name string, thumbs ...Thumbnailer) { _ = "STUB: not implemented"; return }
