package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

const (
	dsn       = "host=localhost user=postgres password=1234 dbname=go_restaurant_app port=5432 sslmode=disable"
	typeFood  = "food"
	typeDrink = "drink"
)

type MenuType string

type MenuItem struct {
	OrderCode string   `json:"order_code"`
	Name      string   `json:"name"`
	Price     int      `json:"price"`
	Type      MenuType `json:"type"`
}

func main() {
	seeder()
	e := echo.New()
	e.GET("/menu", getMenu)
	e.Logger.Fatal(e.Start(":8000"))
}

func getMenu(c echo.Context) error {
	typeMenu := c.FormValue("type")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	var menu []MenuItem
	db.Where(MenuItem{Type: MenuType(typeMenu)}).Find(&menu)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menu,
	})
}

func seeder() {
	foodMenu := []MenuItem{
		{"bakmie", "Bakmie", 35000, typeFood},
		{"bakso", "bakso", 25000, typeFood},
		{"ayam_rica", "Ayam Rica", 35000, typeFood},
		{"bebek_bakar", "Bebek Bakar", 45000, typeFood},
	}

	drinksMenu := []MenuItem{
		{"moccacino", "Moccacino", 30000, typeDrink},
		{"choco_ice", "Choco Ice", 25000, typeDrink},
		{"strawberry_ice", "Strawberry Ice", 35000, typeDrink},
		{"manggo_float", "Manggo Float", 45000, typeDrink},
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&MenuItem{})
	if err := db.First(&MenuItem{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		//Insert seed data
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
