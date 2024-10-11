package main

import "fmt"

type Escola struct {
	nome        string
	professores []*Professor
}

func (e *Escola) AdicionarProfessor(professor *Professor) {
	if !e.temProfessor(professor) {
		e.professores = append(e.professores, professor)
		professor.AdicionarEscola(e)
	}
}

func (e *Escola) temProfessor(professor *Professor) bool {
	for _, p := range e.professores {
		if p == professor {
			return true
		}
	}
	return false
}

func (e *Escola) ExibirProfessores() {
	fmt.Printf("Escola %s tem os seguintes professores:\n", e.nome)
	for _, professor := range e.professores {
		fmt.Printf("- %s\n", professor.nome)
	}
}
