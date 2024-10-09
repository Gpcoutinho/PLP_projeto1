#Adicione um metodo acelerar e frear à classe Carro que altere um atributo
#velocidade. Crie um metodo para exibir a velocidade atual.

class Carro:
    def __init__(self, marca, modelo, cor):
        self.marca = marca
        self.modelo = modelo
        self.cor = cor
        self.velocidade = 0

    def acelerar(self, incremento):
        self.velocidade += incremento
        print(f'O carro acelerou para {self.velocidade} km/h.')

    def frear(self, decremento):
        self.velocidade -= decremento
        if self.velocidade < 0:
            self.velocidade = 0
        print(f'O carro reduziu a velocidade para {self.velocidade} km/h.')

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Cor: {self.cor}')

    def exibir_velocidade(self):
        print(f'A velocidade atual do {self.marca} {self.modelo} é {self.velocidade} km/h.')


carro1 = Carro("Volkswagen", "Polo", "Preto")
carro2 = Carro("Fiat", "Strada", "Branco")
carro3 = Carro("Toyota", "Etios", "Prata")

carro1.exibir_detalhes()
carro2.exibir_detalhes()
carro3.exibir_detalhes()

carro1.acelerar(50)
carro1.exibir_velocidade()

carro1.frear(20)
carro1.exibir_velocidade()

carro1.frear(40)
carro1.exibir_velocidade()
