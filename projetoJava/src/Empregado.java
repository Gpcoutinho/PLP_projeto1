class Empregado {
    private String nome;
    private String cargo;
    private double salario;

    public Empregado(String nome, String cargo, double salario) {
        this.nome = nome;
        this.cargo = cargo;
        this.salario = salario;
    }

    public String infoEmpregado() {
        return nome + ", Cargo: " + cargo + ", Salário: R$ " + salario;
    }
}