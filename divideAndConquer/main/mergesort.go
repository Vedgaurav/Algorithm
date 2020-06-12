package main

import(
	"fmt"
)

func mergesort(start, end int, arr []int) {

	if start < end {

		mid := (end + start) / 2

		fmt.Println("mid ", mid)

		mergesort(start, mid, arr)
		mergesort(mid+1, end, arr)
		merge(start,mid, end, arr)

	}
}

func merge(start,mid, end int, arr []int)(mergeArray []int) {

	leftArraySize := (mid - start) + 1
	rightArraySize := end - mid

	fmt.Println("leftArraySize  ",leftArraySize)
	fmt.Println("rightArraySize ", rightArraySize)

	var leftArray = make([]int, leftArraySize +1 , leftArraySize + 1)
	var rightArray = make([]int, rightArraySize + 1, rightArraySize + 1)

	

	for i := 0; i < leftArraySize; i++ {
		leftArray[i] = arr[start + i]
	}

	for j := 0; j < rightArraySize; j++ {
		rightArray[j] = arr[mid + 1 + j]
	}

	fmt.Println("leftArray ", leftArray)
	fmt.Println("rightArray ", rightArray)

	leftArray[leftArraySize] = 999
	rightArray[rightArraySize] = 999

	fmt.Println("leftArray ", leftArray)
	fmt.Println("rightArray ", rightArray)

	j := 0
	k := 0

	
	for i := start; i <= end ; i++ {

		if leftArray[j] < rightArray[k]{
			arr[i] = leftArray[j]
			j++
		}else{
			arr[i] = rightArray[k]
			k++
		}

	}
	fmt.Println("Merge array ", arr)
	return arr
}
