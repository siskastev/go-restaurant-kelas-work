package model

type MenuType string

type MenuItem struct {
	OrderCode string   `json:"order_code"`
	Name      string   `json:"name"`
	Price     int64    `json:"price"`
	Type      MenuType `json:"type"`
}
