package main

type Animal interface {
	Som() string
	Nome() string
}

type BaseAnimal struct {
	nome string
}

func (a BaseAnimal) Nome() string {
	return a.nome
}
