package seeds

import (
	"context"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedSessions(ctx context.Context, q *pgstore.Queries, fake faker.Faker, movies []pgstore.Movie, rooms []pgstore.Room) ([]pgstore.Session, error) {
	var sessions []pgstore.Session

	for day := 0; day < 7; day++ {
		baseTime := time.Now().AddDate(0, 0, day)

		for _, room := range rooms {
			numSessions := fake.IntBetween(3, 5)
			for i := 0; i < numSessions; i++ {
				movie := movies[fake.IntBetween(0, len(movies)-1)]

				startHour := 10 + (i * 3)
				startTime := time.Date(
					baseTime.Year(),
					baseTime.Month(),
					baseTime.Day(),
					startHour,
					fake.IntBetween(0, 30),
					0, 0,
					time.Local,
				)

				basePrice := 25.0
				if room.IsVip.Bool {
					basePrice = 40.0
				}
				if startHour >= 18 {
					basePrice += 5.0
				}
				finalPrice := basePrice + float64(fake.IntBetween(-2, 5))

				priceInt := new(big.Int)
				priceInt.SetInt64(int64(finalPrice * 100))

				session, err := q.CreateSession(ctx, pgstore.CreateSessionParams{
					MovieID: movie.ID,
					RoomID:  room.ID,
					StartTime: pgtype.Timestamp{
						Time:  startTime,
						Valid: true,
					},
					Price: pgtype.Numeric{
						Int:   priceInt,
						Exp:   -2,
						Valid: true,
					},
				})
				if err != nil {
					return nil, err
				}
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}
