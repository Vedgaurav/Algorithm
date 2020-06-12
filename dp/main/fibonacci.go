package main

import (
	"fmt"
)

func fibonacciLoopSeries(n int) (fibonacci int) {

	if n <=1 {
		return n
	}
	a := 0
	b := 1
	var fibo int
    for i := 0; i <= n-2; i++ {
		fibo = a +  b
		a = b
		b = fibo

	}
	return fibo
}

//Fibonacci ...
//finds the fibonnaci sequence numbers of len n
func fibonacciLoopSeriesDP(n int) (arr1 []int) {

	var arr = make([]int, n, n)
	//fmt.Println(len(arr))

	arr[0] = 0
	arr[1] = 1
	
    for i := 0; i < n-2; i++ {
		arr[i+2] = arr[i] + arr[i+1]
	}
	return arr
}

var count int

func fibonacciRecursionSeriesDP(n int) (a []int) {

	if n < 2 {
		fmt.Println("Value should be greater")
		return
	}

	
	count++
	fmt.Printf("fib(%d) => %d\n", n, count)

	if n == 2 {
		var arr []int
		arr = append(arr, 0)
		arr = append(arr, 1)

		//count++
		//fmt.Printf("fib(%d) => %d\n", n, count)

		return arr
	}

	arr2 := fibonacciRecursionSeriesDP(n - 1)

	temp := arr2[n-2] + arr2[n-3]
	arr2 = append(arr2, temp)

	

	return arr2

}

func fibonacciRecursionSimple(n int) int {

	if n <= 1 {
		count++
		fmt.Printf("fib(%d) => %d\n", n, count)
		return n
	}
	count++
	fmt.Printf("fib(%d) => %d\n", n, count)
	return fibonacciRecursionSimple(n-2) + fibonacciRecursionSimple(n-1)
}

// func singletonMap()(singletMap map[int]int){

// 	var myMap map[int]int

// 	return myMap

// }

var myMap = make(map[int]int)

func fibonacciRecursionSimpleDP(n int) int {

	
	value, check := myMap[n]
	if check {
		fmt.Printf("fib(%d) => got the value from map\n", n)
		return value
	} 
	count++
	fmt.Printf("fib(%d) => %d\n", n, count)

	if n <= 1 {
		
		myMap[n] = n
		//fmt.Println("myMap: ", myMap)
		return n
	}

		temp := fibonacciRecursionSimpleDP(n-2) + fibonacciRecursionSimpleDP(n-1)
		myMap[n] = temp
		//fmt.Println("myMap: ", myMap)
		return temp
	
}
