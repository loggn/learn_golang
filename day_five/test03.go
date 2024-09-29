package main

import "fmt"

func factorial(num int) int {

	if num < 0 {
		fmt.Println("请输入非负整数")
		return -1
	} else if num == 0 {
		return 0
	}

	j := 1
	for i := 1 ; i <= num ; i++ {
		j *= i
	}

	return j
}

func main() {
	var (
		num int
		fac int
	)
	fmt.Scanln(&num)

	fac = factorial(num)
	
	if fac == -1 {
		return
	}
	
	fmt.Println(num, "的阶乘是", fac)
}