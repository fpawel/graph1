package main

import "fmt"

func main() {
	V := []int{0,1,2,3,4,5,6,7,8,9,10}
	E := [][2]int{
		{0, 1},
		{1, 2},
		{2, 0},

		{3, 4},
		{4, 5},
		{5, 3},

		{6, 7},
		{7, 8},
		{8, 9},
		{9, 6},
		{9, 10},
	}

	for _,x := range uccSearch(V,E){
		fmt.Printf("%+v\n",x)
	}
}

func uccSearch(V []int, E [][2]int) [][]int {

	type vertex struct {
		adjacency []int
		passed    bool
		cc int
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
		x.cc = -1
		vs[v] = &x
	}

	numCC := 0
	for _,i := range V{
		if vs[i].passed {
			continue
		}

		numCC++ // новая компонента связности

		// Q - очередь, инициализированная вершиной s
		Q := []int{i}
		for len(Q) > 0 {
			// удалить вершину из начала Q, назвать ее v
			v := Q[0]
			Q = Q[1:]
			vs[v].cc = numCC
			// для каждого ребра (v, w) в списке смежности вершины v
			for _,w := range vs[v].adjacency {
				wi := vs[w]
				if wi.passed {
					continue
				}
				wi.passed = true
				Q = append(Q, w)
			}
		}
	}

	m := make(map[int] []int)
	for _,v := range V {
		m[vs[v].cc] = append(m[vs[v].cc], v)
	}

	var xs [][]int
	for _,vs := range m {
		xs = append(xs, vs)
	}
	return xs
}