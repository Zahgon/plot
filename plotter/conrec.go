// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

type point struct {
	X, Y float64
}

type line struct {
	p1, p2 point
}

func sect(h, p [5]float64, v1, v2 int) float64 { _ = "STUB: not implemented"; return 0 }

// conrecLine performs an operation with a line at a given height derived
// from data over the 2D box interval (i, j) to (i+1, j+1).
type conrecLine func(i, j int, l line, height float64)

// conrec is a Go translation of the C version of CONREC by Paul Bourke:
// http://paulbourke.net/papers/conrec/conrec.c
//
// conrec takes g, an m×n grid function, a sorted slice of contour heights
// and a conrecLine function.
//
// For full details of the algorithm, see the paper at
// http://paulbourke.net/papers/conrec/
func conrec(g GridXYZ, heights []float64, fn conrecLine) { _ = "STUB: not implemented"; return }

// We differ from conrec.c in the assignment of a single value
// in cases (castab in conrec.c). The value of castab[1][1][1] is
// 3, but we set cases[1][1][1] to 0.
//
// axiom: When we have a section of the grid where all the
// Z values are equal, and equal to a contour height we would
// expect to have no internal segments to draw.
//
// This is covered by case g) in Paul Bourke's description of
// the CONREC algorithm (a triangle with three vertices the lie
// on the contour level). He says, "... case g above has no really
// satisfactory solution and fortunately will occur rarely with
// real arithmetic." and then goes on to show the following image:
//
// http://paulbourke.net/papers/conrec/conrec3.gif
//
// which shows case g) in the set where no edge is drawn, agreeing
// with our axiom above.
//
// However, in the iteration over sh at conrec.c +44, a triangle
// with all vertices on the plane is given sh = {0,0,0,0,0} and
// then when the switch at conrec.c +93 happens, castab resolves
// that to case 3 for all values of m.
//
// This is fixed by replacing castab/cases[1][1][1] with 0.

/*
   Note: at this stage the relative heights of the corners and the
   centre are in the h array, and the corresponding coordinates are
   in the xh and yh arrays. The centre of the box is indexed by 0
   and the 4 corners by 1 to 4 as shown below.
   Each triangle is then indexed by the parameter m, and the 3
   vertices of each triangle are indexed by parameters m1,m2,and m3.
   It is assumed that the centre of the box is always vertex 2
   though this isimportant only when all 3 vertices lie exactly on
   the same contour level, in which case only the side of the box
   is drawn.

      vertex 4 +-------------------+ vertex 3
               | \               / |
               |   \    m-3    /   |
               |     \       /     |
               |       \   /       |
               |  m=2    X   m=2   |       the centre is vertex 0
               |       /   \       |
               |     /       \     |
               |   /    m=1    \   |
               | /               \ |
      vertex 1 +-------------------+ vertex 2
*/

// Scan each triangle in the box.

// Line between vertices 1 and 2

// Line between vertices 2 and 3

// Line between vertices 3 and 1

// Line between vertex 1 and side 2-3

// Line between vertex 2 and side 3-1

// Line between vertex 3 and side 1-2

// Line between sides 1-2 and 2-3

// Line between sides 2-3 and 3-1

// Line between sides 3-1 and 1-2
