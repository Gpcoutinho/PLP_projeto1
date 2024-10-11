package main

import "fmt"

type ContaBancaria struct {
	Titular string
	saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
		fmt.Printf("Depósito de R$ %.2f realizado com sucesso.\n", valor)
	} else {
		fmt.Println("O valor do depósito deve ser positivo.")
	}
}

func (c *ContaBancaria) Sacar(valor float64) {
	if valor > 0 && valor <= c.saldo {
		c.saldo -= valor
		fmt.Printf("Saque de R$ %.2f realizado com sucesso.\n", valor)
	} else {
		fmt.Println("Saque não permitido. Verifique o valor e o saldo.")
	}
}

func (c ContaBancaria) ExibirSaldo() {
	fmt.Printf("O saldo atual da conta de %s é R$ %.2f.\n", c.Titular, c.saldo)
}
