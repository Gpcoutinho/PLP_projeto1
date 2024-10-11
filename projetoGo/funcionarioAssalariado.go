package main

type FuncionarioAssalariado struct {
	nome          string
	salarioMensal float64
}

func (f FuncionarioAssalariado) calcularSalario() float64 {
	return f.salarioMensal
}
