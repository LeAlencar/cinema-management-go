package seeds

import (
	"context"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedOrders(ctx context.Context, q *pgstore.Queries, fake faker.Faker, customers []pgstore.Customer, products []pgstore.Product) ([]pgstore.Order, error) {
	orderStatuses := []string{"Completed", "Processing", "Cancelled"}
	var orders []pgstore.Order

	for i := 0; i < 100; i++ {
		customer := customers[fake.IntBetween(0, len(customers)-1)]

		order, err := q.CreateOrder(ctx, pgstore.CreateOrderParams{
			CustomerID: customer.ID,
			TotalAmount: pgtype.Numeric{
				Int:   big.NewInt(0),
				Exp:   -2,
				Valid: true,
			},
			Status: orderStatuses[fake.IntBetween(0, len(orderStatuses)-1)],
		})
		if err != nil {
			return nil, err
		}

		totalAmount := big.NewInt(0)
		numItems := fake.IntBetween(1, 5)

		availableProducts := make([]pgstore.Product, len(products))
		copy(availableProducts, products)
		for j := len(availableProducts) - 1; j > 0; j-- {
			k := fake.IntBetween(0, j)
			availableProducts[j], availableProducts[k] = availableProducts[k], availableProducts[j]
		}

		for j := 0; j < numItems && j < len(availableProducts); j++ {
			product := availableProducts[j]
			quantity := fake.IntBetween(1, 3)

			itemPrice := new(big.Int).Mul(product.Price.Int, big.NewInt(int64(quantity)))

			_, err := q.CreateOrderItem(ctx, pgstore.CreateOrderItemParams{
				OrderID:   order.ID,
				ProductID: product.ID,
				Quantity:  int32(quantity),
				Price: pgtype.Numeric{
					Int:   itemPrice,
					Exp:   -2,
					Valid: true,
				},
			})
			if err != nil {
				return nil, err
			}

			totalAmount = new(big.Int).Add(totalAmount, itemPrice)
		}

		order, err = q.UpdateOrder(ctx, pgstore.UpdateOrderParams{
			ID: order.ID,
			TotalAmount: pgtype.Numeric{
				Int:   totalAmount,
				Exp:   -2,
				Valid: true,
			},
			Status: order.Status,
		})
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
