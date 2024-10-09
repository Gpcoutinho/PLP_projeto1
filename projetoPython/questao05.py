#Utilize polimorfismo para criar uma função que receba uma lista de
#objetos Animal e chame o metodo som de cada um, demonstrando comportamento
#diferente para cada subclasse.


from questao04 import Cachorro, Gato, Pato

def exibir_animais(animais):
    for animal in animais:
        print(f"{animal.nome} diz: {animal.som()}")

cachorro = Cachorro("Appa")
gato = Gato("Mushu")
pato = Pato("Zezinho")

animais = [cachorro, gato, pato]
