package queries

import (
	"fmt"

	"github.com/angusbean/enviro-check/app/models"
	"github.com/jmoiron/sqlx"
)

type CityQueries struct {
	*sqlx.DB
}

//
func (q *CityQueries) ClearCitiesTable() error {
	// Define query string.
	query := `DELETE FROM cities`

	// Send query to database.
	_, err := q.Exec(query)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	fmt.Println("Deletion Complete of Cities in DB")
	return nil
}

func (q *CityQueries) InsertCityInfo(m models.Cities) error {
	for _, c := range m.Cities {
		query := `INSERT INTO cities VALUES ($1, $2, $3)`

		// Send query to database
		_, err := q.Exec(
			query,
			c.ID, c.Lat, c.Long,
		)
		fmt.Println(c.ID)
		if err != nil {
			return err
		}
	}
	fmt.Println("Insert Complete of Cities in db")
	return nil

}
