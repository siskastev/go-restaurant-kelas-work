package menu

import (
	"go-restaurant-kelas-work/internal/model"
	"golang.org/x/net/context"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockMenuRepository -destination=mocks/menu_repository_mock.go -source=repository.go
type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenuOrder(ctx context.Context, orderCode string) (model.MenuItem, error)
}
