package main

import "fmt"

func main() {
	// дано: граф G с вершинами V от 0 до 10
	V := []int{0,1,2,3,4,5,6,7,8,9,10}

	// дано: список рёбер E графа G:
	E := [][2]int{
		{0, 1},
		{0, 3},
		{0, 4},
		{1, 2},
		{1, 3},
		{2, 3},
		{2, 9},
		{3, 4},
		{3, 5},
		{4, 6},
		{5, 7},
		{6, 10},
		{7, 8},
		{7, 4},
		{8, 10},
		{9, 5},
		{9, 7},
		{9, 8},
		{10, 4},
	}

	// найти кратчайший путь между двумя произвольными вершинами графа G
	for _,v := range V{
		for i,x := range augmentedBreathSearch(V,E, v){
			fmt.Printf("%d-%d: %v\n",v, i, x)
		}
		fmt.Println()
	}
}

// augmentedBreathSearch - список истинных расстояний кратчайшего пути до вершины s
func augmentedBreathSearch(V []int, E [][2]int, s int) [][]int {

	type vertex struct {
		adjacency []int
		passed    bool
		path      []int
	}

	// пометить s как разведанную вершину, все остальные как неразведанные
	// дистанция от s до v: О если s == v, иначе +бесконечность
	vs := make(map[int]*vertex)
	for _,v := range V{
		var x vertex
		for _, e := range E {
			if e[0] == v {
				x.adjacency = append(x.adjacency, e[1])
			}
		}
		if v == s {
			x.passed = true
		}
		vs[v] = &x
	}

	// Q - очередь, инициализированная вершиной s
	Q := []int{s}
	for len(Q) > 0 {
		// удалить вершину из начала Q, назвать ее v
		v := Q[0]
		Q = Q[1:]
		// для каждого ребра (v, w) в списке смежности вершины v
		for _,w := range vs[v].adjacency {
			wi := vs[w]
			if wi.passed {
				continue
			}
			wi.passed = true
			wi.path = append(append(wi.path, vs[v].path...), w)
			Q = append(Q, w)
		}
	}

	xs := make([][]int, len(V))
	for i,v := range V {
		xs[i] = vs[v].path
	}
	return xs
}