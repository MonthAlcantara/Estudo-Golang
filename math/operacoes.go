package math

// Eu posso ter mil arquivos, se todos estiverem no mesmo pacote, eu consigo importar
//pelo nome do pacote e usar seus métodos como se estivessem num arquivo só

func Soma(valor int, valor2 int) int {
	//Os métodos começam com letras maiusculas para poder ser exportadas
	// se declarar com letra minuscula o metodo n fica visível pra quem ta de fora
	//Maiuscula -> public, Minuscula -> private
	return soma(valor + valor2)
}

//método private
func soma(valor int) int {
	return valor
}
