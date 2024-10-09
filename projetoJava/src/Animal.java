public abstract class Animal {
    protected String nome;

    public Animal(String nome) {
        this.nome = nome;
    }

    public abstract String som();
}

class Cachorro extends Animal {
    public Cachorro(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return "au au!";
    }
}

class Gato extends Animal {
    public Gato(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return "miau!";
    }
}

class Pato extends Animal {
    public Pato(String nome) {
        super(nome);
    }

    @Override
    public String som() {
        return "quack!";
    }
}



