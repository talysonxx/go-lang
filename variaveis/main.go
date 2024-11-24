package main

import "fmt"

func main() {

	x := 10
	fmt.Printf("x: %v, %T\n", x, x)

	nome := "talyson"
	numerodebytes, _ := fmt.Printf("nome: %v\ntipo: %T\n", nome, nome)
	fmt.Println(numerodebytes)

}
