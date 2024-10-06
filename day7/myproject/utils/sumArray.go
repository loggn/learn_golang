package utils

import (
	"fmt"
)

func SumArray(arr []int) int {
	var sum int

	for _, value := range arr {
		sum += value
	}

	fmt.Println("数组元素的和是", sum)
	return sum
}