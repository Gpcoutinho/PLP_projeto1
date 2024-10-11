package main

type Calculadora struct{}

func (c *Calculadora) SomarDois(x, y float64) float64 {
	return x + y
}

func (c *Calculadora) SomarTres(x, y, z float64) float64 {
	return x + y + z
}
