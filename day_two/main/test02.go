package main

import "fmt"

func main() {
	
	var length, width, height, volume int

	fmt.Scanln(&length)
	fmt.Scanln(&width)
	fmt.Scanln(&height)
	
	volume = length * width * height
	
	fmt.Println("volume =", volume)
}


