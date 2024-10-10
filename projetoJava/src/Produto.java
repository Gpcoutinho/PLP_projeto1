
public class Produto {
    private String nome;
    private double preco;

    public Produto(String nome, double preco) {
        this.nome = nome;
        this.preco = preco;
    }

    public Produto somar(Produto outro) {
        String novoNome = this.nome + " e " + outro.nome;
        double novoPreco = this.preco + outro.preco;
        return new Produto(novoNome, novoPreco);
    }

    @Override
    public String toString() {
        return nome + ": R$ " + String.format("%.2f", preco);
    }
}
