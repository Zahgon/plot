// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter

// johnson implements Johnson's "Finding all the elementary
// circuits of a directed graph" algorithm. SIAM J. Comput. 4(1):1975.
//
// Comments in the johnson methods are kept in sync with the comments
// and labels from the paper.
type johnson struct {
	adjacent graph // SCC adjacency list.
	b        []set // Johnson's "B-list".
	blocked  []bool
	s        int

	stack []int

	result [][]int
}

// cyclesIn returns the set of elementary cycles in the graph g.
func cyclesIn(g graph) [][]int { _ = "STUB: not implemented"; return nil }

// len(j.adjacent) will be the order of g until Tarjan's analysis
// finds no SCC, at which point t.sccSubGraph returns nil and the
// loop breaks.

// We use the previous SCC adjacency to reduce the work needed.

// A_k = adjacency structure of strong component K with least
//       vertex in subgraph of G induced by {s, s+1, ... ,n}.
// Only allow SCCs with >= 2 vertices.

// s = least vertex in V_k

//L3:

// circuit is the CIRCUIT sub-procedure in the paper.
func (j *johnson) circuit(v int) bool { _ = "STUB: not implemented"; return false }

//L1:

// Output circuit composed of stack followed by s.

//L2:

// unblock is the UNBLOCK sub-procedure in the paper.
func (j *johnson) unblock(u int) { _ = "STUB: not implemented"; return }

// tarjan implements Tarjan's strongly connected component finding
// algorithm. The implementation is from the pseudocode at
//
// http://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm?oldid=642744644
type tarjan struct {
	g graph

	index      int
	indexTable []int
	lowLink    []int
	onStack    []bool

	stack []int

	sccs [][]int
}

// newTarjan returns a tarjan with the sccs field filled with the
// strongly connected components of the directed graph g.
func newTarjan(g graph) *tarjan { _ = "STUB: not implemented"; return nil }

// strongconnect is the strongconnect function described in the
// wikipedia article.
func (t *tarjan) strongconnect(v int) {
	_ = "STUB: not implemented"
	// Set the depth index for v to the smallest unused index.
	return
}

// Consider successors of v.

// Successor w has not yet been visited; recur on it.

// Successor w is in stack s and hence in the current SCC.

// If v is a root node, pop the stack and generate an SCC.

// Start a new strongly connected component.

// Add w to current strongly connected component.

// Output the current strongly connected component.

// sccSubGraph returns the graph of the tarjan's strongly connected
// components with each SCC containing at least min vertices.
// sccSubGraph returns nil if there is no SCC with at least min
// members.
func (t *tarjan) sccSubGraph(min int) graph { _ = "STUB: not implemented"; return *new(graph) }

// set is an integer set.
type set map[int]struct{}

// graph is an edge list representation of a graph.
type graph []set

// remove deletes edges that make up the given paths from the graph.
func (g graph) remove(paths [][]int) { _ = "STUB: not implemented"; return }

// subgraph returns a subgraph of g induced by {s, s+1, ... , n}. The
// subgraph is destructively generated in g.
func (g graph) subgraph(s int) graph { _ = "STUB: not implemented"; return *new(graph) }

// clone returns a deep copy of the graph g.
func (g graph) clone() graph { _ = "STUB: not implemented"; return *new(graph) }
