# CRUD User In Memory

API RESTful em Go para gerenciar usuários com armazenamento em memória.

## Requisitos

- [Go](https://golang.org/dl/) 1.21 ou superior

## Instalação

Clone o repositório e instale as dependências:

```bash
git clone https://github.com/Ericles-Miller/crudUserInMemory.git
cd crudUserInMemory
go mod tidy
```

## Executando o projeto

```bash
go run main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Build

Para gerar o binário:

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
├── api/
│   ├── api.go         # rotas
│   ├── controller.go  # handlers HTTP
│   ├── response.go    # struct de resposta
│   └── sendJson.go    # utilitário de resposta JSON
└── application/
    ├── aplication.go  # struct Application e métodos CRUD
    ├── User.go        # struct User
    ├── userBody.go    # struct UserBody
    └── validateUser.go # validações
```
