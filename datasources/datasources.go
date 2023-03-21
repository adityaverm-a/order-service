package datasources

import (
	"gorm.io/gorm"
)

var dataSources DataSources

// DataSources is a struct containing an GORM client.
type DataSources struct {
	GORMClient *gorm.DB
}

// Get returns a pointer to a DataSources struct and an error.
// The function retrieves an GORM client and returns it back to the calling function.
func Get() (*DataSources, error) {
	GORMClient, err := getGORMClient()
	if err != nil {
		return &dataSources, err
	}

	dataSources.GORMClient = GORMClient

	return &dataSources, nil
}

// GetDB helps you to get a connection
func GetSQLClient() *gorm.DB {
	return dataSources.GORMClient
}
