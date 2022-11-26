package database

import "adomaitisc.com/main/app/models"

func (d* DB) CreateMovie(m *models.Movie) error {
	res, err := d.db.Exec(insertMovieSchema, m.MovieID, m.MovieTitle, m.MovieYear, m.MovieLink, m.MovieRating, m.MovieUserRate, m.MovieWatched, m.MovieDirector, m.MoviePoster)
	if err != nil {
		return err
	} // end if

	res.LastInsertId()
	return err
} // end CreateMovie func

func (d* DB) GetMovie() ([]*models.Movie, error) {
	rows, err := d.db.Query(getMovieSchema)
	if err != nil {
		return nil, err
	} // end if

	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		movie := &models.Movie{}

		err = rows.Scan(&movie.MovieID, &movie.MovieTitle, &movie.MovieYear, &movie.MovieLink, &movie.MovieRating, &movie.MovieUserRate, &movie.MovieWatched, &movie.MovieDirector, &movie.MoviePoster)
		if err != nil {
			return nil, err
		} // end if

		movies = append(movies, movie)
	} // end for

	return movies, err
} // end GetMovie func