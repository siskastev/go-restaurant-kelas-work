package rest

import "go-restaurant-kelas-work/internal/usecase/resto"

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{restoUsecase}
}
