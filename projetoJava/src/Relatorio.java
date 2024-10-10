interface Imprimivel {
    void imprimir();
}


class Relatorio implements Imprimivel {
    private String conteudo;

    public Relatorio(String conteudo) {
        this.conteudo = conteudo;
    }

    public void imprimir() {
        System.out.println("Relatório: " + conteudo);
    }
}