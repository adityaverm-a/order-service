package datasources

import (
	"demo/oms/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// getGORMClient connects to a mysql server and returns ...
func getGORMClient() (*gorm.DB, error) {
	inputConfig := config.Config.SQL

	dataSourceName := getFormattedDataSourceName(inputConfig)

	cfg := &gorm.Config{
		SkipDefaultTransaction: true,
	}
	db, err := gorm.Open(mysql.Open(dataSourceName), cfg)
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
