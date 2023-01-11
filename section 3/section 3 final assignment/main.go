package main

import "fmt"

func main() {
	newArr := []int {0,1,2,3,4,5,6,7,8,9,10}

	fmt.Println("array we're checking is ", newArr)

	for _, num := range newArr {
		if num % 2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}