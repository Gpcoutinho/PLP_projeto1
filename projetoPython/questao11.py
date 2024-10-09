#Crie uma classe abstrata Funcionario com um metodo abstrato
#calcularSalario. Derive classes como FuncionarioHorista e FuncionarioAssalariado que
#implementam calcularSalario de formas distintas.

from abc import ABC, abstractmethod


class Funcionario(ABC):

    @abstractmethod
    def calcularSalario(self):
        pass

class FuncionarioHorista(Funcionario):

    def __init__(self, nome, horas_trabalhadas, valor_por_hora):
        self.nome = nome
        self.horas_trabalhadas = horas_trabalhadas
        self.valor_por_hora = valor_por_hora

    def calcularSalario(self):
        return self.horas_trabalhadas * self.valor_por_hora

class FuncionarioAssalariado(Funcionario):

    def __init__(self, nome, salario_mensal):
        self.nome = nome
        self.salario_mensal = salario_mensal

    def calcularSalario(self):
        return self.salario_mensal


funcionario_horista = FuncionarioHorista("Gabriel", 160, 15.00)
print(f"Salário do funcionário horista {funcionario_horista.nome}: R$ {funcionario_horista.calcularSalario():.2f}")

funcionario_assalariado = FuncionarioAssalariado("Ricardo", 3000.00)
print(f"Salário do funcionário assalariado {funcionario_assalariado.nome}: R$ {funcionario_assalariado.calcularSalario():.2f}")
