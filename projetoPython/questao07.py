#Crie uma classe Escola e uma classe Professor. A associação deve permitir
#que uma escola tenha vários professores, e um professor possa lecionar em várias escolas.

# professor.py
class Professor:
    def __init__(self, nome):
        self.nome = nome
        self.escolas = []

    def adicionar_escola(self, escola):
        if escola not in self.escolas:
            self.escolas.append(escola)
            escola.adicionar_professor(self)

    def exibir_escolas(self):
        print(f"Professor {self.nome} leciona nas seguintes escolas:")
        for escola in self.escolas:
            print(f"- {escola.nome}")

class Escola:
    def __init__(self, nome):
        self.nome = nome
        self.professores = []

    def adicionar_professor(self, professor):
        if professor not in self.professores:
            self.professores.append(professor)
            professor.adicionar_escola(self)

    def exibir_professores(self):
        print(f"Escola {self.nome} tem os seguintes professores:")
        for professor in self.professores:
            print(f"- {professor.nome}")


prof1 = Professor("Gabriel")
prof2 = Professor("Ricardo")
prof3 = Professor("Carlos")

escola1 = Escola("Evolução")
escola2 = Escola("Século")

escola1.adicionar_professor(prof1)
escola1.adicionar_professor(prof2)

escola2.adicionar_professor(prof2)
escola2.adicionar_professor(prof3)

escola1.exibir_professores()
escola2.exibir_professores()

prof1.exibir_escolas()
prof2.exibir_escolas()
prof3.exibir_escolas()