package main

import "sync"

type Configuracao struct {
	configuracoes map[string]string
}

var instancia *Configuracao

var mu sync.Mutex

func GetConfiguracao() *Configuracao {
	mu.Lock()
	defer mu.Unlock()

	if instancia == nil {
		instancia = &Configuracao{
			configuracoes: make(map[string]string),
		}
	}
	return instancia
}

func (c *Configuracao) SetConfiguracao(chave, valor string) {
	c.configuracoes[chave] = valor
}

func (c *Configuracao) GetConfiguracao(chave string) string {
	return c.configuracoes[chave]
}
