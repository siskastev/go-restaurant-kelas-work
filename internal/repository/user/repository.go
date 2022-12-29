package user

import (
	"go-restaurant-kelas-work/internal/model"
	"golang.org/x/net/context"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockUserRepository -destination=mocks/user_repository_mock.go -source=repository.go
type Repository interface {
	RegisterUser(ctx context.Context, user model.User) (model.User, error)
	CheckRegister(ctx context.Context, username string) (bool, error)
	GenerateUserHash(ctx context.Context, password string) (hash string, err error)
	VerifyLogin(ctx context.Context, username, password string, user model.User) (bool, error)
	GetUserData(ctx context.Context, username string) (model.User, error)
	CreateUserSession(ctx context.Context, userID string) (model.UserSession, error)
	CheckSession(ctx context.Context, data model.UserSession) (userID string, err error)
}
