package main

import "fmt"

type Contrato struct {
	partes []string
}

// Método Imprimir para a struct Contrato
func (c Contrato) Imprimir() {
	fmt.Printf("Contrato entre: %s\n", join(c.partes, ", "))
}

// Função join para concatenar partes de string com um delimitador
func join(partes []string, delimitador string) string {
	resultado := ""
	for i, parte := range partes {
		if i > 0 {
			resultado += delimitador
		}
		resultado += parte
	}
	return resultado
}
