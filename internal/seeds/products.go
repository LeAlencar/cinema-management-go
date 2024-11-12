package seeds

import (
	"context"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedProducts(ctx context.Context, q *pgstore.Queries, fake faker.Faker) ([]pgstore.Product, error) {
	products := []struct {
		name     string
		category string
		minPrice float64
		maxPrice float64
	}{
		{"Small Popcorn", "Snacks", 5.99, 7.99},
		{"Medium Popcorn", "Snacks", 7.99, 9.99},
		{"Large Popcorn", "Snacks", 9.99, 12.99},
		{"Small Soda", "Beverages", 3.99, 4.99},
		{"Medium Soda", "Beverages", 4.99, 5.99},
		{"Large Soda", "Beverages", 5.99, 6.99},
		{"Water Bottle", "Beverages", 2.99, 3.99},
		{"Nachos", "Snacks", 6.99, 8.99},
		{"Hot Dog", "Snacks", 5.99, 7.99},
		{"Candy Bar", "Candy", 2.99, 3.99},
		{"M&Ms", "Candy", 3.99, 4.99},
		{"Skittles", "Candy", 3.99, 4.99},
		{"Ice Cream", "Desserts", 4.99, 6.99},
		{"Cookie", "Desserts", 2.99, 3.99},
	}

	var createdProducts []pgstore.Product

	for _, p := range products {
		priceRange := p.maxPrice - p.minPrice
		randomPrice := p.minPrice + (float64(fake.IntBetween(0, 100)) / 100.0 * priceRange)

		priceInt := new(big.Int)
		priceInt.SetInt64(int64(randomPrice * 100))

		product, err := q.CreateProduct(ctx, pgstore.CreateProductParams{
			Name: p.name,
			Price: pgtype.Numeric{
				Int:   priceInt,
				Exp:   -2,
				Valid: true,
			},
			StockCount: int32(fake.IntBetween(50, 200)),
			Category:   p.category,
		})
		if err != nil {
			return nil, err
		}
		createdProducts = append(createdProducts, product)
	}

	return createdProducts, nil
}
