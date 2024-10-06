package utils

import "errors"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("不能除以零")
	}
	return a/b, nil
}