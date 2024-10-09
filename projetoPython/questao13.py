#Adicione um metodo estático à classe Matematica que calcula o
#fatorial de um número. Em Python, utilize @staticmethod, em Java static, e em Golang crie
#uma função regular na struct.

class Matematica:

    @staticmethod
    def fatorial(n):
        if n < 0:
            raise ValueError("O fatorial não está definido para números negativos.")
        elif n == 0 or n == 1:
            return 1
        else:
            fatorial = 1
            for i in range(2, n + 1):
                fatorial *= i
            return fatorial

try:
    numero = 5
    resultado = Matematica.fatorial(numero)
    print(f"O fatorial de {numero} é: {resultado}")
except ValueError as e:
    print(e)

try:
    numero_negativo = -3
    resultado_negativo = Matematica.fatorial(numero_negativo)
    print(f"O fatorial de {numero_negativo} é: {resultado_negativo}")
except ValueError as e:
    print(e)

