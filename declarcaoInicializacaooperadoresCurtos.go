package main

import "fmt"

//ponto de ebulicao da gua em Fahrenhe8it
const ebulicaoF = 212.0

func main() {
	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9 //Conversao de F para C
	/*
		fmt.Println("A temperatura de ebulicao em  garus F = ", tempF)
		fmt.Println("A temperatura de ebulicao em graus C = ", tempC)
	*/
	fmt.Printf("A temperatura de ebulicao em f e %g ,e a temperatura em C e %g", tempF, tempC)
}
