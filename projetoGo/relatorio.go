package main

import "fmt"

type Relatorio struct {
	conteudo string
}

func (r Relatorio) Imprimir() {
	fmt.Printf("Relatório: %s\n", r.conteudo)
}
