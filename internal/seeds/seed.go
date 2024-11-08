package seeds

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func Run(ctx context.Context, pool *pgxpool.Pool) error {
	fake := faker.New()
	queries := pgstore.New(pool)

	// Limpa o banco
	if err := cleanDatabase(ctx, pool); err != nil {
		return fmt.Errorf("error cleaning database: %w", err)
	}

	// Cria movies
	movies, err := seedMovies(ctx, queries, fake)
	if err != nil {
		return fmt.Errorf("error seeding movies: %w", err)
	}
	fmt.Printf("Created %d movies\n", len(movies))
	for _, movie := range movies {
		fmt.Printf("- Movie: %s (Genre: %s, Duration: %d min)\n",
			movie.Title, movie.Genre, movie.Duration)
	}
	// Cria rooms
	rooms, err := seedRooms(ctx, queries, fake)
	if err != nil {
		return fmt.Errorf("error seeding rooms: %w", err)
	}

	fmt.Printf("\nCreated %d rooms\n", len(rooms))
	for _, room := range rooms {
		fmt.Printf("- Room %d (Capacity: %d, VIP: %v)\n",
			room.Number, room.Capacity, room.IsVip.Bool)
	}
	// // Cria sessions
	// sessions, err := seedSessions(ctx, queries, fake, movies, rooms)
	// if err != nil {
	// 	return fmt.Errorf("error seeding sessions: %w", err)
	// }
	//
	// // Cria customers
	// customers, err := seedCustomers(ctx, queries, fake)
	// if err != nil {
	// 	return fmt.Errorf("error seeding customers: %w", err)
	// }
	//
	// // Cria employees
	// if err := seedEmployees(ctx, queries, fake); err != nil {
	// 	return fmt.Errorf("error seeding employees: %w", err)
	// }
	//
	// // Cria products
	// products, err := seedProducts(ctx, queries, fake)
	// if err != nil {
	// 	return fmt.Errorf("error seeding products: %w", err)
	// }
	//
	// // Cria orders e order items
	// if err := seedOrders(ctx, queries, fake, customers, products); err != nil {
	// 	return fmt.Errorf("error seeding orders: %w", err)
	// }
	//
	// // Cria tickets
	// if err := seedTickets(ctx, queries, fake, sessions, customers); err != nil {
	// 	return fmt.Errorf("error seeding tickets: %w", err)
	// }

	fmt.Println("Seeding finished successfully!")

	return nil
}

func cleanDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	tables := []string{
		"tickets",
		"order_items",
		"orders",
		"products",
		"employees",
		"sessions",
		"rooms",
		"movies",
		"customers",
	}

	for _, table := range tables {
		if _, err := pool.Exec(ctx, fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)); err != nil {
			return err
		}
	}

	return nil
}
