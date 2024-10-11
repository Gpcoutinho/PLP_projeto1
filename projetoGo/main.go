package main

import (
	"fmt"
)

func ExibirAnimais(animais []Animal) {
	for _, animal := range animais {
		fmt.Printf("%s diz: %s\n", animal.Nome(), animal.Som())
	}
}

func main() {
	/* #1, 2, 6 Exercicio classes carro e motor
	motor1 := Motor{Cilindrada: "1.0", Combustivel: "flex"}
	motor2 := Motor{Cilindrada: "1.4", Combustivel: "gasolina"}
	motor3 := Motor{Cilindrada: "1.3", Combustivel: "etanol"}

	carro1 := Carro{Marca: "Volkswagen", Modelo: "Polo", Cor: "Preto", Motor: motor1}
	carro2 := Carro{Marca: "Fiat", Modelo: "Strada", Cor: "Branco", Motor: motor2}
	carro3 := Carro{Marca: "Toyota", Modelo: "Etios", Cor: "Prata", Motor: motor3}

	carro1.ExibirDetalhes()
	carro2.ExibirDetalhes()
	carro3.ExibirDetalhes()

	carro1.Acelerar(50)
	carro1.ExibirVelocidade()

	carro1.Frear(20)
	carro1.ExibirVelocidade()

	carro1.Frear(40)
	carro1.ExibirVelocidade()
	*/

	/* #3 Exercicio classe ContaBancaria
	conta1 := ContaBancaria{Titular: "Gabriel"}

	conta1.Depositar(500)
	conta1.ExibirSaldo()

	conta1.Sacar(200)
	conta1.ExibirSaldo()

	conta1.Sacar(400)
	conta1.ExibirSaldo()
	*/

	/* #4, 5 Exercicio classe animal
	cachorro := Cachorro{BaseAnimal{nome: "Appa"}}
	gato := Gato{BaseAnimal{nome: "Mushu"}}
	pato := Pato{BaseAnimal{nome: "Zezinho"}}

	animais := []Animal{cachorro, gato, pato}

	ExibirAnimais(animais)
	*/

	/* #7 Classes escola e professor
	prof1 := &Professor{nome: "Gabriel"}
	prof2 := &Professor{nome: "Ricardo"}
	prof3 := &Professor{nome: "Carlos"}

	escola1 := &Escola{nome: "Evolução"}
	escola2 := &Escola{nome: "Século"}

	escola1.AdicionarProfessor(prof1)
	escola1.AdicionarProfessor(prof2)

	escola2.AdicionarProfessor(prof2)
	escola2.AdicionarProfessor(prof3)

	escola1.ExibirProfessores()
	escola2.ExibirProfessores()

	prof1.ExibirEscolas()
	prof2.ExibirEscolas()
	prof3.ExibirEscolas()
	*/

	/*#8 Classes empresa e empregado
	empresa := Empresa{nome: "Tech Solutions"}

	empregado1 := &Empregado{nome: "Gabriel", cargo: "Gerente", salario: 1000}
	empregado2 := &Empregado{nome: "Ricardo", cargo: "Desenvolvedor", salario: 3000}
	empregado3 := &Empregado{nome: "Carlos", cargo: "Analista de testes", salario: 2000}

	empresa.AdicionarEmpregado(empregado1)
	empresa.AdicionarEmpregado(empregado2)
	empresa.AdicionarEmpregado(empregado3)

	fmt.Println(empresa.InfoEmpresa())
	*/

	/*#9 Interface
	relatorio := Relatorio{conteudo: "Análise de vendas do Q3"}
	contrato := Contrato{partes: []string{"Empresa X", "Empresa Y"}}

	relatorio.Imprimir()
	contrato.Imprimir()
	*/

	/*#10 Classe calculadora
	calc := Calculadora{}

	resultadoDois := calc.SomarDois(10.0, 20.0)
	fmt.Printf("A soma de dois números é: %.2f\n", resultadoDois)

	resultadoTres := calc.SomarTres(5.0, 10.0, 15.0)
	fmt.Printf("A soma de três números é: %.2f\n", resultadoTres)
	*/

	/*#11 Classe abstrata funcionario
	funcionarioHorista := FuncionarioHorista{"Gabriel", 160, 15.00}
	fmt.Printf("Salário do funcionário horista %s: R$ %.2f\n", funcionarioHorista.nome, funcionarioHorista.calcularSalario())

	funcionarioAssalariado := FuncionarioAssalariado{"Ricardo", 3000.00}
	fmt.Printf("Salário do funcionário assalariado %s: R$ %.2f\n", funcionarioAssalariado.nome, funcionarioAssalariado.calcularSalario())
	*/

	/*#12 Metodos de conveniencia
	produto1 := NovoProduto("Tapioca", 50.00)
	produto2 := NovoProduto("Goiaba", 30.00)

	produto3 := produto1.Somar(produto2)

	fmt.Println(produto1)
	fmt.Println(produto2)
	fmt.Println(produto3)
	*/

	/*#13 Metodos estaticos
	m := Matematica{}

	numero := 5
	resultado, err := m.Fatorial(numero)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O fatorial de %d é: %d\n", numero, resultado)
	}

	numeroNegativo := -3
	resultadoNegativo, err := m.Fatorial(numeroNegativo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("O fatorial de %d é: %d\n", numeroNegativo, resultadoNegativo)
	}

	*/

	/*#14 padrão Singleton
	config1 := GetConfiguracao()
	config1.SetConfiguracao("cor_tema", "azul")

	config2 := GetConfiguracao()
	config2.SetConfiguracao("font_size", "12px")

	fmt.Println(config1.GetConfiguracao("cor_tema"))
	fmt.Println(config2.GetConfiguracao("font_size"))

	fmt.Println(config1 == config2)
	*/

	/*#15 Classe de exceção personalizada
	conta := NovoContaCorrente(100)
	err := conta.Sacar(150)
	if err != nil {
		fmt.Println(err)
	*/
}
