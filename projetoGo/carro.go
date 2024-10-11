package main

import "fmt"

type Carro struct {
	Marca      string
	Modelo     string
	Cor        string
	Motor      Motor
	Velocidade int
}

func (c *Carro) Acelerar(incremento int) {
	c.Velocidade += incremento
	fmt.Printf("O carro acelerou para %d km/h.\n", c.Velocidade)
}

func (c *Carro) Frear(decremento int) {
	c.Velocidade -= decremento
	if c.Velocidade < 0 {
		c.Velocidade = 0
	}
	fmt.Printf("O carro reduziu a velocidade para %d km/h.\n", c.Velocidade)
}

func (c Carro) ExibirVelocidade() {
	fmt.Printf("A velocidade atual do %s %s é %d km/h.\n", c.Marca, c.Modelo, c.Velocidade)
}

func (c Carro) ExibirDetalhes() {
	fmt.Printf("Marca: %s, Modelo: %s, Cor: %s, Motor: %s\n", c.Marca, c.Modelo, c.Cor, c.Motor.ExibeMotor())
}
