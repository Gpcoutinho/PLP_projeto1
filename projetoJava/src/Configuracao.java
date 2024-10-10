import java.util.HashMap;

public class Configuracao {

    private static Configuracao instancia = null;

    private HashMap<String, String> configuracoes;

    private Configuracao() {
        configuracoes = new HashMap<>();
    }

    public static Configuracao getInstancia() {
        if (instancia == null) {
            instancia = new Configuracao();
        }
        return instancia;
    }

    public void setConfiguracao(String chave, String valor) {
        configuracoes.put(chave, valor);
    }

    public String getConfiguracao(String chave) {
        return configuracoes.getOrDefault(chave, null);
    }
}
