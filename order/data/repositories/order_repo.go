package repositories

import (
	"database/sql"
	"demo/oms/order/data/models"
	"demo/oms/order/domain/repositories"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewOrderRepository(db *sqlx.DB) repositories.OrderRepository {
	return &orderRepo{db: db}
}

type orderRepo struct {
	db *sqlx.DB
}

// GetByID fetches order with the provided id.
func (ur *orderRepo) GetByID(id int64) (*models.Persons, error) {
	var order models.Persons
	err := ur.db.QueryRowx(queryGetByID, id).StructScan(&order)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows returned.")
		}
		return nil, err
	}

	return &order, nil
}
