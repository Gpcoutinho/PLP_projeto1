#Crie uma classe base Animal com metodos como som. Derive classes como
#Cachorro e Gato que implementam o metodo som de maneiras diferentes.

class Animal:
    def __init__(self, nome):
        self.nome = nome

    def som(self):
        pass

class Cachorro(Animal):
    def som(self):
        return "au au!"

class Gato(Animal):
    def som(self):
        return "miau!"

class Pato(Animal):
    def som(self):
        return "quack!"

cachorro = Cachorro("Appa")
gato = Gato("Mushu")
pato = Pato("Zezinho")

print(f"{cachorro.nome} diz: {cachorro.som()}")
print(f"{gato.nome} diz: {gato.som()}")
print(f"{pato.nome} diz: {pato.som()}")

