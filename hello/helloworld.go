package main

import (
	"fmt"
)

func main() {

	//função fmt retorna o número de bytes e erros
	numerodebytes, erros := fmt.Println("Hello World!")
	fmt.Println(numerodebytes, erros)

	_, error := fmt.Println("Hello World!")
	fmt.Println(error)

	x := 57
	y := "10"
	z := true
	fmt.Println(x, y, z)

}
