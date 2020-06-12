package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter the size of array")
	var size int
	fmt.Scanf("%d",&size)

	var arr = make([]int, size, size)

	fmt.Println("Enter the elements of array")
	
	for i:=0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	
}