package main

import (
	"fmt"
)

func main() {
	V := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}
	E := []edge{
		{"a", "b", 10},
		{"a", "c", 5},
		{"b", "d", 30},
		{"b", "g", 15},
		{"c", "f", 10},
		{"d", "c", 5},
		{"d", "f", 3},
		{"d", "e", 10},
		{"e", "f", 5},
		{"e", "h", 10},
		{"e", "i", 20},
		{"e", "j", 30},
		{"f", "h", 20},
		{"f", "a", 30},
		{"g", "d", 5},
		{"g", "e", 5},
		{"g", "k", 15},
		{"h", "i", 5},
		{"h", "n", 20},
		{"i", "l", 5},
		{"i", "m", 5},
		{"i", "n", 10},
		{"j", "i", 40},
		{"j", "l", 5},
		{"j", "g", 10},
		{"k", "j", 17},
		{"k", "l", 20},
		{"l", "m", 10},
		{"m", "n", 15},
		{"m", "l", 15},
		{"n", "m", 12},
	}

	G := newGraph(V, E)
	G.dijkstra("a")
	for _,v := range V {
		fmt.Println(v, ":", G[v].len, G[v].path)
	}
}

type graph map[string]*vertex

type vertex struct {
	adj []tailLen
	len int
	path []string
}

type edge struct {
	head, tail string
	len        int
}

type tailLen struct {
	tail string
	len  int
}

func newGraph(V []string, E []edge) graph {
	g := make(graph)
	for _, v := range V {
		x := vertex{len: -1}
		for _, e := range E {
			if e.head == v {
				x.adj = append(x.adj, tailLen{tail: e.tail, len: e.len})
			}
		}
		g[v] = &x
	}
	return g
}



func (G graph) dijkstra(s string) {
	X := make(StrSet)
	X.Push(s)
	G[s].len = 0

	for {
		Len := maxInt
		var v,w string
		for x := range X {
			xLen := G[x].len
			for _, z := range G[x].adj {
				if !X.Contains(z.tail) && xLen + z.len < Len {
					v = x
					w = z.tail
					Len = xLen + z.len
				}
			}
		}
		if Len == maxInt {
			return
		}
		G[w].len = Len
		G[w].path = append(G[v].path, w)
		X.Push(w)
	}
}

type StrSet map[string]struct{}

func (x StrSet) Contains(i string) bool {
	_, r := x[i]
	return r
}

func (x StrSet) Push(i string) {
	x[i] = struct{}{}
}

func (x StrSet) Pop(i string) {
	delete(x, i)
}

const maxInt = int((^uint(0)) >> 1)
