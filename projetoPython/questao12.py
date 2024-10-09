#Sobrecarga de Operadores (Python) / Métodos de Conveniência (Java, Go) Em Python,
#sobrecarregue o operador + para somar dois objetos Produto baseados no preço. Em Java
#e Golang, crie métodos que permitam essa funcionalidade.

class Produto:
    def __init__(self, nome, preco):
        self.nome = nome
        self.preco = preco

    def __add__(self, outro):
        if isinstance(outro, Produto):
            novo_nome = f"{self.nome} e {outro.nome}"
            novo_preco = self.preco + outro.preco
            return Produto(novo_nome, novo_preco)
        return NotImplemented

    def __str__(self):
        return f'{self.nome}: R$ {self.preco:.2f}'

produto1 = Produto("Tapioca", 50.00)
produto2 = Produto("Goiaba", 30.00)

produto3 = produto1 + produto2

print(produto1)
print(produto2)
print(produto3)
