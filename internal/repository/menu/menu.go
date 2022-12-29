package menu

import (
	"context"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/tracing"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{db}
}

func (m *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()
	var menu []model.MenuItem
	if err := m.db.WithContext(ctx).Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (m *menuRepo) GetMenuOrder(ctx context.Context, orderCode string) (menu model.MenuItem, err error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuOrder")
	defer span.End()
	if err = m.db.WithContext(ctx).Where(model.MenuItem{OrderCode: orderCode}).First(&menu).Error; err != nil {
		return
	}
	return
}
