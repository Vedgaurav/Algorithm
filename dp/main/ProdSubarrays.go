package main

import (
	"fmt"
)



func subArrays(arr []int) {

	size := len(arr)
	prod := 1

	fmt.Printf("Subarrays\n")
	for j := 0; j < size; j++ {

		for k := 0; k < size-j; k++ {

			tempArr := arr[k : k+j+1]
			fmt.Printf("%d ", tempArr)

			prod = prod * prodSubArrays(tempArr)
		}

	}
	fmt.Printf("\nProduct of SubArrays: %d", prod)
}

func prodSubArrays(arr []int) (prod int) {

	size := len(arr)
	prod = 1

	for i := 0; i < size; i++ {

		prod = prod * arr[i]
	}

	return prod
}

func efficientProdSubArrays(arr []int) {

	result := 1

	for i := 0; i < len(arr); i++ {
		product := 1
		for j := i; j < len(arr); j++ {
			product *= arr[j]
			fmt.Printf("Sub Arrays %d", product)
			result *= product
		}
	}
	fmt.Printf("Product of SubArrays: %d", result)
}
