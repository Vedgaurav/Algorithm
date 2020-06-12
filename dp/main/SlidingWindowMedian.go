package main

import (
	"fmt"
	"sort"
)

func main2() {

	fmt.Println("Enter the size array")
	var SIZE int
	fmt.Scanf("%d", &SIZE)

	arr := make([]int, SIZE, SIZE)

	fmt.Println("Enter the elements of array")
	for i := 0; i < SIZE; i++ {

		fmt.Scanf("%d", &arr[i])
	}
	fmt.Printf("array %d", arr)
	fmt.Println("\nEnter the size of the window")
	var window int
	fmt.Scanf("%d", &window)

	for i := 0; i <= len(arr)-window; i++ {

		fmt.Printf("array %d", arr)
		tempArr := make([]int, window, window)

		copy(tempArr[:], arr[i:window+i])
		fmt.Printf("tempArray %d", tempArr)

		sort.Ints(tempArr)
		fmt.Printf("sortedArray %d", tempArr)

		if window%2 == 0 {

			temp1 := len(tempArr) / 2

			median := float32(tempArr[temp1]+tempArr[temp1-1]) / 2
			fmt.Printf("median %f\n", median)

		} else {

			median := (len(tempArr) / 2)
			fmt.Printf("median %d\n", tempArr[median])
		}
	}
}
