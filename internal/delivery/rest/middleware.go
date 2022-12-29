package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go-restaurant-kelas-work/internal/model/constant"
	"go-restaurant-kelas-work/internal/tracing"
	"go-restaurant-kelas-work/internal/usecase/resto"
	"golang.org/x/net/context"
	"net/http"
)

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://restoku.com"},
	}))
	// handler panic
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}

type authMiddleware struct {
	restoUsecase resto.Usecase
}

func (a *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := tracing.CreateSpan(c.Request().Context(), "CheckAuth")
		defer span.End()
		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}
		userID, err := a.restoUsecase.CheckSession(ctx, sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}
		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))
		if err := next(c); err != nil {
			return err
		}
		return nil
	}
}

func GetAuthMiddleware(restoUsecase resto.Usecase) *authMiddleware {
	return &authMiddleware{
		restoUsecase: restoUsecase,
	}
}
