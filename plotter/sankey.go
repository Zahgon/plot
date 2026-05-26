// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/text"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

// A Sankey diagram presents stock and flow data as rectangles representing
// the amount of each stock and lines between the stocks representing the
// amount of each flow.
type Sankey struct {
	// Color specifies the default fill
	// colors for the stocks and flows. If Color is not nil,
	// each stock and flow is rendered filled with Color,
	// otherwise no fill is performed. Colors can be
	// modified for individual stocks and flows.
	Color color.Color

	// StockBarWidth is the widths of the bars representing
	// the stocks. The default value is 15% larger than the
	// height of the stock label text.
	StockBarWidth vg.Length

	// LineStyle specifies the default border
	// line style for the stocks and flows. Styles can be
	// modified for individual stocks and flows.
	LineStyle draw.LineStyle

	// TextStyle specifies the default stock label
	// text style. Styles can be modified for
	// individual stocks.
	TextStyle text.Style

	flows []Flow

	// FlowStyle is a function that specifies the
	// background color and border line style of the
	// flow based on its group name. The default
	// function uses the default Color and LineStyle
	// specified above for all groups.
	FlowStyle func(group string) (color.Color, draw.LineStyle)

	// StockStyle is a function that specifies, for a stock
	// identified by its label and category, the label text
	// to be printed on the plot (lbl), the style of the text (ts),
	// the horizontal and vertical offsets for printing the text (xOff and yOff),
	// the color of the fill for the bar representing the stock (c),
	// and the style of the outline of the bar representing the stock (ls).
	// The default function uses the default TextStyle, color and LineStyle
	// specified above for all stocks; zero horizontal and vertical offsets;
	// and the stock label as the text to be printed on the plot.
	StockStyle func(label string, category int) (lbl string, ts text.Style, xOff, yOff vg.Length, c color.Color, ls draw.LineStyle)

	// stocks arranges the stocks by category.
	// The first key is the category and the seond
	// key is the label.
	stocks map[int]map[string]*stock
}

// StockRange returns the minimum and maximum value on the value axis
// for the stock with the specified label and category.
func (s *Sankey) StockRange(label string, category int) (min, max float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// stock represents the amount of a stock and its plotting order.
type stock struct {
	// receptorValue and sourceValue are the totals of the values
	// of flows coming into and going out of this stock, respectively.
	receptorValue, sourceValue float64

	// label is the label of this stock, and category represents
	// its placement on the category axis. Together they make up a
	// unique identifier.
	label    string
	category int

	// order is the plotting order of this stock compared
	// to other stocks in the same category.
	order int

	// min represents the beginning of the plotting location
	// on the value axis.
	min float64

	// max is min plus the larger of receptorValue and sourceValue.
	max float64
}

// A Flow represents the amount of an entity flowing between two stocks.
type Flow struct {
	// SourceLabel and ReceptorLabel are the labels
	// of the stocks that originate and receive the flow,
	// respectively.
	SourceLabel, ReceptorLabel string

	// SourceCategory and ReceptorCategory define
	// the locations on the category axis of the stocks that
	// originate and receive the flow, respectively. The
	// SourceCategory must be a lower number than
	// the ReceptorCategory.
	SourceCategory, ReceptorCategory int

	// Value represents the magnitute of the flow.
	// It must be greater than or equal to zero.
	Value float64

	// Group specifies the group that a flow belongs
	// to. It is used in assigning styles to groups
	// and creating legends.
	Group string
}

// NewSankey creates a new Sankey diagram with the specified
// flows and stocks.
func NewSankey(flows ...Flow) (*Sankey, error) { _ = "STUB: not implemented"; return nil, nil }

// Here we make sure the stock categories are in the proper order.

// Here we initialize the stock holders.

// Here we figure out the plotting order of the stocks.

// Here we add the current value to the total value of the stocks

// Plot implements the plot.Plotter interface.
func (s *Sankey) Plot(c draw.Canvas, plt *plot.Plot) { _ = "STUB: not implemented"; return }

// sourceFlowPlaceholder and receptorFlowPlaceholder track
// the current plotting location during
// the plotting process.

// Here we draw the flows.

// Here we fill the flow polygons.

// Here we draw the flow edges.

// Here we draw the stocks.

// Here we fill the stock bars.

// poly)

// Here we draw the bottom edge.

// Here we draw the top edge plus vertical edges where there are
// no flows connected.

// stockList returns a sorted list of the stocks in the diagram.
func (s *Sankey) stockList() []*stock { _ = "STUB: not implemented"; return nil }

// stockSorter is a wrapper for a list of *stocks that implements
// sort.Interface.
type stockSorter []*stock

func (s stockSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s stockSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s stockSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// setStockRange sets the minimum and maximum values of the stock plotting locations.
func (s *Sankey) setStockRange(stocks *[]*stock) { _ = "STUB: not implemented"; return }

// bezier creates a bezier curve between the begin and end points.
func (s *Sankey) bezier(begin, end vg.Point) []vg.Point {
	_ = "STUB: not implemented"
	// directionOffsetFrac is the fraction of the distance between begin.X and
	// end.X for the bezier control points.
	return nil
}

// nPoints is the number of points for bezier interpolation.

// DataRange implements the plot.DataRanger interface.
func (s *Sankey) DataRange() (xmin, xmax, ymin, ymax float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

// GlyphBoxes implements the GlyphBoxer interface.
func (s *Sankey) GlyphBoxes(plt *plot.Plot) []plot.GlyphBox { _ = "STUB: not implemented"; return nil }

// Thumbnailers creates a group of objects that can be used to
// add legend entries for the different flow groups in this
// diagram, as well as the flow group labels that correspond to them.
func (s *Sankey) Thumbnailers() (legendLabels []string, thumbnailers []plot.Thumbnailer) {
	_ = "STUB: not implemented"
	return nil, nil
}

// sankeyFlowThumbnailer implements the Thumbnailer interface
// for Sankey flow groups.
type sankeyFlowThumbnailer struct {
	draw.LineStyle
	color.Color
}

// Thumbnail fulfills the plot.Thumbnailer interface.
func (t sankeyFlowThumbnailer) Thumbnail(c *draw.Canvas) {
	_ = "STUB: not implemented"
	// Here we draw the fill.
	return
}

// Here we draw the upper border.

// Here we draw the lower border.
