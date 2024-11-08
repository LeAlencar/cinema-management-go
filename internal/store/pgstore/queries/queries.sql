-- Movies
-- name: GetMovie :one
SELECT * FROM movies WHERE id = $1;

-- name: ListMovies :many
SELECT * FROM movies ORDER BY title;

-- name: CreateMovie :one
INSERT INTO movies (
    title,
    duration,
    genre,
    release_date
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateMovie :one
UPDATE movies
SET
    title = COALESCE($2, title),
    duration = COALESCE($3, duration),
    genre = COALESCE($4, genre),
    release_date = COALESCE($5, release_date),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = $1;

-- Rooms
-- name: GetRoom :one
SELECT * FROM rooms WHERE id = $1;

-- name: ListRooms :many
SELECT * FROM rooms ORDER BY number;

-- name: CreateRoom :one
INSERT INTO rooms (
    number,
    capacity,
    is_vip
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateRoom :one
UPDATE rooms
SET
    number = COALESCE($2, number),
    capacity = COALESCE($3, capacity),
    is_vip = COALESCE($4, is_vip),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms WHERE id = $1;

-- Sessions
-- name: GetSession :one
SELECT * FROM sessions WHERE id = $1;

-- name: ListSessions :many
SELECT * FROM sessions ORDER BY start_time;

-- name: ListSessionsByMovie :many
SELECT * FROM sessions WHERE movie_id = $1 ORDER BY start_time;

-- name: ListSessionsByRoom :many
SELECT * FROM sessions WHERE room_id = $1 ORDER BY start_time;

-- name: CreateSession :one
INSERT INTO sessions (
    movie_id,
    room_id,
    start_time,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateSession :one
UPDATE sessions
SET
    movie_id = COALESCE($2, movie_id),
    room_id = COALESCE($3, room_id),
    start_time = COALESCE($4, start_time),
    price = COALESCE($5, price),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;

-- Customers
-- name: GetCustomer :one
SELECT * FROM customers WHERE id = $1;

-- name: GetCustomerByEmail :one
SELECT * FROM customers WHERE email = $1;

-- name: ListCustomers :many
SELECT * FROM customers ORDER BY name;

-- name: CreateCustomer :one
INSERT INTO customers (
    name,
    email,
    phone
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateCustomer :one
UPDATE customers
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    phone = COALESCE($4, phone),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers WHERE id = $1;

-- Employees
-- name: GetEmployee :one
SELECT * FROM employees WHERE id = $1;

-- name: GetEmployeeByEmail :one
SELECT * FROM employees WHERE email = $1;

-- name: ListEmployees :many
SELECT * FROM employees ORDER BY name;

-- name: CreateEmployee :one
INSERT INTO employees (
    name,
    email,
    role,
    phone
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateEmployee :one
UPDATE employees
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email),
    role = COALESCE($4, role),
    phone = COALESCE($5, phone),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employees WHERE id = $1;

-- Products
-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products ORDER BY name;

-- name: ListProductsByCategory :many
SELECT * FROM products WHERE category = $1 ORDER BY name;

-- name: CreateProduct :one
INSERT INTO products (
    name,
    price,
    stock_count,
    category
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE($2, name),
    price = COALESCE($3, price),
    stock_count = COALESCE($4, stock_count),
    category = COALESCE($5, category),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateProductStock :one
UPDATE products
SET
    stock_count = stock_count + $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;

-- Orders
-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1;

-- name: ListOrders :many
SELECT * FROM orders ORDER BY created_at DESC;

-- name: ListOrdersByCustomer :many
SELECT * FROM orders WHERE customer_id = $1 ORDER BY created_at DESC;

-- name: CreateOrder :one
INSERT INTO orders (
    customer_id,
    total_amount,
    status
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateOrder :one
UPDATE orders
SET
    total_amount = COALESCE($2, total_amount),
    status = COALESCE($3, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;

-- Order Items
-- name: GetOrderItem :one
SELECT * FROM order_items WHERE id = $1;

-- name: ListOrderItems :many
SELECT * FROM order_items WHERE order_id = $1;

-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateOrderItem :one
UPDATE order_items
SET
    quantity = COALESCE($2, quantity),
    price = COALESCE($3, price),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteOrderItem :exec
DELETE FROM order_items WHERE id = $1;

-- Tickets
-- name: GetTicket :one
SELECT * FROM tickets WHERE id = $1;

-- name: ListTickets :many
SELECT * FROM tickets ORDER BY created_at DESC;

-- name: ListTicketsBySession :many
SELECT * FROM tickets WHERE session_id = $1;

-- name: ListTicketsByCustomer :many
SELECT * FROM tickets WHERE customer_id = $1 ORDER BY created_at DESC;

-- name: CreateTicket :one
INSERT INTO tickets (
    session_id,
    customer_id,
    seat_number,
    price
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets WHERE id = $1;

-- Complex Queries
-- name: GetSessionWithDetails :one
SELECT 
    s.*,
    m.title as movie_title,
    m.duration as movie_duration,
    r.number as room_number,
    r.capacity as room_capacity,
    r.is_vip as room_is_vip
FROM sessions s
JOIN movies m ON s.movie_id = m.id
JOIN rooms r ON s.room_id = r.id
WHERE s.id = $1;

-- name: GetOrderWithItems :one
SELECT 
    o.*,
    c.name as customer_name,
    c.email as customer_email,
    json_agg(
        json_build_object(
            'id', oi.id,
            'product_id', oi.product_id,
            'product_name', p.name,
            'quantity', oi.quantity,
            'price', oi.price
        )
    ) as items
FROM orders o
JOIN customers c ON o.customer_id = c.id
LEFT JOIN order_items oi ON o.id = oi.order_id
LEFT JOIN products p ON oi.product_id = p.id
WHERE o.id = $1
GROUP BY o.id, c.name, c.email;

-- name: GetAvailableSeats :many
SELECT 
    generate_series(1, r.capacity) as seat_number
FROM sessions s
JOIN rooms r ON s.room_id = r.id
WHERE s.id = $1
EXCEPT
SELECT CAST(seat_number AS integer)
FROM tickets
WHERE session_id = $1
ORDER BY seat_number;

