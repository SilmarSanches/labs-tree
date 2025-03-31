# labs-tree

Projeto de conclusão de pós-graduação

## Indice
1. [Build-Dev](#build-dev)
2. [DockerCompose](#dockercompose)
3. [Testes](#testes)

## Build-Dev

O app é acessível na porta 8080.

- na raiz do projeto rodar o comando
```bash
go run ./cmd/auction
```

## DockerCompose

O docker compose está na raiz do projeto, então rodar o comando abaixo para a execução:

```bash
docker compose up -d
````

## Testes

- na pasta api foram criados testes para as chamadas dos métodos. Adicionei um método para criar o usuário e retornar o ID dele para poder configurar as outras chamadas.

- no método de criação do leilão foi desenvolvido sem a resposta do leilão criado, não alterei, por tanto, deve se buscar o ID do leilão no mongodb.

- na sequencia temos um método que adiciona os lances.

- finalmente o método para consultar o lance vencedor.
