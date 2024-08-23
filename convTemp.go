package main

import "fmt"

const kelvin = 373.2

func main() {
	cel := (kelvin - 273)
	fmt.Printf("O ponto ebulicao da gaua em celsios e: %g", cel)
}
