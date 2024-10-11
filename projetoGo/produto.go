package main

import "fmt"

type Produto struct {
	nome  string
	preco float64
}

func NovoProduto(nome string, preco float64) Produto {
	return Produto{nome: nome, preco: preco}
}

func (p Produto) Somar(outro Produto) Produto {
	novoNome := fmt.Sprintf("%s e %s", p.nome, outro.nome)
	novoPreco := p.preco + outro.preco
	return NovoProduto(novoNome, novoPreco)
}

func (p Produto) String() string {
	return fmt.Sprintf("%s: R$ %.2f", p.nome, p.preco)
}
