package main

import "fmt"

func main()  {
	var num int

	fmt.Scanln(&num)

	if (num % 2) == 0 {
		fmt.Println(num, "是偶数")
		fmt.Printf("%d 转换为浮点数：%0.1f", num, float32(num))
	} else {
		fmt.Println(num, "是奇数")
		fmt.Printf("%d 转换为浮点数：%0.1f", num, float32(num))
	}
}