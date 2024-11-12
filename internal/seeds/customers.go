package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedCustomers(ctx context.Context, q *pgstore.Queries, fake faker.Faker) ([]pgstore.Customer, error) {
	var customers []pgstore.Customer

	for i := 0; i < 50; i++ {
		firstName := fake.Person().FirstName()
		lastName := fake.Person().LastName()
		email := fake.Internet().Email()
		phoneNumber := fmt.Sprintf("+%d-%d-%d",
			fake.IntBetween(1, 99),        // código do país
			fake.IntBetween(100, 999),     // código de área
			fake.IntBetween(10000, 99999), // número
		)
		customer, err := q.CreateCustomer(ctx, pgstore.CreateCustomerParams{
			Name:  firstName + " " + lastName,
			Email: email,
			Phone: pgtype.Text{String: phoneNumber, Valid: true},
		})
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}
