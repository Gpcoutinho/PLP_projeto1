import java.util.ArrayList;
import java.util.List;

class Professor {
    private String nome;
    private List<Escola> escolas;

    public Professor(String nome) {
        this.nome = nome;
        this.escolas = new ArrayList<>();
    }

    public void adicionarEscola(Escola escola) {
        if (!escolas.contains(escola)) {
            escolas.add(escola);
            escola.adicionarProfessor(this);
        }
    }

    public void exibirEscolas() {
        System.out.println("Professor " + nome + " leciona nas seguintes escolas:");
        for (Escola escola : escolas) {
            System.out.println("- " + escola.getNome());
        }
    }

    public String getNome() {
        return nome;
    }
}