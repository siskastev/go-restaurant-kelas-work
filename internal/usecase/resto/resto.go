package resto

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go-restaurant-kelas-work/internal/model"
	"go-restaurant-kelas-work/internal/model/constant"
	"go-restaurant-kelas-work/internal/repository/menu"
	"go-restaurant-kelas-work/internal/repository/order"
	"go-restaurant-kelas-work/internal/repository/user"
	"go-restaurant-kelas-work/internal/tracing"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUsecase{menuRepo, orderRepo, userRepo}
}

func (r *restoUsecase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()
	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (r *restoUsecase) CreateOrder(ctx context.Context, request model.MenuOrderRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "CreateOrder")
	defer span.End()
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))
	for i, orderProduct := range request.OrderProducts {
		menu, err := r.menuRepo.GetMenuOrder(ctx, orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}
		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menu.OrderCode,
			Qty:        orderProduct.Qty,
			TotalPrice: menu.Price * int64(orderProduct.Qty),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}
	orderData := model.Order{
		OrderID:       uuid.New().String(),
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}
	createOrder, err := r.orderRepo.CreateOrder(ctx, orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createOrder, nil
}
func (r *restoUsecase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()
	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)
	if err != nil {
		return orderData, err
	}
	if orderData.UserID != request.UserID {
		return model.Order{}, errors.New("unauthorized")
	}
	return orderData, nil
}

func (r *restoUsecase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()
	userRegistered, err := r.userRepo.CheckRegister(ctx, request.Username)
	if err != nil {
		return model.User{}, err
	}
	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}
	userHash, err := r.userRepo.GenerateUserHash(ctx, request.Password)
	if err != nil {
		return model.User{}, err
	}
	userData, err := r.userRepo.RegisterUser(ctx, model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	})
	if err != nil {
		return model.User{}, err
	}
	return userData, nil

}

func (r *restoUsecase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()
	userData, err := r.userRepo.GetUserData(ctx, request.Username)
	if err != nil {
		return model.UserSession{}, err
	}
	verified, err := r.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)
	if err != nil {
		return model.UserSession{}, err
	}
	if !verified {
		return model.UserSession{}, errors.New("cant verified user login")
	}
	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)
	if err != nil {
		return model.UserSession{}, err
	}
	return userSession, nil
}

func (r *restoUsecase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", err
	}
	return userID, nil
}
