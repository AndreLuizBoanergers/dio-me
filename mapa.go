package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])
	elemento := make(map[string]string)
	elemento["H"] = "Hidrogenio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Litio"

	fmt.Println(elemento["Li"])
}
