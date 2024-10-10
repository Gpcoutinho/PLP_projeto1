class Contrato implements Imprimivel {
    private String[] partes;

    public Contrato(String[] partes) {
        this.partes = partes;
    }

    public void imprimir() {
        System.out.print("Contrato entre: ");
        for (int i = 0; i < partes.length; i++) {
            System.out.print(partes[i]);
            if (i < partes.length - 1) {
                System.out.print(", ");
            }
        }
        System.out.println();
    }
}