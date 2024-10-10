#Crie uma classe Carro com atributos como marca, modelo, e
#ano. Instancie três objetos dessa classe e exiba os detalhes de cada um.

from questao06 import Motor

class Carro:
    def __init__(self, marca, modelo, cor, motor):
        self.marca = marca
        self.modelo = modelo
        self.cor = cor
        self.motor = motor

    def exibir_detalhes(self):
        print(f'Marca: {self.marca}, Modelo: {self.modelo}, Cor: {self.cor}, motor: {self.motor.exibe_motor()}')

motor1 = Motor("1.0", "flex")
motor2 = Motor("1.4", "gasolina")
motor3 = Motor("1.3", "etanol")


carro1 = Carro("Volkswagen", "Polo", "Preto", motor1)
carro2 = Carro("Fiat", "Strada", "Branco", motor2)
carro3 = Carro("Toyota", "Etios", "Prata", motor3)

carro1.exibir_detalhes()
carro2.exibir_detalhes()
carro3.exibir_detalhes()