public class Carro {
    private String marca;
    private String modelo;
    private String cor;
    private Motor motor;
    private int velocidade;

    public Carro(String marca, String modelo, String cor, Motor motor) {
        this.marca = marca;
        this.modelo = modelo;
        this.cor = cor;
        this.motor = motor;
        this.velocidade = 0;
    }

    public void acelerar(int incremento) {
        this.velocidade += incremento;
        System.out.println("O carro acelerou para " + velocidade + " km/h.");
    }

    public void frear(int decremento) {
        this.velocidade -= decremento;
        if (this.velocidade < 0) {
            this.velocidade = 0;
        }
        System.out.println("O carro reduziu a velocidade para " + velocidade + " km/h.");
    }

    public void exibirDetalhes() {
        System.out.println("Marca: " + marca + ", Modelo: " + modelo + ", Cor: " + cor + ", Motor: " + motor.exibeMotor());
    }

    public void exibirVelocidade() {
        System.out.println("A velocidade atual do " + marca + " " + modelo + " é " + velocidade + " km/h.");
    }
}
