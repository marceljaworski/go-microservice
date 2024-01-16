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
	fmt.Println("ctx----->", ctx)
	key := orderIDKey(order.OrderID)

	_, err = r.db.Exec("INSERT INTO orders_db (order_id, customer_id, product_no, created_at) VALUES ($1,$2,$3,4$)", key, order.CustomerID, order.Products, order.CreatedAt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value inserted")
	}

	return nil
}

func (r *PostgresRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {
	err := ErrNotExist
	return model.Order{}, fmt.Errorf("get order: %w", err)
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
