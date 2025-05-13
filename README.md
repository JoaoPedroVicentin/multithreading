# 🔀 Multithreading com Go - Consulta de CEP

Este projeto tem como objetivo a resolução do desafio prático **Multithreading** da **GoExpert** para obter informações de endereço com base em um CEP.

## 🚀 Desafio

Fazer duas requisições simultâneas para as seguintes APIs:

- [BrasilAPI](https://brasilapi.com.br/docs#tag/CEP)
- [ViaCEP](https://viacep.com.br/)

A aplicação deve:

- [x] Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.
- [x] Exibir o resultado da request no command line com os dados do endereço, bem como qual API a enviou.
- [x] Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## 💻 Como rodar o projeto

1. Clone o repositório:

```bash
git clone https://github.com/JoaoPedroVicentin/multithreading.git
cd multithreading
```

2. Execute o programa (certifique-se de ter o Go instalado):

```bash
go run main.go
```
---

<div align="center">
<h3>👨‍💻</h3>
    <h3> Criado por João Pedro Vicentin!</h3>
    <div>
        <h3>
            <a href="https://www.linkedin.com/in/joaopedrovicentin/" target="_blank">Linkedin</a>
            <a href='https://github.com/JoaoPedroVicentin' target='_blank'>Github</a>
            <a href="https://contate.me/joao-pedro-lopes-vicentin" target="_blank">Whatsapp</a>
        </h3>
    </div>
</div>