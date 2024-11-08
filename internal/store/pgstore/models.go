// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package pgstore

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Customer struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Phone     pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Employee struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Role      string
	Phone     pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Movie struct {
	ID          uuid.UUID
	Title       string
	Duration    int32
	Genre       string
	ReleaseDate pgtype.Timestamp
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type Order struct {
	ID          uuid.UUID
	CustomerID  uuid.UUID
	TotalAmount pgtype.Numeric
	Status      string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

type OrderItem struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	ProductID uuid.UUID
	Quantity  int32
	Price     pgtype.Numeric
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Product struct {
	ID         uuid.UUID
	Name       string
	Price      pgtype.Numeric
	StockCount int32
	Category   string
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

type Room struct {
	ID        uuid.UUID
	Number    int32
	Capacity  int32
	IsVip     pgtype.Bool
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Session struct {
	ID        uuid.UUID
	MovieID   uuid.UUID
	RoomID    uuid.UUID
	StartTime pgtype.Timestamp
	Price     pgtype.Numeric
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type Ticket struct {
	ID         uuid.UUID
	SessionID  uuid.UUID
	CustomerID uuid.UUID
	SeatNumber string
	Price      pgtype.Numeric
	CreatedAt  pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}