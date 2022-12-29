package model

type OrderStatus string

type Order struct {
	OrderID       string         `gorm:"primaryKey" json:"order_id"`
	UserID        string         `gorm:"column:user_id" json:"user_id"`
	Status        OrderStatus    `gorm:"column:status" json:"status"`
	ProductOrders []ProductOrder `json:"product_orders"`
	ReferenceID   string         `gorm:"unique" json:"reference_id"`
}

type ProductOrderStatus string

type ProductOrder struct {
	ID         string             `gorm:"primaryKey;" json:"id"`
	OrderID    string             `gorm:"column:order_id" json:"order_id"`
	OrderCode  string             `gorm:"column:order_code" json:"order_code"`
	Qty        int                `gorm:"column:qty" json:"qty"`
	TotalPrice int64              `gorm:"column:total_price" json:"total_price"`
	Status     ProductOrderStatus `gorm:"column:status" json:"status"`
}

type MenuProductOrderRequest struct {
	OrderCode string `json:"order_code"`
	Qty       int    `json:"qty"`
}

type MenuOrderRequest struct {
	UserID        string                    `json:"-"`
	OrderProducts []MenuProductOrderRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	UserID  string `json:"-"`
	OrderID string `json:"order_id"`
}
