package order

import (
	"errors"
	"fmt"

	"github.com/marceljaworski/go-microservice/model"
)

func orderIDKey(id uint64) string {
	return fmt.Sprintf("order:%d", id)
}

var ErrNotExist = errors.New("order does not exist")

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Orders []model.Order
	Cursor uint64
}
