public class Motor {
    private String cilindrada;
    private String tipoCombustivel;

    public Motor(String cilindrada, String tipoCombustivel) {
        this.cilindrada = cilindrada;
        this.tipoCombustivel = tipoCombustivel;
    }

    public String exibeMotor() {
        return cilindrada + " - " + tipoCombustivel;
    }
}
