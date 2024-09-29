package main

import "fmt"

func numberClass(num int)  {
	switch {
	case num < 0 :
		fmt.Println(num, "为负数")
	case num == 0 :
		fmt.Println(num, "为零")
	case num >= 0 && num <= 10 :
		fmt.Println(num, "属于1 到10 的范围")
	case num > 10 :
		fmt.Println(num, "大于10")		
	}
}

func main(){
	var num int
	fmt.Scanln(&num)
	numberClass(num)
}