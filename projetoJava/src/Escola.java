import java.util.ArrayList;
import java.util.List;

class Escola {
    private String nome;
    private List<Professor> professores;

    public Escola(String nome) {
        this.nome = nome;
        this.professores = new ArrayList<>();
    }

    public void adicionarProfessor(Professor professor) {
        if (!professores.contains(professor)) {
            professores.add(professor);
            professor.adicionarEscola(this);
        }
    }

    public void exibirProfessores() {
        System.out.println("Escola " + nome + " tem os seguintes professores:");
        for (Professor professor : professores) {
            System.out.println("- " + professor.getNome());
        }
    }

    public String getNome() {
        return nome;
    }
}