package constant

import "go-restaurant-kelas-work/internal/model"

const (
	OrderStatusProcessed        model.OrderStatus        = "processed"
	OrderStatusFinished         model.OrderStatus        = "finished"
	OrderStatusFailed           model.OrderStatus        = "failed"
	ProductOrderStatusPreparing model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinished  model.ProductOrderStatus = "finished"
)
