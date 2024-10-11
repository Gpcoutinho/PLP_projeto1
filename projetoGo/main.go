package main

func main() {
	motor1 := Motor{Cilindrada: "1.0", Combustivel: "flex"}
	motor2 := Motor{Cilindrada: "1.4", Combustivel: "gasolina"}
	motor3 := Motor{Cilindrada: "1.3", Combustivel: "etanol"}

	carro1 := Carro{Marca: "Volkswagen", Modelo: "Polo", Cor: "Preto", Motor: motor1}
	carro2 := Carro{Marca: "Fiat", Modelo: "Strada", Cor: "Branco", Motor: motor2}
	carro3 := Carro{Marca: "Toyota", Modelo: "Etios", Cor: "Prata", Motor: motor3}

	carro1.ExibirDetalhes()
	carro2.ExibirDetalhes()
	carro3.ExibirDetalhes()
}
