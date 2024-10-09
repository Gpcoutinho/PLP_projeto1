#Crie uma classe de exceção personalizada
#SaldoInsuficienteException em Java e Python, ou error em Golang, que seja lançada
#quando houver uma tentativa de saque superior ao saldo disponível na classe
#ContaBancaria.

class SaldoInsuficienteException(Exception):
    def __init__(self, saldo_atual, valor_saque):
        self.saldo_atual = saldo_atual
        self.valor_saque = valor_saque
        super().__init__(f'Saldo insuficiente: saldo atual é R$ {saldo_atual}, tentativa de saque R$ {valor_saque}.')

class ContaBancaria:
    def __init__(self, saldo_inicial=0):
        self.saldo = saldo_inicial

    def sacar(self, valor):
        if valor > self.saldo:
            raise SaldoInsuficienteException(self.saldo, valor)
        self.saldo -= valor
        return self.saldo

    def depositar(self, valor):
        self.saldo += valor
        return self.saldo

    def get_saldo(self):
        return self.saldo

try:
    conta = ContaBancaria(100)
    conta.sacar(150)
except SaldoInsuficienteException as e:
    print(e)
