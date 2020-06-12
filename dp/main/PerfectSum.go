package main

import (
	"fmt"
)



func perfectSum(arr []int, perfect int) {
	count := 0
	for s := 0; s < len(arr); s++ {
		var tempArr []int
		sum := 0
		for i := s; i < len(arr); i++ {

			sum = sum + arr[i]
			count++
			tempArr = append(tempArr, arr[i])

			fmt.Printf("**tempArrr%d", tempArr)
			fmt.Printf("sum %d", sum)

			if sum == perfect {
				fmt.Println("\n===================init==============================================================================\n")
				fmt.Printf("[%d] = %d", tempArr, sum)
				fmt.Println("\n====================init===============================================================================\n")
				break
			}

			for j := i; j < len(arr); j++ {
				result := sum
				var tempArr2 []int
				tempArr2 = tempArr
				fmt.Printf("**tempArrr2%d**", tempArr2)
				for k := j + 1; k < len(arr); k++ {

					fmt.Printf("\ntempArr initial %d", tempArr2)
					tempArr2 = append(tempArr2, arr[k])

					result = result + arr[k]
					count++
					fmt.Printf("result  %d", result)

					fmt.Printf("tempArr after %d\n", tempArr2)

					if result == perfect {
						fmt.Println("\n===============================================================================================\n")
						fmt.Printf("[%d] = %d", tempArr2, perfect)
						fmt.Println("\n================================================================================================\n")
					}
				}
			}
		}
	}

	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%count: ", count)
}
