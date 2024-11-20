# Projeto de Banco de Dados

Este projeto tem como objetivo apresentar o conhecimento adquirido durante a disciplina de Banco de Dados. O projeto foi desenvolvido utilizando
Golang, PostgreSQL e Docker.

Integrantes: 

- Leandro de Brito Alencar RA: 22.222.034-5

- Caio Arnoni RA: 22.221.019-7

- Mateus Rocha RA: 22.222.002-2

---

# Sistema de Gerenciamento de Cinema

## Requisitos do Sistema

1. **Cliente (Customer)**:
   - Cada cliente possui um identificador único, nome, e-mail (que deve ser único), telefone e datas de criação e atualização.
   - Clientes podem comprar ingressos ou realizar pedidos.

2. **Funcionário (Employee)**:
   - Cada funcionário possui um identificador único, nome, e-mail (que deve ser único), função (cargo), telefone e datas de criação e atualização.

3. **Filme (Movie)**:
   - Cada filme possui um título, duração em minutos, gênero, data de lançamento e datas de criação e atualização.

4. **Sala (Room)**:
   - Cada sala possui um número identificador, capacidade de assentos, uma flag que indica se é VIP, e datas de criação e atualização.

5. **Sessão (Session)**:
   - Cada sessão está vinculada a um filme e uma sala.
   - Possui um horário de início, preço, e datas de criação e atualização.

6. **Ingresso (Ticket)**:
   - Cada ingresso está relacionado a uma sessão e um cliente.
   - Possui um número de assento, preço, e datas de criação e atualização.

7. **Pedido (Order)**:
   - Cada pedido está vinculado a um cliente.
   - Contém o valor total, status do pedido, e datas de criação e atualização.

8. **Itens do Pedido (OrderItem)**:
   - Cada item de pedido está associado a um pedido e um produto.
   - Possui quantidade, preço e datas de criação e atualização.

9. **Produto (Product)**:
   - Cada produto possui um nome, preço, estoque, categoria, e datas de criação e atualização.

---

## Modelagem do Banco de Dados

A modelagem do banco foi realizada utilizando um diagrama Entidade-Relacionamento (ER) e está normalizada até a Terceira Forma Normal (3FN). A estrutura de relacionamentos é a seguinte:

1. **Clientes (Customer)** podem realizar **pedidos (Order)** e comprar **ingressos (Ticket)**.
2. **Funcionários (Employee)** são registrados com informações de contato e função.
3. **Filmes (Movie)** são exibidos em **salas (Room)** através de **sessões (Session)**.
4. **Ingressos (Ticket)** estão vinculados a **sessões (Session)** e **clientes (Customer)**.
5. **Pedidos (Order)** podem conter múltiplos **produtos (Product)** por meio de **itens do pedido (OrderItem)**.
6. **Produtos (Product)** têm preço, estoque e categorias definidas.

---

## Diagrama Entidade-Relacionamento (ER)

O diagrama foi criado seguindo os relacionamentos descritos acima. Veja como as entidades estão relacionadas:

- **Customer** está relacionado com:
  - **Order**: Um cliente pode ter muitos pedidos.
  - **Ticket**: Um cliente pode comprar vários ingressos.

- **Movie** está relacionado com:
  - **Session**: Um filme pode ter várias sessões.

- **Room** está relacionado com:
  - **Session**: Uma sala pode exibir várias sessões.

- **Session** está relacionado com:
  - **Ticket**: Uma sessão pode gerar muitos ingressos.

- **Order** está relacionado com:
  - **OrderItem**: Um pedido pode conter vários itens.
  - **Customer**: Um pedido pertence a um cliente.

- **OrderItem** está relacionado com:
  - **Product**: Um item de pedido se refere a um produto específico.

---
## Questões

1. **Quais clientes realizaram pedidos com valor total superior a 100?**
```sql
SELECT customers.name, customers.email, orders.total_amount
FROM customers
JOIN orders ON customers.id = orders.customer_id
WHERE orders.total_amount > 100;
```

2. **Quais filmes estão atualmente sendo exibidos em sessões?**
```sql
SELECT movies.title, movies.genre, sessions.start_time
FROM movies
JOIN sessions ON movies.id = sessions.movie_id
WHERE sessions.start_time > CURRENT_TIMESTAMP
```

3. **Qual é a capacidade total das salas VIP?**
```sql
SELECT SUM(rooms.Capacity) AS TotalCapacity
FROM rooms
WHERE rooms.is_vip = TRUE;
```

4. **Quais produtos estão com estoque inferior a 10 unidades?**
```sql
SELECT products.name, products.stock_count
FROM products
WHERE products.stock_count > 100;
```

5. **Quais são os 5 filmes mais longos cadastrados no sistema?**
```sql
SELECT movies.title, movies.duration
FROM movies
ORDER BY movies.duration DESC
LIMIT 5;
```

6. **Quais clientes compraram ingressos para sessões em uma sala específica?**
```sql
SELECT customers.name, customers.email, rooms.number AS RoomNumber
FROM customers
JOIN tickets ON customers.id = tickets.customer_id
JOIN sessions ON tickets.session_id = sessions.id
JOIN rooms ON sessions.room_id = rooms.id
WHERE rooms.number = 1;
```

7. **Quais funcionários têm o cargo de 'Gerente'?**
```sql
SELECT employees.name, employees.email, employees.role
FROM employees
WHERE employees.role = 'Manager';
```

8. **Quais sessões estão programadas para começar nas próximas 24 horas?**
```sql
SELECT movies.title, rooms.number as room_number, sessions.start_time
FROM sessions
JOIN movies ON sessions.movie_id = movies.id
JOIN rooms ON sessions.room_id = rooms.id
WHERE sessions.start_time BETWEEN CURRENT_TIMESTAMP AND CURRENT_TIMESTAMP + INTERVAL '1 DAY';
```

9. **Quais pedidos contêm produtos da categoria 'Alimentos'?**
```sql
SELECT orders.id AS order_id, products.name, products.category
FROM orders
JOIN order_items ON orders.id = order_items.order_id
JOIN products ON order_items.product_id = products.id
WHERE products.category = 'Snacks';
```

10. **Quantos ingressos foram vendidos para cada sessão?**
```sql
SELECT Session.id AS SessionID, Movie.title, COUNT(Ticket.id) AS TicketCount
FROM sessions as Session
JOIN movies as Movie ON Session.movie_id = Movie.id
JOIN tickets as Ticket ON Session.id = Ticket.session_id
GROUP BY Session.id, Movie.title;
```
## Diagrama Relacional

![image](https://github.com/user-attachments/assets/4c7980e6-dafe-42a1-bcfa-e81f994b745f)

---

## Como executar o projeto
Para rodar o projeto é necessário ter o Docker instalado em sua máquina. Assim como o Golang.

### Clonar o repositório
```bash
git clone https://github.com/lealencar/cinema-management-go.git
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

