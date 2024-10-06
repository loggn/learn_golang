package main

import (
	"fmt"
	"myproject/utils"
)

func main() {
	var arr = []int{7, 2, 9, 4}
	var arruil []int
	fmt.Println("数组元素的和是",utils.SumArray(arr))
	max, err := utils.FindMax(arruil)
	if err != nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println("切片中最大的值是", max)
	}

	var bank utils.Account
	var num float64
	bank.Balance = 1000.0
	num = 100
	bank.Deposit(num)
	num = 1000
	err1 := bank.Withdtaw(num)
	if err1 != nil {
		fmt.Println("错误:", err1)
	} 
}