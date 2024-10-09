#Implemente o padrão de projeto Singleton para garantir que apenas uma
#instância de uma classe Configuracao seja criada.


class Configuracao:
    _instancia = None

    def __new__(cls, *args, **kwargs):
        if cls._instancia is None:
            cls._instancia = super(Configuracao, cls).__new__(cls)
            cls._instancia.inicializar()
        return cls._instancia

    def inicializar(self):
        self.configuracoes = {}

    def set_configuracao(self, chave, valor):
        self.configuracoes[chave] = valor

    def get_configuracao(self, chave):
        return self.configuracoes.get(chave, None)


config1 = Configuracao()
config1.set_configuracao('cor_tema', 'azul')

config2 = Configuracao()
config2.set_configuracao('font_size', '12px')

print(config1.get_configuracao('cor_tema'))
print(config2.get_configuracao('font_size'))

print(config1 is config2)
