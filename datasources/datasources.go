package datasources

import "github.com/jmoiron/sqlx"

var dataSources DataSources

// DataSources is a struct containing an SQLX client.
type DataSources struct {
	SQLXClient *sqlx.DB
}

// Get returns a pointer to a DataSources struct and an error.
// The function retrieves an SQLX client and returns it back to the calling function.
func Get() (*DataSources, error) {

	// getSQLXClient() retrieves an SQLX client.
	SQLXClient, err := getSQLXClient()
	if err != nil {
		return &dataSources, err
	}

	dataSources.SQLXClient = SQLXClient

	return &dataSources, nil
}
