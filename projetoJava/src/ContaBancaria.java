public class ContaBancaria {
    private double saldo;
    private String titular;

    public ContaBancaria(String titular) {
        this.titular = titular;
        this.saldo = 0.0;
    }

    public void depositar(double valor) {
        if (valor > 0) {
            saldo += valor;
            System.out.printf("Depósito de R$ %.2f realizado com sucesso.%n", valor);
        } else {
            System.out.println("O valor do depósito deve ser positivo.");
        }
    }

    public void sacar(double valor) {
        if (valor > 0 && valor <= saldo) {
            saldo -= valor;
            System.out.printf("Saque de R$ %.2f realizado com sucesso.%n", valor);
        } else {
            System.out.println("Saque não permitido. Verifique o valor e o saldo.");
        }
    }

    public void exibirSaldo() {
        System.out.printf("O saldo atual da conta de %s é R$ %.2f.%n", titular, saldo);
    }

    public static void main(String[] args) {
        ContaBancaria conta1 = new ContaBancaria("Gabriel");
        conta1.depositar(500);
        conta1.exibirSaldo();

        conta1.sacar(200);
        conta1.exibirSaldo();

        conta1.sacar(400);
        conta1.exibirSaldo();
    }
}
