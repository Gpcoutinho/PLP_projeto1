#Modele uma classe Empresa que agregue uma lista de objetos Empregado.
#Cada empregado deve ter atributos como nome, cargo, e salario.


class Empregado:
    def __init__(self, nome, cargo, salario):
        self.nome = nome
        self.cargo = cargo
        self.salario = salario

    def info_empregado(self):
        return f"{self.nome}, Cargo: {self.cargo}, Salário: R$ {self.salario}"


class Empresa:
    def __init__(self, nome):
        self.nome = nome
        self.empregados = []

    def adicionar_empregado(self, empregado):
        self.empregados.append(empregado)

    def info_empresa(self):
        info = f"Empresa: {self.nome}\nEmpregados:\n"
        for empregado in self.empregados:
            info += f"- {empregado.info_empregado()}\n"
        return info

empresa = Empresa("Tech Solutions")

empregado1 = Empregado("Gabriel", "Gerente", 1000)
empregado2 = Empregado("Ricardo", "Desenvolvedor", "2 cremosinhos")
empregado3 = Empregado("Carlos", "Analista de testes", 2000)

empresa.adicionar_empregado(empregado1)
empresa.adicionar_empregado(empregado2)
empresa.adicionar_empregado(empregado3)

print(empresa.info_empresa())
