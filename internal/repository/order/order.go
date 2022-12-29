package order

import (
	"context"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/tracing"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func GetOrderRepository(db *gorm.DB) Repository {
	return &orderRepository{db}
}

func (o *orderRepository) CreateOrder(ctx context.Context, order model.Order) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()
	if err := o.db.WithContext(ctx).Create(order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (o *orderRepository) GetOrderInfo(ctx context.Context, orderID string) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()
	data := model.Order{}
	if err := o.db.WithContext(ctx).Where(model.Order{OrderID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
