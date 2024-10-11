package main

type Cachorro struct {
	BaseAnimal
}

func (c Cachorro) Som() string {
	return "au au!"
}
