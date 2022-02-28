package main

import (
	"fmt"
	"net/http" //pacote para conexão com a internet
	"os"       // usado para comunicação com o SO
)

func main() {
	//O GO só tem for para repetição... se eu escrever for e não passar mais nada, ele vai agir como um While(true) no java
	for {

		fmt.Println("1- Iniciar monitoramento")
		fmt.Println("2- Exibir logs")
		fmt.Println("0- Sair do Programa")

		comando := leComando()
		println("Voce escolheu a opcao", comando)

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			exibeLogs()
		case 0:
			encerraPrograma()
			os.Exit(0) // Passar zero para S.O. indica que o programa finalizou sem falhas
		default:
			fmt.Println("Escolha um numero valido")
			os.Exit(-1) // Passar menos 1 para S.O. indica que o programa finalizou com falhas
		}
	}
}
func leComando() int {
	var comando int
	//Scan para capturar interação com o usuario (System.in do Java)
	//Como eu declarei "comando" com int...o scan infere que o valor recebido é um numero
	//& para indicar que o comando representa o endereço na memoria
	fmt.Scan(&comando)
	fmt.Scanf("%d", comando)
	return comando

}
func iniciaMonitoramento() {
	fmt.Println("Monitoramento iniciado")
	site := "https://www.alura.com.br"
	// No Go um método pode retornar mais de um valor
	// o Get do http por exemplo retorna um resp *http.Response, err error
	resp, _ := http.Get(site)
	fmt.Println(resp)
}

func exibeLogs() {
	fmt.Println("Exibindo logs")
}

func encerraPrograma() {
	fmt.Println("Programa finalizado")
}

// Dessa forma eu consigo devolver mais de um retorno e quem chamar pode escolher pegar algum ou todos os retornos
func devolveNomeEIdade() (string, int, int) {
	nome := "Junior"
	idade := 31
	return nome, idade, 11
}
