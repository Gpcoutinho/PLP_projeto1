package main

type Gato struct {
	BaseAnimal
}

func (g Gato) Som() string {
	return "miau!"
}
