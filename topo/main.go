package main

import (
	"fmt"
	"sort"
)

func main() {
	V := []string{
		"пинать хуи",
		"припарковать автомобиль",
		"вымыть руки",
		"купить колбасу",
		"прийти домой",
		"посрать",
		"накормить кота",
		"приготовить ужин",
		"поесть",
		"играть в компьютер",
		"спать",
	}
	E := [][2]string {
		{
			"припарковать автомобиль",
			"купить колбасу",
		},
		{
			"припарковать автомобиль",
			"прийти домой",
		},
		{
			"купить колбасу",
			"прийти домой",
		},
		{
			"прийти домой",
			"посрать",
		},
		{
			"прийти домой",
			"вымыть руки",
		},
		{
			"посрать",
			"вымыть руки",
		},
		{
			"вымыть руки",
			"накормить кота",
		},
		{
			"вымыть руки",
			"приготовить ужин",
		},
		{
			"накормить кота",
			"приготовить ужин",
		},

		{
			"приготовить ужин",
			"поесть",
		},

		{
			"поесть",
			"играть в компьютер",
		},
		{
			"поесть",
			"пинать хуи",
		},
		{
			"пинать хуи",
			"спать",
		},
		{
			"играть в компьютер",
			"спать",
		},
	}

	topologicalSort(V,E)

	for _,v := range V {
		fmt.Printf("%s\n", v)
	}

}

func topologicalSort(V []string, E [][2]string){
	level := len(V)
	graph := newGraph(V,E)
	for v, x := range graph{
		if !x.passed {
			dfsTopological(graph, v, &level)
		}
	}

	sort.Slice(V, func(i, j int) bool {
		return graph[V[i]].index < graph[V[j]].index
	})

	for i := range V {
		V[i] = fmt.Sprintf( "[%d] %s", graph[V[i]].index, V[i])
	}
}

func dfsTopological(g graph, s string, level *int){
	g[s].passed = true
	for _,v := range g[s].adjacency{
		if !g[v].passed {
			dfsTopological(g, v, level)
		}
	}
	g[s].index = *level
	*level -= 1
}

func newGraph(V []string, E [][2]string) graph {
	vs := make(graph)
	for _,v := range V{
		var x vertex
		for _, e := range E {
			if e[0] == v {
				x.adjacency = append(x.adjacency, e[1])
			}
		}
		vs[v] = &x
	}
	return vs
}

type graph = map[string]*vertex

type vertex struct {
	adjacency []string
	passed    bool
	index int
}