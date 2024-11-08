package seeds

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jaswdr/faker"

	"cinema-project-go/internal/store/pgstore"
)

func seedMovies(ctx context.Context, q *pgstore.Queries, fake faker.Faker) ([]pgstore.Movie, error) {
	var movies []pgstore.Movie
	genres := []string{"Action", "Comedy", "Drama", "Horror", "Sci-Fi", "Adventure", "Animation"}
	releaseDate := fake.Time().TimeBetween(
		time.Now().AddDate(-1, 0, 0),
		time.Now().AddDate(1, 0, 0),
	)

	for i := 0; i < 10; i++ {
		movie, err := q.CreateMovie(ctx, pgstore.CreateMovieParams{
			Title:    fake.Music().Name(),
			Duration: int32(fake.IntBetween(90, 180)),
			Genre:    genres[fake.IntBetween(0, len(genres)-1)],
			ReleaseDate: pgtype.Timestamp{
				Time:  releaseDate,
				Valid: true,
			},
		})
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
