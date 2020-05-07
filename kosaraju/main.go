package main

import (
	"fmt"
)

func main() {
	V := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	E := [][2]int{
		{1, 5},
		{2, 9},
		{3, 1},
		{3, 1},
		{4, 9},
		{4, 2},
		{5, 3},
		{6, 8},
		{6, 11},
		{7, 5},
		{7, 4},
		{8, 9},
		{8, 10},
		{8, 11},
		{9, 5},
		{9, 7},
		{10, 2},
		{10, 6},
		{11, 3},
	}

	G := newGraph(V, E)

	G.kosoraju(V)

	// print strong connectivity areas
	for _, v := range G.groupScc() {
		if len(v) > 1 {
			fmt.Printf("%d\n", v)
		}
	}
}

func newGraph(V []int, E [][2]int) graph {
	g := make(graph)
	for _, v := range V {
		var x vertex
		for _, e := range E {
			if e[0] == v {
				x.adjacency = append(x.adjacency, e[1])
			}
			if e[1] == v {
				x.adjacencyRev = append(x.adjacencyRev, e[0])
			}
		}
		g[v] = &x
	}
	return g
}

type graph map[int]*vertex

type vertex struct {
	adjacency    []int
	passed       bool
	adjacencyRev []int
	passedRev    bool
	scc          int
}

// kosoraju search for strong connectivity areas in a directed graph
func (g graph) kosoraju(V []int)  {

	// first mark all vertexes as unexplored
	// the first depth search traversing computes reversed graph topological order
	level := len(V)
	topoRev := make([]int, len(V))
	for v, x := range g {
		if !x.passedRev {
			g.dfsTopoRev(v, &level, topoRev)
		}
	}

	// the second depth search traversing finds strongly connected components
	scc := 0
	for i := range topoRev {
		v := topoRev[i]
		if !g[v].passed {
			scc++
			g.dfsScc(v, scc) // set values of scc
		}
	}
}

// dfsTopoRev computes topological order of the reversed graph using depth first search traversing
func (g graph) dfsTopoRev(s int, level *int, topoRev []int) {
	g[s].passedRev = true
	for _, v := range g[s].adjacencyRev {
		if !g[v].passedRev {
			g.dfsTopoRev(v, level, topoRev)
		}
	}
	topoRev[*level-1] = s
	*level -= 1
}

// dfsScc marks vertices by index of strong connectivity using depth first search traversing
func (g graph) dfsScc(s int, scc int) {
	g[s].passed = true
	for _, v := range g[s].adjacency {
		if !g[v].passed {
			g.dfsScc(v, scc)
		}
	}
	g[s].scc = scc
}

// groupScc groups areas of strong connectivity
func (g graph) groupScc() map[int][]int {
	m := make(map[int][]int)
	for v, x := range g {
		m[x.scc] = append(m[x.scc], v)
	}
	return m
}