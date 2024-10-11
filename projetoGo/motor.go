package main

import "fmt"

type Motor struct {
	Cilindrada  string
	Combustivel string
}

func (m Motor) ExibeMotor() string {
	return fmt.Sprintf("%s, %s", m.Cilindrada, m.Combustivel)
}
