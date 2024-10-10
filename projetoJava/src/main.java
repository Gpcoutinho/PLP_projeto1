import java.util.ArrayList;
import java.util.List;


public class main {
    public static void main(String[] args) {
   /* #1, 2, 6// Exercicio classe carro e motor

        Motor motor1 = new Motor("1.2", "flex");
        Motor motor2 = new Motor("1.5", "gasolina");
        Motor motor3 = new Motor("1.6", "etanol");

        Carro carro1 = new Carro("Honda", "Civic", "Azul", motor1);
        Carro carro2 = new Carro("Nissan", "Versa", "Prata", motor2);
        Carro carro3 = new Carro("Hyundai", "HB20", "Vermelho", motor3);

        carro1.exibirDetalhes();
        carro2.exibirDetalhes();
        carro3.exibirDetalhes();

        carro1.acelerar(50);
        carro1.exibirVelocidade();

        carro1.frear(20);
        carro1.exibirVelocidade();

        carro1.frear(40);
        carro1.exibirVelocidade();
   */

   /* #3// Exercicio classe conta bancaria
        ContaBancaria conta1 = new ContaBancaria("Gabriel");

        conta1.depositar(500);
        conta1.exibirSaldo();

        conta1.sacar(200);
        conta1.exibirSaldo();

        conta1.sacar(400);
        conta1.exibirSaldo();
   */

   /* #4, 5// Exercicio classe animal

        List<Animal> animais = new ArrayList<>();

        Cachorro cachorro = new Cachorro("Appa");
        Gato gato = new Gato("Mushu");
        Pato pato = new Pato("Zezinho");

        animais.add(cachorro);
        animais.add(gato);
        animais.add(pato);

        exibirAnimais(animais);
    }
    public static void exibirAnimais(List<Animal> animais) {
        for (Animal animal : animais) {
            System.out.println(animal.nome + " diz: " + animal.som());
        }

   */

   /* #7// Exercicio classe escola e professor

        Professor prof1 = new Professor("Gabriel");
        Professor prof2 = new Professor("Ricardo");
        Professor prof3 = new Professor("Carlos");

        Escola escola1 = new Escola("Evolução");
        Escola escola2 = new Escola("Século");

        escola1.adicionarProfessor(prof1);
        escola1.adicionarProfessor(prof2);

        escola2.adicionarProfessor(prof2);
        escola2.adicionarProfessor(prof3);

        escola1.exibirProfessores();
        escola2.exibirProfessores();

        prof1.exibirEscolas();
        prof2.exibirEscolas();
        prof3.exibirEscolas();

   */

   /* #8// Exercicio classe empresa e empregado
        Empresa empresa = new Empresa("Tech Solutions");

        Empregado empregado1 = new Empregado("Gabriel", "Gerente", 1000);
        Empregado empregado2 = new Empregado("Ricardo", "Desenvolvedor", 3000);
        Empregado empregado3 = new Empregado("Carlos", "Analista de testes", 2000);

        empresa.adicionarEmpregado(empregado1);
        empresa.adicionarEmpregado(empregado2);
        empresa.adicionarEmpregado(empregado3);

        System.out.println(empresa.infoEmpresa());

   */

   /* #9// Exercicio interface
        Relatorio relatorio = new Relatorio("Análise de vendas do Q3");
        Contrato contrato = new Contrato(new String[]{"Empresa X", "Empresa Y"});

        relatorio.imprimir();
        contrato.imprimir();

   */

   /* #10// Exercicio calculadora
        Calculadora calc = new Calculadora();

        int resultado1 = calc.somar(10, 20);
        System.out.println("Resultado da soma de dois números: " + resultado1);

        int resultado2 = calc.somar(5, 10, 15);
        System.out.println("Resultado da soma de três números: " + resultado2);

   */

   /* #11// Exercicio classe abstrata funcionario
        FuncionarioHorista funcionarioHorista = new FuncionarioHorista("Gabriel", 160, 15.00);
        System.out.printf("Salário do funcionário horista %s: R$ %.2f%n", funcionarioHorista.nome, funcionarioHorista.calcularSalario());

        FuncionarioAssalariado funcionarioAssalariado = new FuncionarioAssalariado("Ricardo", 3000.00);
        System.out.printf("Salário do funcionário assalariado %s: R$ %.2f%n", funcionarioAssalariado.nome, funcionarioAssalariado.calcularSalario());

   */

   /* #12// Exercicio sobrecarga de operador
        Produto produto1 = new Produto("Tapioca", 50.00);
        Produto produto2 = new Produto("Goiaba", 30.00);

        Produto produto3 = produto1.somar(produto2);

        System.out.println(produto1);
        System.out.println(produto2);
        System.out.println(produto3);

    */

    /* #13// Exercicio calcular fatorial com metodo estatico
        int numero = 5;
            int resultado = Matematica.fatorial(numero);
            System.out.println("O fatorial de " + numero + " é: " + resultado);

            try {
                int numeroNegativo = -3;
                int resultadoNegativo = Matematica.fatorial(numeroNegativo);
                System.out.println("O fatorial de " + numeroNegativo + " é: " + resultadoNegativo);
            } catch (IllegalArgumentException e) {
                System.out.println(e.getMessage());
            }

    */

    /* #14// Exercicio Singleton
        Configuracao config1 = Configuracao.getInstancia();
                config1.setConfiguracao("cor_tema", "azul");

                Configuracao config2 = Configuracao.getInstancia();
                config2.setConfiguracao("font_size", "12px");

                System.out.println(config1.getConfiguracao("cor_tema"));
                System.out.println(config2.getConfiguracao("font_size"));

                System.out.println(config1 == config2);
    */

    /* #15// Exercicio classe exception

        try {
            ContaCorrente conta = new ContaCorrente(100);
            conta.sacar(150);
        } catch (SaldoInsuficienteException e) {
            System.out.println(e.getMessage());
        }

    */

    }
}



