package queries

import (
	"fmt"

	"github.com/angusbean/enviro-check/app/models"
	"github.com/jmoiron/sqlx"
)

type CityQueries struct {
	*sqlx.DB
}

//ClearCitiesTable removes all rows from cities table is DB
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

//InsertCityInfo add the city ID, Lat and Long from JSON file into DB
func (q *CityQueries) InsertCityInfo(m models.Cities) error {

	for _, c := range m.Cities {
		query := `INSERT INTO cities VALUES ($1, $2, $3)`

		// Send query to database
		_, err := q.Exec(
			query,
			c.ID, c.Lat, c.Long,
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("Insert Complete of Cities in db")

	return nil

}

//LocateCity checks the DB for the closest city ID based on the lat and long provided
func (q *CityQueries) LocateCity(lat float64, long float64) (int, error) {
	var closestCityID int

	query := `SELECT id FROM cities ORDER BY ABS((lat - $1) + (long - $2)) LIMIT 1`

	// Send query to database
	err := q.QueryRow(query, lat, long).Scan(&closestCityID)
	if err != nil {
		return closestCityID, err
	}

	return closestCityID, nil
}
