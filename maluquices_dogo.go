package main

import "fmt"

func main() {
	nome := "montival"
	//Quando coloco o & na frente da variável, ele me retorna o endereço na memória e não o valor da variável
	fmt.Println(&nome)
	fmt.Println(*&nome)

	//Esse ateristico é um ponteiro
	//Ou seja ele é um apontamento pra algum lugar na memória
	var nome2 *string
	// Aqui eu estou pegando o endereço de nome e copiando esse endereço para nome2, ja que ele é um apontamento devido o * na frente da variavel
	nome2 = &nome

	println(nome2)
	//Sendo assim se eu imprimir o *nome2 ele vai me retornar o valor, uma vez que nome2 é o apontamento pra memoria
	println(*nome2)

	/*
		Se a variavel for a apontamento pra memória e eu usar o * ele me mostra o valor
		Se a variavel for o valor e eu usar o * ele me mostra o endereço na memória
		& na frente no dome da variável eu ja estou criando a variável com o fim de salvar o endereço na memoria
		Sendo assim se eu fizer *&nome_da_variavel, ele me mostrará o valor dela
	*/
}
