package main

import "fmt"

type Professor struct {
	nome    string
	escolas []*Escola
}

func (p *Professor) AdicionarEscola(escola *Escola) {
	if !p.temEscola(escola) {
		p.escolas = append(p.escolas, escola)
		escola.AdicionarProfessor(p)
	}
}

func (p *Professor) temEscola(escola *Escola) bool {
	for _, e := range p.escolas {
		if e == escola {
			return true
		}
	}
	return false
}

func (p *Professor) ExibirEscolas() {
	fmt.Printf("Professor %s leciona nas seguintes escolas:\n", p.nome)
	for _, escola := range p.escolas {
		fmt.Printf("- %s\n", escola.nome)
	}
}
