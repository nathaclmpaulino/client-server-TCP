# client-server-TCP

Este repositório contém uma aplicação cliente servidor construída sobre protocolo TCP que tem como objetivo o servidor receber e processar uma string qualquer, convertendo todas as letras minúsculas para maiúsculas.

## Organização 

Este projeto está dividido de acordo com a seguinte estruturação:

1. client/
    1. client.go
2. server/
    1. server.go
3. README.md

## Para executar o código

Baixe o compilador de Go seguindo orientações do site da [Golang](https://golang.org/dl/) 
Assumindo que seu path atual seja a pasta base (~/client-server-TCP) 
Use os dois comandos:
1. go run server/server.go
2. go run client/client.go

A partir disso - digite qualquer palavra no cliente, e o próprio servidor processa a palavra e a redireciona para strings uppercase.