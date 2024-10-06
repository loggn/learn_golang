package utils


import (
	"errors"
)

func FindMax(slice []int) (int, error) {
	if slice == nil {
		return 0, errors.New("输入切片为空")
	}
	var max int

	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max, nil
}