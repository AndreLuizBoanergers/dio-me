package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Tres")
		case 4:
			fmt.Println("Quatro")
		case 5:
			fmt.Println("Cinco")
		}
	}
}
