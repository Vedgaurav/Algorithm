package main

import(
	"fmt"
)

func main() {

	fmt.Println("Enter the graph")
	

	// graphDijsktra := [9][9]int{ {0,4,0,0,0,0,0,8,0},
	// 							{4,0,8,0,0,0,0,11,0},
	// 							{0,8,0,7,0,4,0,0,2},
	// 							{0,0,7,0,9,14,0,0,0},
	// 							{0,0,0,9,0,10,0,0,0},
	// 							{0,0,4,14,10,0,2,0,0},
	// 							{0,0,0,0,0,2,0,1,6},
	// 							{8,11,0,0,0,0,1,0,7},
	// 							{0,0,2,0,0,0,6,7,0}}					
	


	//dijkstraSpt(graphDijsktra)

	// graphBellmanFord := [5][5]int {{0,-1,4,0,0},
	// 							   {0,0,3,2,2},
	// 							   {0,0,0,0,0},
	// 							   {0,1,5,0,0},
	// 							   {0,0,0,-3,0}}

	graphBellmanFordNegCycle := [8][8]int{{0,4,4,0,0,0,0,0},
										  {0,0,0,0,0,0,0,0},
										  {0,0,0,0,4,-2,0,0},
										  {3,0,2,0,0,0,0,0},
										  {0,0,0,1,0,0,-2,0},
										  {0,3,0,0,-3,0,0,0},
										  {0,0,0,0,0,2,0,2},
										  {0,0,0,0,-2,0,0,0}}

	//bellmanFord(0,graphBellmanFord)
	bellmanFord(0,graphBellmanFordNegCycle)
	
}