package main

func mergesort(start, end int, arr []int){

	
	
	if len(arr) > 1{

	mid := len(arr)/2

	mergesort(start, mid, arr)
	mergesort(mid+1, end, arr)
	merge(start, end, arr)

	}
}

func merge(start, end int, arr []int){

	
}