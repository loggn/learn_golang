package main

import "fmt"

// 定义一个结构体
type Adder struct{}

func (a Adder) AddInt(x, y int) int {
    return x + y
}

func (a Adder) AddFloat(x, y float64) float64 {
    return x + y
}

func main() {
    adder := Adder{}
    
    fmt.Println("AddInt:", adder.AddInt(3, 4))           // 输出：AddInt: 7
    fmt.Println("AddFloat:", adder.AddFloat(3.5, 4.5))   // 输出：AddFloat: 8
}
