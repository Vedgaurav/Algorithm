package main

import (
	"fmt"
)

func bellmanFord(source int, graph [8][8]int) {

	size := len(graph)
	var distance = make([]int, size, size)

	for i := 0; i < size; i++ {
		if i == 0 {
			distance[i] = 0
			continue
		}
		// distance[i] = math.MaxInt64  //9223372036854775807
		// fmt.Println("int64 value : ", distance[i])
		distance[i] = 999
	}

	for u := 0; u < size; u++ {
		for v := 0; v < size; v++ {

			if graph[u][v] != 0 {

				if distance[v] > distance[u]+graph[u][v] {
					distance[v] = distance[u] + graph[u][v]
				}
			}
		}
	}
	fmt.Println(distance)
	for u := 0; u < size; u++ {
		for v := 0; v < size; v++ {

			if graph[u][v] != 0 {

				if distance[v] > distance[u]+graph[u][v] {
					fmt.Println("there is negative cycle in the graph")
				}
			}
		}
	}

}
