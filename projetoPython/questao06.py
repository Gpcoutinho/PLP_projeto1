#Implemente uma classe Motor e associe-a a uma classe Carro. A classe
#Carro deve ter um objeto do tipo Motor como um de seus atributos.

class Motor:
    def __init__(self, volume, combustivel):
        self.volume = volume
        self.combustivel = combustivel

    def exibe_motor(self):
        return f"Volume: {self.volume} Combustivel: {self.combustivel}"
