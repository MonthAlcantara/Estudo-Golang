package main

import (
	"errors"
	"fmt"
)

//O Go não é totalmente orientado a objetos. Possui esse Structs que se assemelha a um Objeto em Java
type Pessoa struct {
	//Variaveis iniciando com letra maiuscula para ser publico
	Nome  string
	Idade int
}

/* Essa seria a forma de fazer o método andou() ser um método da struct pessoa.
func (p Pessoa) andou() {
	fmt.Println(p.Nome, "andou")
}
*/
func andou(pessoa Pessoa) {
	fmt.Println(pessoa.Nome, "andou")
}

//O go permite ter dois retornos pra o mesmo método, nesse caso usei string e error
func andou_com_dois_retornos(pessoa Pessoa) (string, error) {
	if pessoa.Nome != "Montival" {
		//A ordem do retorno é a da assinatura do método (string, error). Eu posso escolher omitir um retorno e postar apenas um
		return "", errors.New("Nome tem que ser Montival")
	}
	// nil seria o null do java. Nesse caso só envio a string e não o erro
	return pessoa.Nome + "andou", nil
}

//Go n tem exception...o tratamento é feito desse jeito abaixo. O erro no Go é cidad~eo de primeira
func main() {
	// Em vez de new é passado algo como um Json a variavel
	var montival Pessoa
	montival = Pessoa{
		Nome:  "Montival",
		Idade: 31,
	}
	fmt.Println(montival)
	andou(montival)
	//Aqui eu estou jogando os dois retornos do método em duas variaveis res e err...como o res (string) n vem... eu posso passar um _ e ignorar ele
	_, err := andou_com_dois_retornos(Pessoa{Nome: "Monti"})
	if err != nil {
		fmt.Println("Deu erro")
		fmt.Println(err.Error())

	}
	andou_com_dois_retornos(montival)
}
