#Em Java e Golang, defina uma interface para imprimivel que
#tenha um metodo imprimir. Implemente essa interface em classes como Relatório e
#Contrato. Em Python, use a abordagem de protocolo ou classes abstratas.


from abc import ABC, abstractmethod

class Imprimivel(ABC):

    @abstractmethod
    def imprimir(self):
        pass

class Relatorio(Imprimivel):

    def __init__(self, conteudo):
        self.conteudo = conteudo

    def imprimir(self):
        print(f"Relatório: {self.conteudo}")

class Contrato(Imprimivel):

    def __init__(self, partes):
        self.partes = partes

    def imprimir(self):
        print(f"Contrato entre: {', '.join(self.partes)}")

relatorio = Relatorio("Análise de vendas do Q3")
contrato = Contrato(["Empresa X", "Empresa Y"])

relatorio.imprimir()
contrato.imprimir()
