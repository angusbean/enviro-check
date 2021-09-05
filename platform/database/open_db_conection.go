package database

import "github.com/angusbean/enviro-check/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
	*queries.CityQueries // load queries from City model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
		CityQueries: &queries.CityQueries{DB: db}, // from City model
	}, nil
}
