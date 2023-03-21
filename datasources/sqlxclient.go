package datasources

import (
	"demo/oms/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// getSQLXClient connects to a mysql server and returns ...
func getSQLXClient() (*sqlx.DB, error) {
	inputConfig := config.Config.SQL

	dataSourceNames := getFormattedDataSourceName(inputConfig)

	db, err := sqlx.Open(inputConfig.DriverName, dataSourceNames)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to SQL!")

	return db, nil
}

func getFormattedDataSourceName(sqlConfig config.SQLConfig) string {
	dataSourceNameFormat := sqlConfig.DataSourceNameFormat

	if dataSourceNameFormat == "" {
		dataSourceNameFormat = "%s:%s@tcp(%s)/%s"
	}

	user := sqlConfig.ConnConfig.User
	password := sqlConfig.ConnConfig.Password
	host := sqlConfig.ConnConfig.Host
	db := sqlConfig.ConnConfig.DB

	dataSourceName := fmt.Sprintf(dataSourceNameFormat, user, password, host, db)

	return dataSourceName
}
