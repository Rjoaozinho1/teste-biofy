# teste-biofy

## Descrição Geral
O projeto teste-biofy é uma API CRUD desenvolvida com foco no padrão de arquitetura RESTful. Seu objetivo principal é permitir as operações básicas de Criar, Listar, Atualizar e Deletar mensagem de um banco de dados relacional PostgreSQL.

## Tecnologias Utilizadas
As tecnologias utilizadas para o desenvolvimento do projeto incluem:

Go: Linguagem de programação principal para o back-end.

HTML, CSS, JS: Tecnologias principais para o front-end.

PostgreSQL: Banco de dados relacional.

Mux: Biblioteca de roteamento para Go.

JWT (JSON Web Token): Para autenticação de usuários.

Docker: Para containerização e setup do ambiente.

Swagger: Ferramentas para documentação da API.

## Setup e Execução
Siga os passos abaixo para executar o projeto localmente:

Pré-requisitos
- Docker
- Docker Compose

Instrução para instalação do Docker:

### Windows and macOS

Docker Compose esta incluido
[Docker Desktop](https://www.docker.com/products/docker-desktop)
para Windows e macOS.

### Linux

Pode instalar os binarios do [Docker Compose](https://github.com/docker/compose)

Renomeando o binario para `docker-compose` e copiando o `$HOME/.docker/cli-plugins` 

Ou copie para um desses diretorios para instalar em todo o sistema:

* `/usr/local/lib/docker/cli-plugins` OU `/usr/local/libexec/docker/cli-plugins`
* `/usr/lib/docker/cli-plugins` OU `/usr/libexec/docker/cli-plugins`

Talvez seja necessario fazer o binario executavel com `chmod +x`

### Passos para configurar

Clone o repositório:
```sh
git clone git@github.com:Rjoaozinho1/teste-biofy.git
cd teste-biofy
```

Suba o container da aplicação:
```sh
docker-compose up --build
```

Vá a seu navegador e digite essa url:

`http://localhost:2026`

Entre com esse usuario para teste:

`Email: 123@teste.com`
`Senha: 123`

## Exemplos de Uso dos Endpoints

### Criar Mensagem
Endpoint: POST /mensagem

Exemplo de cURL:
```sh
curl -X POST http://localhost:2026/mensagem \
-H "Content-Type: application/json" \
-H "Authorization: Bearer SEU_TOKEN_AQUI" \
-d '{
  "nome": "Novo Mensagem",
  "mensagem": "Este é um exemplo"
}'
```
### Atualizar Mensagem
Endpoint: PUT /mensagem/{id}

Exemplo de cURL:
```sh
curl -X PUT http://localhost:2026/mensagem/uuid \
-H "Content-Type: application/json" \
-H "Authorization: Bearer SEU_TOKEN_AQUI" \
-d '{
  "nome": "Mensagem Atualizado",
  "mensagem": "Conteúdo atualizado"
}'
```
### Obter Mensagem por ID
Endpoint: GET /mensagem/{id}

Exemplo de cURL:
```sh
curl -X GET http://localhost:2026/mensagem/uuid
-H "Authorization: Bearer SEU_TOKEN_AQUI" \
```
### Deletar Mensagem
Endpoint: DELETE /mensagem/{id}

Exemplo de cURL:
```sh
curl -X DELETE http://localhost:2026/mensagem/uuid
-H "Authorization: Bearer SEU_TOKEN_AQUI" \
```
## Documentação da API
A documentação detalhada da API pode ser encontrada nos seguintes formatos:
Swagger: Swagger UI
Para acessar a interface Swagger acesse:

`http://localhost:2026/doc.html`
