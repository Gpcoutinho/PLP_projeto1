class FuncionarioHorista extends Funcionario{
    private int horasTrabalhadas;
    private double valorHora;

    public FuncionarioHorista(String nome, int horasTrabalhadas, double valorHora) {
        this.nome = nome;
        this.horasTrabalhadas = horasTrabalhadas;
        this.valorHora = valorHora;


    }

    @Override
    public double calcularSalario() {
        return horasTrabalhadas * valorHora;
    }
}

