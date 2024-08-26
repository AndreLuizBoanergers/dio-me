package main

import "fmt"

func main() {
	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	fatia := arr[2:5]
	fmt.Println(fatia)
	//
	fatia1 := []int{1, 2, 3, 4}
	fatia2 := append(fatia1, 5, 6, 7)
	fmt.Println(fatia1, fatia2)
}
