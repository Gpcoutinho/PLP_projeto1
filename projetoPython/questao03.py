#Implemente uma classe ContaBancaria com atributos saldo, titular e
#metodos depositar e sacar. Use encapsulamento para proteger o atributo saldo.

class ContaBancaria:
    def __init__(self, titular):
        self.titular = titular
        self.__saldo = 0.0

    def depositar(self, valor):
        if valor > 0:
            self.__saldo += valor
            print(f'Depósito de R$ {valor:.2f} realizado com sucesso.')
        else:
            print('O valor do depósito deve ser positivo.')

    def sacar(self, valor):
        if 0 < valor <= self.__saldo:
            self.__saldo -= valor
            print(f'Saque de R$ {valor:.2f} realizado com sucesso.')
        else:
            print('Saque não permitido. Verifique o valor e o saldo.')

    def exibir_saldo(self):
        print(f'O saldo atual da conta de {self.titular} é R$ {self.__saldo:.2f}.')

conta1 = ContaBancaria("Gabriel")
conta1.depositar(500)
conta1.exibir_saldo()

conta1.sacar(200)
conta1.exibir_saldo()

conta1.sacar(400)
conta1.exibir_saldo()
