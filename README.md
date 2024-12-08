# teste-biofy

## 📝 Descrição Geral
O projeto teste-biofy é uma API CRUD desenvolvida com foco no padrão de arquitetura RESTful. Seu objetivo principal é permitir as operações básicas de Criar, Listar, Atualizar e Deletar itens de um banco de dados relacional PostgreSQL.

## 🛠 Tecnologias Utilizadas
As tecnologias utilizadas para o desenvolvimento do projeto incluem:

Go: Linguagem de programação principal para o back-end.
HTML, CSS, JS: Tecnologias principais para o front-end.
PostgreSQL: Banco de dados relacional.
Mux: Biblioteca de roteamento para Go.
JWT (JSON Web Token): Para autenticação de usuários.
Docker: Para containerização e setup do ambiente.
Swagger: Ferramentas para documentação da API.

## 🚀 Setup e Execução
Siga os passos abaixo para executar o projeto localmente:

Pré-requisitos
- Go (versão 1.20 ou superior)
- Docker

### Passos para configurar

Clone o repositório:
```sh
git clone git@github.com:Rjoaozinho1/teste-biofy.git
cd teste-biofy
```

Suba o container do PostgreSQL dentro da pasta back/:
```sh
docker-compose up -d
```

Faça o build do projeto back-end e execute-o:
```sh
go build -o app
./app
```

Vá a seu navegador e digite essa url:

`http://localhost:2026`

Entre com esse usuario para teste:

`Email: 123@teste.com`
`Senha: 123`

## 📌 Exemplos de Uso dos Endpoints
Aqui estão alguns exemplos para testar os principais endpoints da API:

### Criar Item
Endpoint: POST /itens

Exemplo de cURL:
```sh
curl -X POST http://localhost:2026/itens \
-H "Content-Type: application/json" \
-d '{
  "nome": "Novo Item",
  "mensagem": "Este é um exemplo"
}'
```
### Atualizar Item
Endpoint: PUT /itens/{id}
```sh
Exemplo de cURL:
curl -X PUT http://localhost:2026/itens/uuid \
-H "Content-Type: application/json" \
-d '{
  "nome": "Item Atualizado",
  "mensagem": "Conteúdo atualizado"
}'
```
### Obter Item por ID
Endpoint: GET /itens/{id}
```sh
Exemplo de cURL:
curl -X GET http://localhost:2026/itens/uuid
```
### Deletar Item
Endpoint: DELETE /itens/{id}
```sh
Exemplo de cURL:
curl -X DELETE http://localhost:2026/itens/uuid
```
## 🛡 Documentação da API
A documentação detalhada da API pode ser encontrada nos seguintes formatos:
Swagger: Swagger UI
Para acessar a interface Swagger acesse:

`http://localhost:2026/doc.html`