package main

type ContaCorrente struct {
	saldo float64
}

func NovoContaCorrente(saldoInicial float64) *ContaCorrente {
	return &ContaCorrente{saldo: saldoInicial}
}

func (c *ContaCorrente) Sacar(valor float64) error {
	if valor > c.saldo {
		return &SaldoInsuficienteException{saldoAtual: c.saldo, valorSaque: valor}
	}
	c.saldo -= valor
	return nil
}

func (c *ContaCorrente) Depositar(valor float64) {
	c.saldo += valor
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
