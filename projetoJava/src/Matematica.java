public class Matematica {

    public static int fatorial(int n) throws IllegalArgumentException {
        if (n < 0) {
            throw new IllegalArgumentException("O fatorial não está definido para números negativos.");
        } else if (n == 0 || n == 1) {
            return 1;
        } else {
            int fatorial = 1;
            for (int i = 2; i <= n; i++) {
                fatorial *= i;
            }
            return fatorial;
        }
    }
}