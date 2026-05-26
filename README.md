# CRUD User

API RESTful em Go para gerenciar usuários com persistência em PostgreSQL.

## Tecnologias

- [Go](https://golang.org/dl/) 1.21 ou superior
- [PostgreSQL](https://www.postgresql.org/) 15
- [pgx](https://github.com/jackc/pgx) — driver PostgreSQL
- [sqlc](https://sqlc.dev/) — geração de queries type-safe
- [Goose](https://github.com/pressly/goose) — migrations
- [chi](https://github.com/go-chi/chi) — roteador HTTP
- [Docker](https://www.docker.com/) — banco de dados local

## Requisitos

- Go 1.21 ou superior
- Docker e Docker Compose

## Instalação

Clone o repositório e instale as dependências:

```bash
git clone https://github.com/Ericles-Miller/crudUserInMemory.git
cd crudUserInMemory
go mod tidy
```

## Configuração

Crie um arquivo `.env` na raiz do projeto com base no `.env.example`:

```bash
cp .env.example .env
```

Preencha a variável:

```env
DATABASE_URL=postgres://myuser:mysecretpassword@localhost:5432/mvc_database
```

## Executando o projeto

Suba o banco de dados:

```bash
docker compose up -d
```

Execute a aplicação (as migrations rodam automaticamente no startup):

```bash
go run main.go
```

Ou com hot reload:

```bash
air
```

O servidor estará disponível em `http://localhost:8080`.

## Build

```bash
go build -o server .
./server
```

## Endpoints

| Método | Rota | Descrição |
|---|---|---|
| POST | `/api/users` | Criar usuário |
| GET | `/api/users` | Listar todos os usuários |
| GET | `/api/users/{id}` | Buscar usuário por ID |
| PUT | `/api/users/{id}` | Atualizar usuário |
| DELETE | `/api/users/{id}` | Deletar usuário |

## Exemplo de payload

```json
{
  "firstName": "Jane",
  "lastName": "Doe",
  "biography": "Desenvolvedora apaixonada por tecnologia e open source."
}
```

## Regras de validação

| Campo | Tipo | Regras |
|---|---|---|
| `firstName` | string | obrigatório, mín. 2, máx. 20 caracteres |
| `lastName` | string | obrigatório, mín. 2, máx. 20 caracteres |
| `biography` | string | obrigatório, mín. 20, máx. 450 caracteres |

## Estrutura do projeto

```
.
├── main.go
├── docker-compose.yml
├── sqlc.yaml
├── api/
│   ├── api.go              # rotas
│   ├── controller.go       # handlers HTTP
│   ├── response.go         # struct de resposta
│   └── sendJson.go         # utilitário de resposta JSON
├── application/
│   ├── aplication.go       # serviço — regras de negócio e CRUD
│   ├── userBody.go         # struct de entrada
│   └── validateUser.go     # validações
└── internal/
    └── database/
        ├── connection.go   # conexão com o banco (pgxpool)
        ├── migrations.go   # execução automática de migrations
        ├── migrations/     # arquivos SQL de migration (Goose)
        ├── queries/        # queries SQL (sqlc)
        └── pgstore/        # código gerado pelo sqlc
```
