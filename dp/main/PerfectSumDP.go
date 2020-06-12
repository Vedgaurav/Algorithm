package main

import (
	"fmt"

	"encoding/json"
)


func perfectSumDP(arr []int, perfect int) {

	var subsetSumMap = make(map[string]int)
	count := 0
	for s := 0; s < len(arr); s++ {
		var tempArr []int
		sum := 0
		checkKey := false
		for i := s; i < len(arr); i++ {

			//sum = sum + arr[i]

			tempArr = append(tempArr, arr[i])
			subsetString, _ := json.Marshal(tempArr)
			fmt.Printf("marshalling subsetString %d", subsetString)
			sum, checkKey = subsetSumMap[string(subsetString)]
			fmt.Printf("sum, checkKey %d %t", sum, checkKey)
			if !checkKey {
				fmt.Println("came inside if statement")
				sum = sum + arr[i]
				count++
				fmt.Println("sum = == ", sum)
				subsetSumMap[string(subsetString)] = sum
			}

			fmt.Println("map: ", subsetSumMap)
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

					subsetString, _ := json.Marshal(tempArr2)
					fmt.Printf("marshalling subsetString %d", subsetString)
					value, checkKey2 := subsetSumMap[string(subsetString)]
					fmt.Printf("value, checkKey %d %t", value, checkKey)
					if !checkKey2 {
						fmt.Println("came inside if statement")
						result = result + arr[k]
						count++
						subsetSumMap[string(subsetString)] = result
					}else{
						result = value
					}
					fmt.Println("map: ", subsetSumMap)

					//result = result + arr[k]
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
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%count: ", count)
}
