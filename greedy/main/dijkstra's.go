package main

import (
	"fmt"
)

func dijkstraSpt(graph [9][9]int) {

	size := len(graph)
	source:= 0
	fmt.Println("size: ", size)

	//1st Step make a nil set for storing the vertices
	var sptSet = make(map[int]struct{})
	var sptSetRemaining = make(map[int]struct{})

	for i:=0; i< size; i++{
		sptSetRemaining[i] = struct{}{}
	}

	//2D Slices Method I =>allocate a single array and point the individual slices into it.
	//As length of the slices are fixed is not going to grow or shrink
	// distanceValue := make([][]int, size)
	// rows := make([]int, size*size)
	// for i := range distanceValue {
	// 	distanceValue[i], rows = rows[:size], rows[size:]
	// }
	var distanceValue = make([]int, size, size)

		
	//2nd Step make a distance value source = 0 rest all values infinity
	for i := 0; i < size; i++ {
		if i == source  {
				distanceValue[i] = 0
				continue
			}
			distanceValue[i] = 999
		}

	//3rd Step get min value from distanceValue if not  present in sptSet store and
	//update the adjacent vertices of u. v =  distanceValue of source + weight(u,v) < distanceValue(v)

	for len(sptSetRemaining) != 0 {
		fmt.Println("length of sptSet: ", len(sptSet))

		//3.a Pick a vertex u which is not in sptSet and has minimum distance value
		minFromDistanceValue, vertexu := minValueFrom2D(distanceValue, sptSetRemaining)
		fmt.Println("Min Value ", minFromDistanceValue)
		fmt.Println("Vertex ", vertexu)

		_, ok := sptSet[vertexu]
		if !ok {
			sptSet[vertexu] = struct{}{}
			delete(sptSetRemaining, vertexu)
		}

		distanceValue = updateAdjVerticesOfU(vertexu, distanceValue, graph)
		for i := range distanceValue{
			fmt.Println(distanceValue[i])
		}
	}

	fmt.Println("=========================================================================================")
	for i := range distanceValue{
		fmt.Printf("%d ====> %d\n",i,distanceValue[i])
	}
	
	fmt.Println("=========================================================================================")

}

func updateAdjVerticesOfU(u int, distanceValue []int, graph [9][9]int) (distValue []int) {

	for adjVertexv := 0; adjVertexv < len(graph); adjVertexv++ {

		if graph[u][adjVertexv] != 0 {

			if distanceValue[u] + graph[u][adjVertexv] < distanceValue[adjVertexv] {

				distanceValue[adjVertexv] = distanceValue[u] + graph[u][adjVertexv]
				
			}

		}
	}
	return distanceValue
}

func minValueFrom2D(distanceValue []int, sptSetRemaining map[int]struct{}) (minFromDistanceValue, vertex int) {

	var vertexu int
	var min int
	min = 999
	
	for j := range sptSetRemaining{
			
			if min < distanceValue[j] {
				continue
			} else {
				min = distanceValue[j]
				vertexu = j
			}
		}
	return min, vertexu
}

//2D Slice Method II =>If the slices might grow or shrink, they should be allocated
//independently to avoid overwriting the next line
// distanceValue := make([][]int, size)
//

// for i:= range distanceValue {
// 	distanceValue[i] = make([]int, size)
// }
// for  j:= range distanceValue {
// 	fmt.Println(j)
// 	fmt.Println(&distanceValue[0][0])
// 	fmt.Println(&distanceValue[0][1])
// 	fmt.Println(&distanceValue[0][2])
// 	fmt.Println(&distanceValue[0][3])
// }
// fmt.Println("distance value: ",distanceValue)
// //fmt.Println("rows: ", rows)
// fmt.Println("distanceValueSize: ", len(distanceValue))
