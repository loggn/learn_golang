package main

import (
	"fmt"
	"myproject/utils"
)

func main() {
	arr := utils.SpliString("Hello Go Language")
	var str []string =nil
	for _, value := range arr{
		str = append(str, "'", value, "'")
	} 
	fmt.Println(str)


	num, err := utils.Divide(10, 0)
	if err != nil {
		fmt.Println("错误：", err)
	} else {
		fmt.Println(num)
	}

	m := utils.WordCount("go is fun and go is powerful")
	fmt.Println(m)
}