package main

import "fmt"

type Empresa struct {
	nome       string
	empregados []*Empregado
}

func (e *Empresa) AdicionarEmpregado(empregado *Empregado) {
	e.empregados = append(e.empregados, empregado)
}

func (e Empresa) InfoEmpresa() string {
	info := fmt.Sprintf("Empresa: %s\nEmpregados:\n", e.nome)
	for _, empregado := range e.empregados {
		info += fmt.Sprintf("- %s\n", empregado.InfoEmpregado())
	}
	return info
}
