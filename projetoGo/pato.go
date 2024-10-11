package main

type Pato struct {
	BaseAnimal
}

func (p Pato) Som() string {
	return "quack!"
}
