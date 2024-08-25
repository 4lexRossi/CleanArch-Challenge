package database

import (
	"database/sql"

	"github.com/4lexRossi/CleanArch-Challenge/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	// Prepare statement helps to insert various orders at the same time
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (c *OrderRepository) List() ([]entity.Order, error) {
	rows, err := c.Db.Query("SELECT id, price, finalPrice FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []entity.Order{}
	for rows.Next() {
		var id string
		var price, finalPrice float64
		if err := rows.Scan(&id, &price, &finalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, entity.Order{ID: id, Price: price, FinalPrice: finalPrice})
	}
	return orders, nil
}
