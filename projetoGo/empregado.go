package main

import "fmt"

type Empregado struct {
	nome    string
	cargo   string
	salario interface{}
}

func (e Empregado) InfoEmpregado() string {
	return fmt.Sprintf("%s, Cargo: %s, Salário: R$ %v", e.nome, e.cargo, e.salario)
}
