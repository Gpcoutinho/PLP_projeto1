import java.util.ArrayList;
import java.util.List;


class Empresa {
    private String nome;
    private List<Empregado> empregados;

    public Empresa(String nome) {
        this.nome = nome;
        this.empregados = new ArrayList<>();
    }

    public void adicionarEmpregado(Empregado empregado) {
        empregados.add(empregado);
    }

    public String infoEmpresa() {
        StringBuilder info = new StringBuilder("Empresa: " + nome + "\nEmpregados:\n");
        for (Empregado empregado : empregados) {
            info.append("- ").append(empregado.infoEmpregado()).append("\n");
        }
        return info.toString();
    }
}