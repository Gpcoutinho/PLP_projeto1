package main

import "errors"

type Matematica struct{}

func (Matematica) Fatorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("o fatorial não está definido para números negativos")
	} else if n == 0 || n == 1 {
		return 1, nil
	}

	fatorial := 1
	for i := 2; i <= n; i++ {
		fatorial *= i
	}
	return fatorial, nil
}
