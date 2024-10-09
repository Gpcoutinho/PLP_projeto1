#Sobrecarga de Métodos (Java) / Métodos com Nomes Diferentes (Python, Go)
#Implemente uma classe Calculadora com metodos somar que aceita diferentes números
#de parâmetros (dois ou três números). Em Golang, use funções com nomes diferentes
#para diferentes quantidades de parâmetros.

class Calculadora:

    def somar(self, *numeros):
        if len(numeros) < 2:
            raise ValueError("Informe pelo menos dois números para somar.")
        elif len(numeros) > 3:
            raise ValueError("Somente até três números são permitidos.")

        return sum(numeros)

calc = Calculadora()

resultado_2 = calc.somar(10, 20)
print(f"Soma de dois números: {resultado_2}")

resultado_3 = calc.somar(5, 10, 15)
print(f"Soma de três números: {resultado_3}")

try:
    calc.somar(5)
except ValueError as e:
    print(e)

try:
    calc.somar(1, 2, 3, 4)
except ValueError as e:
    print(e)
