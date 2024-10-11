package main

type FuncionarioHorista struct {
	nome             string
	horasTrabalhadas float64
	valorPorHora     float64
}

func (f FuncionarioHorista) calcularSalario() float64 {
	return f.horasTrabalhadas * f.valorPorHora
}
