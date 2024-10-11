package main

import "fmt"

type Carro struct {
	Marca  string
	Modelo string
	Cor    string
	Motor  Motor
}

func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Cor: %s, Motor: %s\n", c.Marca, c.Modelo, c.Cor, c.Motor.ExibeMotor())
}
