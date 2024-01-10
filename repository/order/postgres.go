package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/marceljaworski/go-microservice/model"
)

type PostgresRepo struct {
	db *sql.DB
}

func (r PostgresRepo) Insert(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}
	fmt.Println("data----->", data)
	key := orderIDKey(order.OrderID)

	_, err = r.db.Exec("INSERT INTO  orders(order_id,customer_id,created_at) VALUES($1,$2,$3)", key, order.CustomerID, order.CreatedAt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value inserted")
	}

	return nil
}

func (r *PostgresRepo) FindByID(ctx context.Context, id uint64) error {

	return nil
}

func (r *PostgresRepo) DeleteByID(ctx context.Context, id uint64) error {

	return nil
}

func (r *PostgresRepo) Update(ctx context.Context, order model.Order) error {

	return nil
}

func (r PostgresRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {

	fmt.Println("Hello, Docker!")

	return FindResult{}, nil
}
