package seeds

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedRooms(ctx context.Context, q *pgstore.Queries, fake faker.Faker) ([]pgstore.Room, error) {
	var rooms []pgstore.Room

	for i := 1; i <= 8; i++ {
		room, err := q.CreateRoom(ctx, pgstore.CreateRoomParams{
			Number:   int32(i),
			Capacity: int32(fake.IntBetween(50, 200)),
			IsVip: pgtype.Bool{
				Bool:  i > 6, // Rooms 7 and 8 are VIP
				Valid: true,
			},
		})
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
