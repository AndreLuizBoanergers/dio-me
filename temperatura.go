package main

import "fmt"

//ponto de ebulicao da gua em Fahrenhe8it
const ebulicaoF float64 = 212.0

func main() {
	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9 //Conversao de F para C
	fmt.Println("A temperatura de ebulicao em  garus F = ", tempF)
	fmt.Println("A temperatura de ebulicao em graus C = ", tempC)
}
