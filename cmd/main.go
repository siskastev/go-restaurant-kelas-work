package main

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/labstack/echo/v4"
	"go-restaurant-kelas-work/internal/database"
	"go-restaurant-kelas-work/internal/delivery/rest"
	"go-restaurant-kelas-work/internal/logger"
	mRepo "go-restaurant-kelas-work/internal/repository/menu"
	oRepo "go-restaurant-kelas-work/internal/repository/order"
	uRepo "go-restaurant-kelas-work/internal/repository/user"
	"go-restaurant-kelas-work/internal/tracing"
	rUsecase "go-restaurant-kelas-work/internal/usecase/resto"
	"time"
)

const (
	dsn = "host=localhost user=postgres password=1234 dbname=go_restaurant_app port=5432 sslmode=disable"
)

func main() {
	logger.Init()
	tracing.Init("http://localhost:14268/api/traces")
	e := echo.New()
	db := database.GetDB(dsn)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetOrderRepository(db)
	userRepo, err := uRepo.GetUserRepository(db, secret, 1, 64*1024, 4, 32, signKey, 5*time.Minute)
	if err != nil {
		panic(err)
	}
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)
	menuHandler := rest.NewHandler(restoUsecase)
	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, menuHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
