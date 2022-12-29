package database

import (
	"errors"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/model/constant"
	"gorm.io/gorm"
)

func seeder(db *gorm.DB) {
	//migrate schema
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{"bakmie", "Bakmie", 35000, constant.TypeFood},
		{"bakso", "bakso", 25000, constant.TypeFood},
		{"ayam_rica", "Ayam Rica", 35000, constant.TypeFood},
		{"bebek_bakar", "Bebek Bakar", 45000, constant.TypeFood},
	}

	drinksMenu := []model.MenuItem{
		{"moccacino", "Moccacino", 30000, constant.TypeDrink},
		{"choco_ice", "Choco Ice", 25000, constant.TypeDrink},
		{"strawberry_ice", "Strawberry Ice", 35000, constant.TypeDrink},
		{"manggo_float", "Manggo Float", 45000, constant.TypeDrink},
	}

	if err := db.First(&model.MenuItem{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//Insert seed data
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
