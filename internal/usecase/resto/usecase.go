package resto

import (
	"context"
	"go-restaurant-kelas-work/internal/model"
)

// Usecase --
//
//go:generate mockgen -package=mocks -mock_names=Usecase=MockRestoUsecase -destination=mocks/resto_usecase_mock.go -source=usecase.go
type Usecase interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	CreateOrder(ctx context.Context, request model.MenuOrderRequest) (model.Order, error)
	GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error)
	RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error)
	Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
