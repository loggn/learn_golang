package main

import "fmt"

func main(){
	var (
		num1 float64
		num2 float64
		operator string
	)

	fmt.Scanf("%f %f %s", &num1, &num2, &operator)

	if operator == "+" {
		fmt.Printf("%v", num1+num2)	
	} else if operator == "-" {
		fmt.Printf("%v", num1-num2)	
	} else if operator == "*" {
		fmt.Printf("%v", num1*num2)	
	} else if operator == "/" {
		if num2 == 0 {
			fmt.Println("不能除以0")
		} else{
				fmt.Printf("%v", num1/num2)	
		}
	} else {
		fmt.Println("输入的操作符不合法")
	}

}