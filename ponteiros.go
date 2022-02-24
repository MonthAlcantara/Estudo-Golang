package main

type Pessoa struct {
	Nome  string
	Idade int
}

//Esse cara muda apenas o nome naquele contexto..não altera a struct em si
func (p Pessoa) setNome(nome string) {
	p.Nome = nome
}

//Esse cara altera a struct pq o que é passado para o método é o ponteiro que aponta para a struct
// sendo assim, alterando aqui, o struct é alterado
func (p *Pessoa) setNomePonteiro(nome string) {
	p.Nome = nome
}
