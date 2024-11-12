package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedEmployees(ctx context.Context, q *pgstore.Queries, fake faker.Faker) ([]pgstore.Employee, error) {
	roles := []string{"Manager", "Cashier", "Usher", "Projectionist", "Cleaner", "Concession Worker"}
	var employees []pgstore.Employee

	for _, role := range roles {
		firstName := fake.Person().FirstName()
		lastName := fake.Person().LastName()
		email := fake.Internet().Email()

		employee, err := q.CreateEmployee(ctx, pgstore.CreateEmployeeParams{
			Name:  firstName + " " + lastName,
			Email: email,
			Role:  role,
			Phone: pgtype.Text{
				String: fake.Phone().Number(),
				Valid:  true,
			},
		})
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	for i := 0; i < 14; i++ {
		firstName := fake.Person().FirstName()
		lastName := fake.Person().LastName()
		email := fake.Internet().Email()
		phoneNumber := fmt.Sprintf("+%d-%d-%d",
			fake.IntBetween(1, 99),        // código do país
			fake.IntBetween(100, 999),     // código de área
			fake.IntBetween(10000, 99999), // número
		)

		employee, err := q.CreateEmployee(ctx, pgstore.CreateEmployeeParams{
			Name:  firstName + " " + lastName,
			Email: email,
			Role:  roles[fake.IntBetween(0, len(roles)-1)],
			Phone: pgtype.Text{String: phoneNumber, Valid: true},
		})
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
