# Projeto de Banco de Dados

Este projeto tem como objetivo apresentar o conhecimento adquirido durante a disciplina de Banco de Dados. O projeto foi desenvolvido utilizando
Golang, PostgreSQL e Docker.

Integrantes: 

- Leandro de Brito Alencar RA: 22.222.034-5

- Caio Arnoni RA: 22.221.019-7

- Mateus Rocha RA: 22.222.002-2

## Como executar o projeto
Para rodar o projeto é necessário ter o Docker instalado em sua máquina. Assim como o Golang.

### Clonar o repositório
```bash
git clone https://github.com/lealencar/music-db-project.git
```

### Rodar o banco de dados
```bash
docker-compose up -d
```

### Criar o .env com o seguinte conteúdo
```bash
# DB_HOST=localhost
# DB_USER=admin
# DB_PASSWORD=admin@123
# DB_NAME=cinema
# DB_PORT=5432
```

### Rodar o generate do sqlc
```bash
go generate ./... && sqlc generate -f ./internal/store/pgstore/sqlc.yaml

```

### Rodar as migrations do banco de dados
```bash
go run cmd/tools/terndotenv/main.go
```

### Rodar o seed do banco de dados
```bash
go run cmd/tools/seed/main.go
```

### Rodar o projeto
```bash
go run cmd/server/main.go
```

### Verificar se o projeto está rodando
```bash
curl localhost:8080/ping
// {"message":"pong","timestamp":"2024-11-18T23:01:47.679964-03:00"}%
```

