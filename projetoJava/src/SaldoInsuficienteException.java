
public class SaldoInsuficienteException extends Exception {
    private double saldoAtual;
    private double valorSaque;

    public SaldoInsuficienteException(double saldoAtual, double valorSaque) {
        super("Saldo insuficiente: saldo atual é R$ " + saldoAtual + ", tentativa de saque R$ " + valorSaque + ".");
        this.saldoAtual = saldoAtual;
        this.valorSaque = valorSaque;
    }

    public double getSaldoAtual() {
        return saldoAtual;
    }

    public double getValorSaque() {
        return valorSaque;
    }
}


