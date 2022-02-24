package main

import (
	"curso-go/math" // O caminho de onde a classe está a partir de src
	"fmt"
)

//Go é fortemente tipado como o Java
var a string

func main() {
	a = "Montival"
	//a = "Montival" -> Forma de declarar e ja atribuir valor
	fmt.Println(a)
	resultado := math.Soma(2, 4)
	fmt.Printf("%T \n", resultado) // %T imprime o tipo da variável
	fmt.Printf("%v \n", resultado) // %v imprime o valor da variavel

	fmt.Printf("O numero digitado é %v", math.Soma_impares(3))
}
