package seeds

import (
	"context"
	"fmt"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedTickets(ctx context.Context, q *pgstore.Queries, fake faker.Faker, sessions []pgstore.Session, customers []pgstore.Customer) ([]pgstore.Ticket, error) {
	var tickets []pgstore.Ticket

	for _, session := range sessions {
		numTickets := fake.IntBetween(3, 10)
		usedSeats := make(map[string]bool)

		for i := 0; i < numTickets; i++ {
			var seatNumber string
			for {
				row := string(rune('A' + fake.IntBetween(0, 9)))
				number := fake.IntBetween(1, 20)
				seatNumber = fmt.Sprintf("%s%d", row, number)

				if !usedSeats[seatNumber] {
					usedSeats[seatNumber] = true
					break
				}
			}

			customer := customers[fake.IntBetween(0, len(customers)-1)]

			basePrice := session.Price.Int
			variation := big.NewInt(int64(fake.IntBetween(-200, 200)))
			ticketPrice := new(big.Int).Add(basePrice, variation)

			ticket, err := q.CreateTicket(ctx, pgstore.CreateTicketParams{
				SessionID:  session.ID,
				CustomerID: customer.ID,
				SeatNumber: seatNumber,
				Price: pgtype.Numeric{
					Int:   ticketPrice,
					Exp:   -2,
					Valid: true,
				},
			})
			if err != nil {
				return nil, err
			}
			tickets = append(tickets, ticket)
		}
	}

	return tickets, nil
}
