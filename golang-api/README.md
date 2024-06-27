# Golang API

## Descrição
Como usar a API:
Ler os dados do CSV:

Endpoint: GET /read
Retorna todos os registros do CSV em formato JSON.
Adicionar um novo registro ao CSV:

Endpoint: POST /write
Body: JSON com os campos do registro a ser adicionado.
Adiciona o novo registro ao CSV e retorna o status 201 Created.
Atualizar um registro existente no CSV:

Endpoint: PUT /update
Body: JSON com os campos atualizados do registro. O campo PartnerId é usado para encontrar o registro a ser atualizado.
Atualiza o registro correspondente no CSV e retorna o status 200 OK.

## Execução
Para executar a aplicação, use os comandos:

```sh
go run cmd/main.go
