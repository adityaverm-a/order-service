package repositories

import (
	"fmt"
)

var (
	queryGetByID = fmt.Sprintf("SELECT * FROM %s WHERE PersonID = ?", orderTable)
)
