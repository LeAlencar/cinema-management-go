package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title       string
	Duration    int
	Genre       string
	ReleaseDate time.Time
	Sessions    []Session
}

type Room struct {
	gorm.Model
	Number   int
	Capacity int
	IsVIP    bool
	Sessions []Session
}

type Session struct {
	gorm.Model
	MovieID   uint
	RoomID    uint
	StartTime time.Time
	Price     float64
	Movie     Movie
	Room      Room
	Tickets   []Ticket
}

type Customer struct {
	gorm.Model
	Name    string
	Email   string
	Phone   string
	Tickets []Ticket
	Orders  []Order
}

type Ticket struct {
	gorm.Model
	SessionID  uint
	CustomerID uint
	SeatNumber string
	Price      float64
	Session    Session
	Customer   Customer
}

type Employee struct {
	gorm.Model
	Name  string
	Email string
	Role  string
	Phone string
}

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	StockCount int
	Category   string
	OrderItems []OrderItem
}

type Order struct {
	gorm.Model
	CustomerID  uint
	TotalAmount float64
	Status      string
	Customer    Customer
	OrderItems  []OrderItem
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
	Order     Order
	Product   Product
}
