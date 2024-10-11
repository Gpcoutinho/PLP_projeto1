package main

import "fmt"

type SaldoInsuficienteException struct {
	saldoAtual float64
	valorSaque float64
}

func (e *SaldoInsuficienteException) Error() string {
	return fmt.Sprintf("Saldo insuficiente: saldo atual é R$ %.2f, tentativa de saque R$ %.2f.", e.saldoAtual, e.valorSaque)
}
