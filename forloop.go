package main

import (
	"fmt"
)

func main() {
	num3 := 44
	fmt.Println(num3)
	var count int = 10
	hello(count)
}
func hello(count int) {
	for i := 0; i <= count; i++ {
		fmt.Println(i)

	}
}
