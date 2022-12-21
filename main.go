package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MenuItem struct {
	Name      string `json:"name"`
	OrderCode string `json:"order_code"`
	Price     int    `json:"price"`
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []MenuItem{
		{"Bakmie", "bakmie", 35000},
		{"Bakso", "bakso", 25000},
		{"Ayam Rica", "ayam_rica", 35000},
		{"Bebek Bakar", "bebek_bakar", 45000},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func getDrinksMenu(c echo.Context) error {
	drinksMenu := []MenuItem{
		{"Moccacino", "moccacino", 30000},
		{"Ice Choco", "choco_ice", 25000},
		{"Ice Strawberry", "strawberry_ice", 35000},
		{"Manggo Float", "manggo_float", 45000},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinksMenu,
	})
}

func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drinks", getDrinksMenu)
	e.Logger.Fatal(e.Start(":8000"))
}
